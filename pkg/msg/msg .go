package msg

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"roofix/pkg/util/log"
)

type MsgKey string

//go:embed msg.json
var file []byte
var messages []struct {
	Key  MsgKey `json:"key"`
	EnUS string `json:"en-US"`
}
var locale = "en-US"

func init() {
	if err := json.Unmarshal(file, &messages); err != nil {
		log.Error(err)
	}
}

const (
	BadRequest  MsgKey = "400"
	ServerError MsgKey = "500"
	// AlreadyExists need arg entityName
	AlreadyExists MsgKey = "AlreadyExists"
	// AccountLockedUntil will requires an arg
	AccountLockedUntil MsgKey = "AccountLockedUntil"
	ConfirmPwdMismatch MsgKey = "ConfirmPwdMismatch"
	OneOfParamMissing  MsgKey = "OneOfParamMissing"
	// ParamMissing require param name arg
	ParamMissing MsgKey = "ParamMissing"
	// FailedToSave require entity name arg
	FailedToSave MsgKey = "FailedToSave"
	// FailedToFind will need name arg
	FailedToFind            MsgKey = "FailedToFind"
	FailedToInstallPkg      MsgKey = "FailedToInstallPkg"
	WrongLoginCred          MsgKey = "WrongLoginCred"
	WrongPwd                MsgKey = "WrongPwd"
	AccountNotActive        MsgKey = "AccountNotActive"
	PartnerAccountNotActive MsgKey = "PartnerAccountNotActive"
	NotAuthorized           MsgKey = "NotAuthorized"
	NotAllowed              MsgKey = "NotAllowed"
	AuthorizedToOrgAdmin    MsgKey = "AuthorizedToOrgAdmin"
	NotFound                MsgKey = "NotFound"
	JobNotAssigned          MsgKey = "JobNotAssigned"
)

func Read(key MsgKey, args ...interface{}) string {
	for _, m := range messages {
		if m.Key != key {
			continue
		}

		switch locale {
		case "en-US":
			if len(args) == 0 {
				return fmt.Sprint(m.EnUS)
			}

			return fmt.Sprintf(m.EnUS, args)
		}
	}

	log.Error(errors.New(fmt.Sprintf("message wih key \"%s\" does not found", key)))
	return ""
}

func AsError(key MsgKey, args ...interface{}) error {
	return errors.New(Read(key, args...))
}
