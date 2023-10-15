package account

import (
	"context"
	"roofix/ent"
	"roofix/ent/user"
	"roofix/pkg/enum"
)

// checkAdminCanSaveAdmin is to confirm that only Admin can create new Admin user account
func checkAdminSavingAdmin(ctxRole, inpRole enum.Role) bool {
	admin := enum.RoleAdmin
	if ctxRole == admin || inpRole != admin {
		return true
	}

	// validate, only Admin can update another Admin user's
	return inpRole == admin && ctxRole == admin
}

func checkEmailInUse(ctx context.Context, email, skipID string) bool {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.User.Query().Where(user.EmailEQ(email))
	if skipID != "" {
		qry.Where(user.IDNEQ(skipID))
	}

	exist, _ := qry.Exist(ctx)

	return exist
}
