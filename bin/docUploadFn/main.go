/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/config"
	"roofix/pkg/document"
	"roofix/pkg/enum"
	"roofix/pkg/job"
	"roofix/pkg/util/log"
	"strings"
	"sync"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, evt events.S3Event) (string, error) {
	config.PrintBuildInfo()
	fmt.Printf("message count: %d\n", len(evt.Records))

	var wg sync.WaitGroup
	for _, rec := range evt.Records {
		wg.Add(1)
		go process(ctx, rec, &wg)
	}
	wg.Wait()

	m := "Done!"
	fmt.Println(m)
	return m, nil
}

func process(ctx context.Context, r events.S3EventRecord, wg *sync.WaitGroup) {
	defer wg.Done()

	bucket := r.S3.Bucket.Name
	key := r.S3.Object.Key
	if err := document.SetReady(ctx, bucket, key); err != nil {
		log.Error(err)
	} else {
		parts := strings.Split(key, "/")
		l := len(parts)
		if l > 0 {
			switch enum.DocFolder(parts[0]) {
			case enum.FolderJobDocs:
				job.NotifyDocsChange(ctx, parts[1])
			}
		}
	}
}
