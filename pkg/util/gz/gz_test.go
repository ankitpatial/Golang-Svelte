/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  14/06/22, 4:48 PM
 */

package gz

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	data := "some string data to zip"
	gz, err := Zip([]byte(data))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("ziped data: %s\n", string(gz))
	// unzip now
	d, err := UnZip(gz)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("got back: %s", d)
}
