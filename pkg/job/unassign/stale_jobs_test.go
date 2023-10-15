/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package unassign

import (
	"context"
	"testing"
)

func TestClearExpiredInvites(t *testing.T) {
	ctx := context.Background()
	err := StaleJobs(ctx)
	if err != nil {
		t.Error(err)
	}
}
