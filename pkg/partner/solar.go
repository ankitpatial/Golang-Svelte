package partner

import (
	"context"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/pkg/util/uid"
)

// QuickCreateSolar is a quick way to create solar partner with only name and id
// must be called to help create solar partner on estimate creation
func QuickCreateSolar(ctx context.Context, db *ent.Client, externalID, name string) (string, error) {
	if externalID == "" || name == "" {
		return "", msg.AsError(msg.ParamMissing, "externalID or name")
	}

	id := uid.ULID()
	qry := db.Partner.Create().SetID(id).SetExternalID(externalID).SetName(name).SetType(enum.PartnerSolar)
	if err := qry.Exec(ctx); err != nil {
		return "", err
	}

	log.Info("created solar partner %s", name)
	return id, nil
}
