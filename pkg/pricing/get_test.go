/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package pricing

import (
	"context"
	"testing"
)

func TestGetPricing(t *testing.T) {
	res, err := ForAll(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Log(len(res), "records found")
}

func TestPricingForPostal(t *testing.T) {
	ctx := context.Background()
	t.Run("check for wrong value", func(t *testing.T) {
		list, _, err := priceList(ctx, "xyz", "1231", 0, 0)
		if err != nil {
			t.Error(err)
			return
		}

		if list != nil {
			t.Error("expecting list be nil")
		}

	})

	t.Run("exact match", func(t *testing.T) {
		list, note, err := priceList(ctx, "Alabama", "35233", 33.5118197, -86.8011326)
		if err != nil {
			t.Error("expected err as nil but got", err)
			return
		}

		if list == nil {
			t.Error("expected list has values but got nil")
			return
		}

		t.Log(len(list), "items found", note)
	})

	t.Run("nearby zip", func(t *testing.T) {
		list, note, err := priceList(ctx, "New York", "10018", 40.7524632, -73.9847395)
		if err != nil {
			t.Error("expected err as nil but got", err)
			return
		}

		if list == nil {
			t.Error("expected list has values but got nil")
			return
		}

		t.Log(len(list), "items found", note)
	})
}
