/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/fatih/color"
	"log"
	"roofix/config"
	"roofix/pkg/util/rt"
	"runtime"
	"time"
)

var (
	blue  = color.New(color.FgBlue)
	isDev = config.IsDevEnv()
)

type Client interface {
	GetObject(ctx context.Context, bucket, key string) ([]byte, error)
	PresignGetObject(ctx context.Context, bucket, key string, d time.Duration) (string, error)
	PresignPutObject(
		ctx context.Context, bucket, key string, contentType *string, meta map[string]string, d time.Duration,
	) (string, error)
	PutObject(ctx context.Context, bucket, key string, data []byte, contentType *string) error
	CopyObject(ctx context.Context, sourceBucket, sourceKey, destBucket, destKey string) error
	DeleteObject(ctx context.Context, bucket, key string) error
	Exists(ctx context.Context, bucket, key string) bool
}

var cl Client

func init() {
	if config.IsDevEnv() {
		h, _ := config.AppHome()
		cl = &LocalFS{
			HomeDir: h,
		}
		return
	}

	cl = &S3{
		Client: s3.NewFromConfig(config.AwsConfig()),
	}
}

func GetObject(ctx context.Context, bucket, key string) ([]byte, error) {
	return cl.GetObject(ctx, bucket, key)
}

func Exists(ctx context.Context, bucket, key string) bool {
	return cl.Exists(ctx, bucket, key)
}

func GetSignedUrl(ctx context.Context, bucket, key string, d time.Duration) (string, error) {
	return cl.PresignGetObject(ctx, bucket, key, d)
}

func PutSignedUrl(
	ctx context.Context, bucket, key string, contentType *string, meta map[string]string, d time.Duration,
) (string, error) {
	return cl.PresignPutObject(ctx, bucket, key, contentType, meta, d)
}

func PutObject(ctx context.Context, bucket, key string, data []byte, contentType *string) error {
	return cl.PutObject(ctx, bucket, key, data, contentType)
}

func CopyObject(ctx context.Context, sourceBucket, sourceKey, destBucket, destKey string) error {
	return cl.CopyObject(ctx, sourceBucket, sourceKey, destBucket, destKey)
}

func DeleteObject(ctx context.Context, bucket, key string) error {
	return cl.DeleteObject(ctx, bucket, key)
}
func info(msg string, arg ...any) {
	caller := rt.CallerInfo(runtime.Caller(1))

	a := make([]any, 0, 3+len(arg))
	a = append(a, ts(), blue.Sprint("INFO"), caller)
	a = append(a, arg...)

	log.Printf("%s %v %s - "+msg, a...)
}

func ts() string {
	if isDev {
		return time.Now().Format("2006-01-02T15:04:05")
	}

	return ""
}
