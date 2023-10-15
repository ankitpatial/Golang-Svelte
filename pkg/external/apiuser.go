package external

import (
	"bytes"
	"context"
	"encoding/json"
	"roofix/ent"
	"roofix/ent/apiuser"
	"roofix/pkg/msg"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"roofix/pkg/util/rand"
	"roofix/server/model"
	"strings"
)

// AuthenticateApiUser an api user with random password
func AuthenticateApiUser(ctx context.Context, username, pwd string) (string, error) {
	if username == "" || pwd == "" {
		return "", msg.AsError(msg.OneOfParamMissing)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// check name is available
	u, err := db.ApiUser.Query().Where(apiuser.UsernameEqualFold(username)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", err
	}

	if u == nil {
		return "", msg.AsError(msg.NotFound, "Website User")
	}

	if !u.Active {
		return "", msg.AsError(msg.AccountNotActive)
	}

	if match := crypt.CompareHash(u.PwdHash, pwd); !match {
		return "", msg.AsError(msg.WrongPwd)
	}

	return u.ID, nil
}

func GetApiUser(ctx context.Context, username string) *ent.ApiUser {
	if username == "" {
		return nil
	}

	db := ent.GetClient()
	defer db.CloseClient()

	u, err := db.ApiUser.Query().Where(apiuser.UsernameEqualFold(username)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil
	}

	return u
}

func CreateApiUser(ctx context.Context, name string) (string, error) {
	n := strings.Trim(name, "")
	if n == "" {
		return "", msg.AsError(msg.ParamMissing, "name")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// check name is available
	if e, _ := db.ApiUser.Query().Where(apiuser.UsernameEqualFold(n)).Exist(ctx); e {
		return "", msg.AsError(msg.AlreadyExists, "Website Username")
	}

	// generate random pwd
	pwd, hash := newPwdWithHash()
	err := db.ApiUser.Create().
		SetUsername(n).
		SetPwdHash(hash).
		Exec(ctx)
	if err != nil {
		return "", err
	}

	return pwd, nil
}

func Update(ctx context.Context, input *model.APIUserInput) error {
	if input == nil {
		return msg.AsError(msg.ParamMissing, input)
	}

	if input.CbAPIAuth == nil {
		log.Warn("nothing new to save")
		return nil
	}

	db := ent.GetClient()
	defer db.CloseClient()

	auth := input.CbAPIAuth

	var out map[string]string
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input.CbAPIEndpoints); err != nil {
		return err
	}

	if err := json.NewDecoder(buf).Decode(&out); err != nil {
		return err
	}

	qry := db.ApiUser.UpdateOneID(input.ID).
		SetCbAPIAuth(apiuser.CbAPIAuth(auth.String())).
		SetCbAPIEndpoints(out)

	url := input.CbAPIURL
	if url != nil {
		qry.SetCbAPIURL(*url)
	} else {
		qry.ClearCbAPIURL()
	}

	usr := input.CbAPIUser
	if usr != nil {
		qry.SetCbAPIUser(*usr)
	} else {
		qry.ClearCbAPIUser()
	}

	pwd := input.CbAPIPwd
	if pwd != nil {
		qry.SetCbAPIPwd(*pwd)
	} else {
		qry.ClearCbAPIPwd()
	}

	tkn := input.CbAPIToken
	if tkn != nil {
		qry.SetCbAPIToken(*tkn)
	} else {
		qry.ClearCbAPIToken()
	}

	return qry.Exec(ctx)
}

func SetNewApiUserPwd(ctx context.Context, id string) (string, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// check name is available
	if e, _ := db.ApiUser.Query().Where(apiuser.ID(id)).Exist(ctx); !e {
		return "", msg.AsError(msg.NotFound, "Website User")
	}

	pwd, hash := newPwdWithHash()
	if err := db.ApiUser.UpdateOneID(id).SetPwdHash(hash).Exec(ctx); err != nil {
		return "", err
	} else {
		return pwd, nil
	}
}

func ChangeApiUserStatus(ctx context.Context, id string, isActive bool) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// check name is available
	if e, _ := db.ApiUser.Query().Where(apiuser.ID(id)).Exist(ctx); !e {
		return false, msg.AsError(msg.NotFound, "Website User")
	}

	if err := db.ApiUser.UpdateOneID(id).SetActive(isActive).Exec(ctx); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func newPwdWithHash() (string, string) {
	pwd := rand.RandomStr(25, 2, 5, 5)
	hash := crypt.Hash(pwd)

	return pwd, hash
}
