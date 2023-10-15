/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package document

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/document"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/notification"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/str"
	"strings"
	"time"
)

// docID will be hash of file path
func docID(bucket, key string) string {
	if strings.HasPrefix(key, "/") {
		return crypt.MD5Hash(fmt.Sprintf("%s%s", bucket, key))
	}
	return crypt.MD5Hash(fmt.Sprintf("%s/%s", bucket, key))
}

// Save a document entry in DB with active=false (marking a attempt of file upload)
// active=true will be set 'true' by storage upload watcher
// record with active=true will be considered uploaded successfully
func Save(ctx context.Context, bucket, key string, inp *Input, ready bool) (*Info, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	id := docID(bucket, key)
	userID := account.CtxUserID(ctx)

	var err error
	if exist, _ := db.Document.Query().Where(document.ID(id)).Exist(ctx); exist {
		qry := db.Document.
			UpdateOneID(id).
			SetUpdaterID(userID).
			SetName(inp.Name).
			SetFilename(inp.FileName).
			SetContentSize(inp.ContentSize).
			SetReady(ready)
		if inp.ContentType == nil {
			qry.ClearContentType()
		} else {
			qry.SetContentType(*inp.ContentType)
		}

		err = qry.Exec(ctx)
	} else {
		err = db.Document.Create().
			SetID(id).
			SetCreatorID(userID).
			SetBucket(bucket).
			SetKey(key).
			SetFolder(inp.Folder).
			SetDir(inp.Dir).
			SetSection(inp.Section).
			SetName(inp.Name).
			SetFilename(inp.FileName).
			SetNillableContentType(inp.ContentType).
			SetContentSize(inp.ContentSize).
			SetReady(ready).
			Exec(ctx)
	}

	if err != nil {
		return nil, err
	}

	return &Info{
		ID:          id,
		Key:         key,
		Folder:      inp.Folder,
		Section:     inp.Section,
		Name:        inp.Name,
		Filename:    inp.FileName,
		ContentType: inp.ContentType,
		ContentSize: inp.ContentSize,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Ready:       false,
	}, err
}

func SetReady(ctx context.Context, bucket, key string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	id := docID(bucket, key)
	// get document by its ID
	doc, err := db.Document.Query().
		Where(document.ID(id)).
		Select(
			document.FieldID,
			document.FieldFolder,
			document.FieldDir,
			document.FieldCreatorID,
		).
		Only(ctx)
	if err != nil {
		return err
	}

	// update document ready status
	if err = doc.Update().SetReady(true).Exec(ctx); err != nil {
		return err
	}

	// dir is in many cases the entity ID
	dir := str.Val(doc.Dir)
	if dir == "" {
		return nil
	}

	// notify creator on its own upload
	// supposed to be WS msg to update UI
	if doc.CreatorID != "" {
		channel := enum.ChannelPing
		switch doc.Folder {
		case enum.FolderEstimates:
			channel = enum.ChannelEstimate
		case enum.FolderJobDocs:
			channel = enum.ChannelJob
		}

		if channel != enum.ChannelPing {
			msg := &notification.Message{
				Channel: channel,
				Topic:   enum.TopicFileUpload,
				RefID:   dir,
				Title:   "New File Uploaded",
				Data: map[string]interface{}{
					"id": dir,
				},
			}
			notification.Send(ctx, msg, notification.ToUser(doc.CreatorID))
		}
	}

	return nil
}
