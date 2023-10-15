/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  14/06/22, 1:22 PM
 */

package gz

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Zip(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}

	if err := gz.Flush(); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func UnZip(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data)
	var r io.Reader
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}

	var resB bytes.Buffer
	if _, err := resB.ReadFrom(r); err != nil {
		return nil, err
	}

	return resB.Bytes(), nil
}
