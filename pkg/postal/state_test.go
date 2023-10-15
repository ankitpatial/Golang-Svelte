/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package postal

import "testing"

func TestSearch(t *testing.T) {
	res := Search("a", 10)
	t.Log("got ", len(res))
}
