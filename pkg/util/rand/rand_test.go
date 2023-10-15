/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  06/04/22, 12:16 PM
 */

package rand

import (
	"testing"
)

func TestRandomPin(t *testing.T) {
	a := []int{5, 6, 7, 8, 9, 10}
	for _, i := range a {
		pin := RandomPin(i)
		t.Log("pin", pin)
		if pin == "" {
			t.Errorf("expected a pin of len %d but got empty", i)
			break
		}

		if len(pin) != i {
			t.Errorf("expected len: %d got %d for pin %s", i, len(pin), pin)
			break
		}
	}
}

func FuzzRandomPin(f *testing.F) {
	testCase := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range testCase {
		f.Add(i)
	}

	f.Fuzz(func(t *testing.T, i int) {
		pin := RandomPin(i)
		if pin == "" {
			t.Errorf("expected a pin of len %d but got empty", i)
			return
		}

		if len(pin) != i {
			t.Errorf("expected len: %d got %d for pin %s", i, len(pin), pin)
			return
		}
	})

}
