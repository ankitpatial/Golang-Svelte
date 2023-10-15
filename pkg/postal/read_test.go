/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package postal

import (
	"os"
	"testing"
)

func TestReadUSGeoNamesCsv(t *testing.T) {
	//ctx: = context.Background()
	data, _ := os.ReadFile("./sample/us.csv")
	res := ReadUSGeoNamesCsv(data)
	if len(res) == 0 {
		t.Error("No data found")
		return
	}
}
