/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"testing"
)

func TestApiToken(t *testing.T) {
	ctx := context.Background()
	_, token, err := apiToken(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	if token == "" {
		t.Error("failed to get token")
		return
	}

	t.Log("got token ", token)
}
