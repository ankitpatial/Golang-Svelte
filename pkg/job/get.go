/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package job

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/job"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/pkg/util/log"
)

func Address(j *ent.Job) string {
	if j == nil || j.Edges.HomeOwner == nil {
		return ""
	}

	ho := j.Edges.HomeOwner
	return fmt.Sprintf("%s %s, %s, %s %s", ho.StreetNumber, ho.StreetName, ho.City, ho.State, ho.Zip)
}

func AddressWithNL(j *ent.Job) string {
	if j == nil || j.Edges.HomeOwner == nil {
		return ""
	}

	ho := j.Edges.HomeOwner
	return fmt.Sprintf("%s %s, %s\n%s %s", ho.StreetNumber, ho.StreetName, ho.City, ho.State, ho.Zip)
}

func withCreatorAndRequester(ctx context.Context, id string) *ent.Job {
	db := ent.GetClient()
	defer db.CloseClient()

	j, err := db.Job.Query().
		WithCreator(func(c *ent.UserQuery) { c.Select(user.FieldID) }).
		WithRequester(func(p *ent.PartnerQuery) { p.Select(partner.FieldID) }).
		WithHomeOwner().
		Where(job.ID(id)).
		Select(job.FieldProgress).
		Only(ctx)
	if err != nil {
		log.Error(err)
	}

	return j
}
