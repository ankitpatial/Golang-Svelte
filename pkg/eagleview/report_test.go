/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"testing"
)

func TestGetReport(t *testing.T) {
	ctx := context.Background()
	// use a job from DB
	reportID := uint(49807225)
	if res, err := GetReport(ctx, reportID); err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("placed order %v", res)
	}

}
