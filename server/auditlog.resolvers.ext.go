package server

import (
	"context"
	"roofix/ent"
	"roofix/ent/apiuser"
	"roofix/ent/auditlog"
	"roofix/pkg/util/str"
	"roofix/server/model"
	"strings"
)

func queryAuditLog(
	ctx context.Context, apiUserID, search *string, page model.PageInput, orderBy *ent.AuditLogOrder,
) (*ent.AuditLogConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.AuditLog.Query().WithUser().WithAPIUser()
	apiUID := strings.TrimSpace(str.Val(apiUserID))
	if apiUID != "" {
		qry.Where(auditlog.HasAPIUserWith(apiuser.ID(apiUID)))
	}

	q := strings.TrimSpace(str.Val(search))
	if q != "" {
		qry.Where(auditlog.Or(
			auditlog.ActionContainsFold(q),
			auditlog.DescriptionContainsFold(q),
		))
	}
	return qry.Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithAuditLogOrder(orderBy))
}
