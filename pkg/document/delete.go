package document

import (
	"context"
	"roofix/ent"
	"roofix/pkg/util/log"
	"roofix/pkg/util/storage"
)

// Delete document from db and storage
func Delete(ctx context.Context, id string) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// pull record
	d, err := db.Document.Get(ctx, id)
	if err != nil {
		return false, err
	}

	// delete from DB
	if err := db.Document.DeleteOneID(id).Exec(ctx); err != nil {
		return false, err
	}

	log.Info("DELETED doc(%s) from database", id)

	// remove from storage
	bucket := d.Bucket
	key := d.Key
	if err := storage.DeleteObject(ctx, bucket, key); err != nil {
		// by now entry is deleted from db
		// so just report error and do not return it
		log.Error(err)
	} else {
		log.Info("DELETED doc(%s) from storage", id)
	}

	return false, err
}
