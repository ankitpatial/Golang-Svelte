package external

import (
	"context"
	"roofix/ent"
	"roofix/ent/apiuser"
	"roofix/ent/apiusertoken"
	"roofix/pkg/msg"
)

func SaveApiUserToken(ctx context.Context, username, tokenType, tokenID, refreshTokenID string) error {
	// validate input
	if username == "" || tokenType == "" || tokenID == "" {
		return msg.AsError(msg.OneOfParamMissing)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// get api
	apiUser, err := db.ApiUser.Query().Where(apiuser.UsernameEqualFold(username)).First(ctx)
	if err != nil && ent.IsNotFound(err) {
		return msg.AsError(msg.NotFound, "Api User")
	} else if err != nil {
		return err
	}

	// save token
	_, err = db.ApiUserToken.
		Create().
		SetAPIUserID(apiUser.ID).
		SetTokenType(tokenType).
		SetTokenID(tokenID).
		SetRefreshTokenID(refreshTokenID).
		Save(ctx)

	return err
}

func GetApiUserToken(ctx context.Context, tokenID string) (*ent.ApiUserToken, error) {
	// validate input
	db := ent.GetClient()
	defer db.CloseClient()

	// get toke
	return db.ApiUserToken.Query().Where(apiusertoken.TokenID(tokenID)).First(ctx)
}
