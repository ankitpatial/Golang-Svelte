/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

/*
	Simsaw Software Pvt. Ltd.
	Author: Ankit Patial <ankit@simsaw.com>
*/

package response

import (
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
	"roofix/pkg/util/log"
	"roofix/pkg/util/validate"
)

type ApiResponse struct {
	Data             interface{}                      `json:"data,omitempty"`
	Error            string                           `json:"error,omitempty"`
	ValidationErrors []*validate.FieldValidationError `json:"validationErrors,omitempty"`
}

func Ok(w http.ResponseWriter, data interface{}) {
	res(w, http.StatusOK, ApiResponse{
		Data: data,
	})
}

func Created(w http.ResponseWriter, data interface{}) {
	res(w, http.StatusCreated, ApiResponse{
		Data: data,
	})
}

func String(w http.ResponseWriter, r *http.Request, v string) {
	render.PlainText(w, r, v)
}

func Html(w http.ResponseWriter, r *http.Request, v string) {
	render.HTML(w, r, v)
}

func BadRequest(w http.ResponseWriter, err error) {
	if validate.IsStructValidationErr(err) {
		res(w, http.StatusBadRequest, ApiResponse{
			Error:            err.Error(),
			ValidationErrors: err.(*validate.StructValidationError).Errors,
		})
		return
	}

	res(w, http.StatusBadRequest, ApiResponse{
		Error: err.Error(),
	})
}

func ServerError(w http.ResponseWriter, err error) {
	res(w, http.StatusInternalServerError, ApiResponse{
		Error: err.Error(),
	})
}

func res(w http.ResponseWriter, status int, v ApiResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	b, err := json.Marshal(v)
	if err != nil {
		log.Error(err)
	}

	_, _ = w.Write(b)
}
