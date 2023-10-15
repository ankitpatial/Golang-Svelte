/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"roofix/ent"
	"roofix/pkg/util/uid"
	"testing"
)

func TestPlaceOrder(t *testing.T) {
	ctx := context.Background()
	// use a job from DB
	db := ent.GetClient()
	job, _ := db.Job.Get(ctx, "6037d665-38b5-4939-a169-7978c9fa14bb")
	refID := uid.ULID()
	if res, err := PlaceOrder(ctx, job, refID); err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("placed order %v", res)
	}

}
