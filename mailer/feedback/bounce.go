/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package feedback

import (
	"context"
	"encoding/json"
	"fmt"
	"roofix/config"
	"roofix/pkg/util/log"
	"roofix/pkg/util/ptr"
	"roofix/pkg/util/storage"
	"roofix/pkg/util/str"
)

func Bounce(ctx context.Context, mail *MailInfo, bounce *BounceInfo) {
	if mail == nil || bounce == nil {
		log.Warn("one of required param is missing")
		return
	}

	if bounce.Type != "Permanent" {
		log.Info("ignore reporting '%s' bounce type", bounce.Type)
		return
	}

	for _, v := range bounce.Recipients {
		log.Info("bounce: %s", v.Email)
		reportBounce(ctx, v)
	}
}

func reportBounce(ctx context.Context, recipient *BounceRecipient) {
	bucket := config.LogBucket()
	key := fmt.Sprintf("mail/bounce/%s", str.TrimAndToLower(recipient.Email))
	data, _ := json.Marshal(recipient)
	if err := storage.PutObject(ctx, bucket, key, data, ptr.Str("text/plain")); err != nil {
		log.Error(err)
	}
}
