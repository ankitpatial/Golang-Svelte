package rfx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"roofix/ent"
	"roofix/ent/user"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"roofix/pkg/util/uid"
	"roofix/pkg/util/validate"
	"roofix/server/router/response"
)

type UserInput struct {
	ID                      string  `json:"id" validate:"required"`
	Email                   *string `json:"email"`
	FirstName               *string `json:"firstName"`
	LastName                *string `json:"lastName"`
	Phone                   *string `json:"phone"`
	Location                *string `json:"location"` // formatted Google Map Address
	Role                    *uint8  `json:"role"`
	InternalRole            *uint8  `json:"internalRole"`
	Status                  *uint8  `json:"status"`
	AcceptedGeneralTerms    *bool   `json:"acceptedGeneralTerms"`
	AcceptedTermsAndPrivacy *bool   `json:"acceptedTermsAndPrivacy"`
	//CompanyID               *string `json:"companyID"`
	//FullName string
	//Theme        uint8 // Dark: 1, Light: 2
}

// SaveUserHandler to save user
func SaveUserHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	ctx := r.Context()
	action := audit.UserSave
	apiUID := account.CtxApiUserID(ctx)
	failed := func(err error) {
		if id != "" {
			audit.ApiOpFailed(ctx, action, apiUID, fmt.Errorf("save with with refID: %s failed with error, %v", id, err))
		} else {
			audit.ApiOpFailed(ctx, action, apiUID, err)
		}
		// return back with 200 with failureStatus
		response.Ok(w, SaveResponse{
			FailureStatus: err.Error(),
		})
	}

	// bind modal
	var inp UserInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		audit.ApiOpFailed(ctx, action, apiUID, err)
		failed(err)
		return
	}

	if err := validate.Struct(inp); err != nil {
		failed(err)
		return
	}

	db := ent.GetClient()
	defer db.CloseClient()

	//
	// user existence must be checked firs by externalID then by email
	//

	var exist bool
	// first check, by externalID
	if u, err := ent.GetClient().User.Query().Where(user.ExternalID(inp.ID)).Select(user.FieldID).First(ctx); err != nil {
		if ent.IsNotFound(err) {
			// second check, by email
			if inp.Email != nil {
				if u, err = db.User.Query().Where(user.Email(*inp.Email)).Select(user.FieldID).First(ctx); err != nil {
					if !ent.IsNotFound(err) {
						failed(log.Wrap(err, "unexpected"))
						return
					}
				} else {
					log.Info("user exist, found by email")
					exist = true
					id = u.ID
				}
			}

		} else {
			failed(log.Wrap(err, "unexpected"))
			return
		}
	} else {
		log.Info("user exist, found by externalID")
		exist = true
		id = u.ID
	}

	var message string
	if exist {
		log.Info("update user")
		if saveErr := updateUser(ctx, db, id, &inp); saveErr != nil {
			failed(log.Wrap(saveErr, "update user"))
			return
		}

		message = "successfully updated user"
		logUserActivity(ctx, db, "created", id, apiUID)
	} else {
		log.Info("create user")
		id = uid.ULID()
		if saveErr := createUser(ctx, db, id, &inp); saveErr != nil {
			failed(log.Wrap(saveErr, "create user"))
			return
		}

		message = "successfully created user"
		logUserActivity(ctx, db, "created", id, apiUID)
	}

	response.Ok(w, SaveResponse{
		Message: message,
		Success: true,
	})
}

// createUser in our database
func createUser(ctx context.Context, db *ent.Client, id string, inp *UserInput) error {
	// validate require fields to create user
	// must have email, firstName, lastName & Role
	if inp.Email == nil || inp.FirstName == nil || inp.LastName == nil || inp.Role == nil {
		return fmt.Errorf("email, firstName, lastName & role are required")
	}

	// create
	return db.User.Create().
		SetID(id).
		SetExternalID(inp.ID).
		SetEmail(*inp.Email).
		SetFirstName(*inp.FirstName).
		SetLastName(*inp.LastName).
		SetPwd(crypt.Hash(uid.ULID())). // set random password
		SetRole(toUserRole(num.ReadUInt8Ptr(inp.Role), num.ReadUInt8Ptr(inp.InternalRole))).
		SetStatus(toUserStatus(num.ReadUInt8Ptr(inp.Status))).
		SetNillablePhone(inp.Phone).
		SetNillableLocation(inp.Phone).
		SetNillableAcceptedGeneralTerms(inp.AcceptedGeneralTerms).
		SetNillableAcceptedTermsNPrivacy(inp.AcceptedTermsAndPrivacy).
		Exec(ctx)
}

// updateUser in our database
// - do not update email
// - only update status as per internal role only if its one-of SubAdmin | Estimator | Reviewer
func updateUser(ctx context.Context, db *ent.Client, id string, inp *UserInput) error {
	u := db.User.UpdateOneID(id)
	// first name
	if inp.FirstName != nil {
		u.SetFirstName(*inp.FirstName)
	}
	// last name
	if inp.LastName != nil {
		u.SetLastName(*inp.LastName)
	}
	// status
	if inp.Status != nil {
		u.SetStatus(toUserStatus(num.ReadUInt8Ptr(inp.Status)))
	}
	// other nillable fields
	u.
		SetNillablePhone(inp.Phone).
		SetNillableLocation(inp.Location).
		SetNillableAcceptedGeneralTerms(inp.AcceptedGeneralTerms).
		SetNillableAcceptedTermsNPrivacy(inp.AcceptedTermsAndPrivacy)

	return u.Exec(ctx)
}

// log user activity
func logUserActivity(ctx context.Context, db *ent.Client, action, uID string, apiUID string) {
	creator := account.CtxApiUsername(ctx)
	desc := fmt.Sprintf("%s by APIUser: %s", action, creator)
	// log activity
	if e := db.UserActivity.Create().SetUserID(uID).SetDescription(desc).SetCreatorAPIID(apiUID).Exec(ctx); e != nil {
		log.Error(e)
	}
}

// toUserRole will convert given bubble app user role int to user role enum
func toUserRole(role, internalRol uint8) enum.Role {
	switch role {
	case 2:
		if internalRol > 0 {
			switch internalRol {
			case 1:
				return enum.RoleAdmin
			case 2:
				return enum.RoleSubAdmin
			case 3:
				return enum.RoleEstimator
			case 4:
				return enum.RoleReviewer
			}
		}
		return enum.RoleAdmin
	default:
		return enum.RoleSiteUser
	}
}

// toUserStatus will convert given bubble app user status int to user accountStatus
func toUserStatus(status uint8) enum.AccountStatus {
	switch status {
	case 1:
		return enum.AccountStatusActive
	case 2:
		return enum.AccountStatusDisabled
	default:
		return enum.AccountStatusPending
	}
}
