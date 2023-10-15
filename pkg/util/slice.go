/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  10/05/22, 12:55 PM
 */

package util

func IndexOf[T comparable](data []T, element T) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
