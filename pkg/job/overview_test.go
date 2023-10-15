/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package job

import (
	"context"
	"testing"
)

func TestOverview(t *testing.T) {
	ctx := context.Background()

	t.Run("by region", func(t *testing.T) {
		if _, err := OverviewByRegion(ctx); err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("by state", func(t *testing.T) {
		if _, err := OverviewState(ctx); err != nil {
			t.Error(err)
			return
		}
	})
}
