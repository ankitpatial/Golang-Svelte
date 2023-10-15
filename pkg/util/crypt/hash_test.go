/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import "testing"

func TestMD5Int(t *testing.T) {
	i := MD5Int("email")
	// 4,174,743,548,339,191,484
	// 9,223,372,036,854,775,807
	t.Log("i", i)
	i2 := MD5Int("EMAIL sd fs df sdf ksdkf skdf skdf ksdfksdfk ")
	t.Log("i2", i2)
}
