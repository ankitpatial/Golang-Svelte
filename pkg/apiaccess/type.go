/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package apiaccess

type Input struct {
	ID       string `json:"ID"`
	Name     string `json:"name" validate:"required"`
	URL      string `json:"URL" validate:"required,url"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Key      string `json:"key,omitempty"`
	Secret   string `json:"secret,omitempty"`
}
