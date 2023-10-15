package document

import (
	"context"
	"roofix/config"
	"roofix/pkg/account"
	"roofix/pkg/util/storage"
	"roofix/pkg/util/validate"
	"time"
)

// UploadRequest will presign upload url & will save a document entry in DB with doc.read = false.
// Upon file upload to storage, db.read will be set as 'true'(lambda 'docUploadFn' will do it)
func UploadRequest(ctx context.Context, key string, inp *Input, meta map[string]string) (*Info, error) {
	// validate input
	if err := validate.Struct(inp); err != nil {
		return nil, err
	}

	// set creator to meta
	if meta == nil {
		meta = make(map[string]string)
	}
	meta["creatorID"] = account.CtxUserID(ctx)
	meta["creator"] = account.CtxUserNameEmail(ctx)

	// presign URL
	bucket := config.DataBucket()
	span := time.Minute * 60
	url, err := storage.PutSignedUrl(ctx, bucket, key, inp.ContentType, meta, span)
	if err != nil {
		return nil, err
	}

	// save document info to DB
	info, err := Save(ctx, bucket, key, inp, false)
	if err != nil {
		return nil, err
	}
	info.UploadUrl = url
	info.Meta = meta
	return info, err // done!
}
