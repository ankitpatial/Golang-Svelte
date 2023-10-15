package installation

import (
	"context"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/partner"
	"time"
)

func Approve(ctx context.Context, id, ownerEmail, ownerPhone string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	j, err := db.InstallationJob.Get(ctx, id)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	qry := db.InstallationJob.
		UpdateOneID(id).
		SetApproval(enum.ApprovalApproved).
		SetApprovalAt(now).
		SetStatus(enum.InstallationStatusNew).
		SetStatusAt(now)

	switch j.Type {
	case enum.InstallationTypeHVAC:
		qry.SetAssignedPartnerID(partner.EcostellaID)
	case enum.InstallationTypeSmartHome:
		qry.SetAssignedPartnerID(partner.ResideoID)
	}

	if ownerEmail != "" {
		qry.SetOwnerEmail(ownerEmail)
	}

	if ownerPhone != "" {
		qry.SetOwnerPhone(ownerPhone)
	}

	return qry.Exec(ctx)
}

func Deny(ctx context.Context, id, reason string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	return db.InstallationJob.
		UpdateOneID(id).
		SetApproval(enum.ApprovalDenied).
		SetApprovalAt(time.Now().UTC()).
		SetDenyReason(reason).
		Exec(ctx)
}

func RemoveDeny(ctx context.Context, id string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	return db.InstallationJob.
		UpdateOneID(id).
		SetApproval(enum.ApprovalPending).
		SetApprovalAt(time.Now().UTC()).
		SetDenyReason("").
		Exec(ctx)
}
