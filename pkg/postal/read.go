/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package postal

import (
	"bytes"
	"encoding/csv"
	"io"
	"roofix/pkg/util/log"
	"strconv"
	"strings"
)

func ReadUSGeoNamesCsv(data []byte) []*Detail {
	var result []*Detail
	r := csv.NewReader(bytes.NewReader(data))
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Error(err)
			return nil
		}

		arr := strings.Split(rec[0], "\t")
		lat, _ := strconv.ParseFloat(arr[9], 64)
		lng, _ := strconv.ParseFloat(arr[10], 64)
		acc, _ := strconv.ParseUint(arr[11], 10, 8)
		result = append(result, &Detail{
			Country:   arr[0],
			Zip:       arr[1],
			City:      arr[2],
			State:     arr[3],
			StateAbr:  arr[4],
			Province:  arr[5],
			Latitude:  lat,
			Longitude: lng,
			Accuracy:  uint8(acc),
		})
	}

	return result
}
