package installation

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/installationjob"
	"roofix/pkg/enum"
	"roofix/pkg/util"
	"roofix/pkg/util/convert"
	"roofix/pkg/util/str"
	"roofix/server/model"
)

func SearchPending(
	ctx context.Context, installationType enum.InstallationType, partnerID, creatorID *string, approvalIn []enum.Approval,
	search *string, dtRange []string,
	page *model.PageInput) (*ent.InstallationJobConnection, error,
) {
	db := ent.GetClient()
	defer db.CloseClient()
	by := &ent.InstallationJobOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.InstallationJobOrderFieldCreatedAt,
	}

	qry := db.InstallationJob.
		Query().
		Where(installationjob.TypeEQ(installationType)).
		Where(func(j *sql.Selector) {
			// partner or creator check
			if partnerID != nil || creatorID != nil {
				var p []*sql.Predicate
				// requester partnerI filter
				if partnerID != nil {
					p = append(p, sql.EQ(installationjob.RequestingPartnerColumn, *partnerID))
					p = append(p, sql.EQ(installationjob.AssignedPartnerColumn, *partnerID))
				}
				// creator filter
				if creatorID != nil {
					p = append(p, sql.EQ(installationjob.CreatorColumn, *creatorID))
				}

				j.Where(sql.Or(p...))
			}

			// status in
			if len(approvalIn) > 0 {
				j.Where(sql.In(j.C(installationjob.FieldApproval), convert.ToAny(approvalIn)...))
			}

			// date range
			if len(dtRange) == 2 {
				f := util.ParseISODate(dtRange[0])
				t := util.ParseISODate(dtRange[1])
				j.Where(sql.And(
					sql.GTE(j.C(installationjob.FieldCreatedAt), f),
					sql.LTE(j.C(installationjob.FieldCreatedAt), t),
				))
			}
		})

	return searchFields(qry, search).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithInstallationJobOrder(by))
}

func SearchApproved(
	ctx context.Context, installationType enum.InstallationType, partnerID, creatorID *string, statusIn []enum.InstallationStatus,
	search *string, dtRange []string,
	page *model.PageInput) (*ent.InstallationJobConnection, error,
) {
	db := ent.GetClient()
	defer db.CloseClient()
	by := &ent.InstallationJobOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.InstallationJobOrderFieldApprovalAt,
	}

	qry := db.InstallationJob.
		Query().
		Where(installationjob.TypeEQ(installationType)).
		Where(func(j *sql.Selector) {
			// partner or creator check
			if partnerID != nil || creatorID != nil {
				var p []*sql.Predicate
				// requester partnerI filter
				if partnerID != nil {
					p = append(p, sql.EQ(installationjob.RequestingPartnerColumn, *partnerID))
					p = append(p, sql.EQ(installationjob.AssignedPartnerColumn, *partnerID))
				}
				// creator filter
				if creatorID != nil {
					p = append(p, sql.EQ(installationjob.CreatorColumn, *creatorID))
				}

				j.Where(sql.Or(p...))
			}

			// status in
			if len(statusIn) > 0 {
				j.Where(sql.In(j.C(installationjob.FieldStatus), convert.ToAny(statusIn)...))
			}

			// approval
			j.Where(sql.EQ(j.C(installationjob.FieldApproval), enum.ApprovalApproved))
			// date range
			if len(dtRange) == 2 {
				f := util.ParseISODate(dtRange[0])
				t := util.ParseISODate(dtRange[1])
				j.Where(sql.And(
					sql.GTE(j.C(installationjob.FieldApprovalAt), f),
					sql.LTE(j.C(installationjob.FieldApprovalAt), t),
				))
			}
		})

	return searchFields(qry, search).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithInstallationJobOrder(by))
}

func searchFields(qry *ent.InstallationJobQuery, search *string) *ent.InstallationJobQuery {
	if s := str.TrimSpace(search); s != "" {
		// search by: company name, sales rep name, customer name, state, date created, etc.
		qry.Where(installationjob.Or(
			installationjob.OwnerNameContainsFold(s),
			installationjob.OwnerPhone(s),
			installationjob.OwnerEmail(s),
			installationjob.OwnerAddress(s),
		))
	}

	return qry
}
