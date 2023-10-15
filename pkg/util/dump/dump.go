package dump

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"os"
	"path/filepath"
	"roofix/config"
	"roofix/pkg/util/dir"
	"roofix/pkg/util/ptr"
	"roofix/pkg/util/rt"
	"runtime"
	"time"
)

type withStack struct {
	error
	*rt.Stack
}

func Error(err error) {
	if err == nil {
		return
	}

	caller := rt.CallerInfo(runtime.Caller(1))

	// print
	fmt.Printf("ERROR  [%s] %s", caller, err)

	// wrap with caller trace
	err = &withStack{
		err,
		rt.Callers(),
	}

	ErrorLog(time.Now().UTC().Format(time.RFC3339), caller, fmt.Sprintf("%+v", err))
}

func ErrorLog(caller, e, stamp string) {
	b, err := json.Marshal(map[string]string{
		"timeStamp": stamp,
		"caller":    caller,
		"error":     e,
	})
	if err != nil {
		fmt.Printf("ERROR %s", err)
		return
	}

	t := time.Now().UTC()
	bucket := config.LogBucket()
	key := fmt.Sprintf("%s/%s.json", t.Format(time.DateOnly), t.Format("15-04-05.000000000"))

	if config.IsDevEnv() {
		err = toDisk(bucket, key, b)
	} else {
		err = toS3(context.Background(), bucket, key, b)
	}

	if err != nil {
		fmt.Printf("ERROR %s", err)
	}
}

func toS3(ctx context.Context, bucket, key string, data []byte) error {
	inp := &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		ACL:         types.ObjectCannedACLPrivate,
		Body:        bytes.NewReader(data),
		ContentType: ptr.Str("application/json"),
	}

	if _, err := s3.NewFromConfig(config.AwsConfig()).PutObject(ctx, inp); err != nil {
		return err
	}

	return nil
}

func toDisk(bucket, key string, data []byte) error {
	h, _ := config.AppHome()
	path := filepath.Join(h, bucket, key)
	fmt.Printf("PutObject %s", path)

	if err := dir.Create(filepath.Dir(path)); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0740)
}
