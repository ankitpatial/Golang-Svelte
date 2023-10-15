/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"roofix/config"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/estimate"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/pkg/util/storage"
	"roofix/server/router/response"
)

const MaxUploadSize = (1024 * 1024) * 100 // 100MB

// fileDelivery from eagle view
func fileDelivery(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	action := audit.FileDelivery

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, err)
		return
	}

	if !r.URL.Query().Has("RefId") {
		err := msg.AsError(msg.ParamMissing, "RefId")
		response.BadRequest(w, err)
		audit.OpFailed(ctx, action, err)
		return
	}

	// check refID is valid
	estID := r.URL.Query().Get("RefId")
	_, err := estimate.GetByID(ctx, estID)
	if err != nil {
		response.BadRequest(w, msg.AsError(msg.NotFound, "Order"))
		audit.OpFailed(ctx, action, err)
		return
	}

	//reportId := r.URL.Query().Get("ReportId")
	//fFormatID := r.URL.Query().Get("FileFormatId")
	//fTypeID := r.URL.Query().Get("FileTypeId")

	// read file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		response.BadRequest(w, err)
		audit.OpFailed(ctx, action, err)
		return
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	// copy file to buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		response.ServerError(w, err)
		audit.OpFailed(ctx, action, err)
		return
	}

	// save file to storage
	fileName := fileHeader.Filename
	log.Info("fileName: %s", fileName)
	bucket := config.DataBucket()
	key := fmt.Sprintf("%s/%s/%s%s", enum.FolderEstimates, estID, fileHeader.Filename)
	if err := storage.PutObject(ctx, bucket, key, buf.Bytes(), nil); err != nil {
		response.ServerError(w, err)
		audit.OpFailed(ctx, action, err)
		return
	}

	// save measurement order history

	// done!
	response.Ok(w, "success")
}

func fileDeliveryConfirm(w http.ResponseWriter, r *http.Request) {

}
