package partner

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/ent"
	"roofix/ent/job"
	entPartner "roofix/ent/partner"
	"roofix/pkg/enum"
	"roofix/pkg/util/str"
)

// Stats by partner type
func Stats(ctx context.Context, partnerType enum.Partner, search *string, skip, take int) ([]*JobStats, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	var result []*JobStats
	q := str.Val(search)
	err := db.Partner.
		Query().
		Where(
			entPartner.TypeEQ(partnerType),
			func(p *sql.Selector) {
				j := sql.Table(job.Table)
				p.LeftJoin(j).On(p.C(entPartner.FieldID), j.C(job.RoofingPartnerColumn))

				p.GroupBy(entPartner.FieldID, entPartner.FieldName)
				p.OrderBy(entPartner.FieldName)

				jStatus := j.C(job.FieldProgress)

				if q != "" {
					p.Where(sql.ContainsFold(p.C(entPartner.FieldName), q))
				}

				tpl := "sum(if(t1.progress = '%s' and t1.progress_flag_at <= utc_timestamp, 1, 0)) as %s"
				p.Select(
					p.C(entPartner.FieldID),
					p.C(entPartner.FieldName),
					p.C(entPartner.FieldStatus),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as newCount", jStatus, enum.JobProgressNew),
					fmt.Sprintf(tpl, enum.JobProgressNew, "newCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as contactedCount", jStatus, enum.JobProgressCustomerContacted),
					fmt.Sprintf(tpl, enum.JobProgressCustomerContacted, "contactedCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as confirmedCount", jStatus, enum.JobProgressJobDetailsConfirmed),
					fmt.Sprintf(tpl, enum.JobProgressJobDetailsConfirmed, "confirmedCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as permittingCount", jStatus, enum.JobProgressPermitting),
					fmt.Sprintf(tpl, enum.JobProgressPermitting, "permittingCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as scheduledCount", jStatus, enum.JobProgressScheduled),
					fmt.Sprintf(tpl, enum.JobProgressScheduled, "scheduledCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as inProgressCount", jStatus, enum.JobProgressInProgress),
					fmt.Sprintf(tpl, enum.JobProgressInProgress, "inProgressCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as installedCount", jStatus, enum.JobProgressInstalled),
					fmt.Sprintf(tpl, enum.JobProgressInstalled, "installedCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as invoicedCount", jStatus, enum.JobProgressInvoiced),
					fmt.Sprintf(tpl, enum.JobProgressInvoiced, "invoicedCountFlagged"),
					fmt.Sprintf("sum(if(%s = '%s', 1, 0)) as delayedCount", jStatus, enum.JobProgressDelayed),
					fmt.Sprintf("count(%s) total", j.C(job.FieldID)),
					"sum(if(t1.progress_flag_at is not null and t1.progress_flag_at <= utc_timestamp, 1, 0))  totalFlagged",
				)
			}).
		Offset(skip).Limit(take).Select().
		Scan(ctx, &result)

	return result, err
}
