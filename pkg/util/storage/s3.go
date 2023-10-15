/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"strings"
	"time"
)

type S3 struct {
	Client *s3.Client
}

func (s *S3) GetObject(ctx context.Context, bucket, key string) ([]byte, error) {
	info("GET s3://%s/%s", bucket, key)
	in := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	out, err := s.Client.GetObject(ctx, in)
	if err != nil {
		if strings.Contains(err.Error(), "NoSuchKey") {
			return nil, nil
		}

		return nil, fmt.Errorf("storage GetObject: %w", err)
	}

	content := &bytes.Buffer{}
	_, err = content.ReadFrom(out.Body)
	if err != nil {
		return nil, err
	}

	return content.Bytes(), nil
}

func (s *S3) Exists(ctx context.Context, bucket, key string) bool {
	in := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := s.Client.HeadObject(ctx, in)
	if err != nil {
		return false
	}

	return true
}

func (s *S3) PresignGetObject(ctx context.Context, bucket, key string, d time.Duration) (string, error) {
	svc := s3.NewPresignClient(s.Client)

	inp := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	req, err := svc.PresignGetObject(ctx, inp, func(options *s3.PresignOptions) {
		options.Expires = d
	})

	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func (s *S3) PresignPutObject(
	ctx context.Context, bucket, key string, contentType *string, meta map[string]string, d time.Duration) (string, error) {
	svc := s3.NewPresignClient(s.Client)

	inp := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: contentType,
	}

	if len(meta) > 0 {
		inp.Metadata = meta
	}

	req, err := svc.PresignPutObject(ctx, inp, func(options *s3.PresignOptions) {
		options.Expires = d
	})

	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func (s *S3) PutObject(ctx context.Context, bucket, key string, data []byte, contentType *string) error {
	info("PutObject %s/%s", bucket, key)
	inp := &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		ACL:         types.ObjectCannedACLPrivate,
		Body:        bytes.NewReader(data),
		ContentType: contentType,
	}

	if _, err := s.Client.PutObject(ctx, inp); err != nil {
		return err
	}

	return nil
}

func (s *S3) CopyObject(ctx context.Context, sourceBucket, sourceKey, destBucket, destKey string) error {
	info("MoveObject from: %s/%s, to: %s/%s", sourceBucket, sourceKey, destBucket, destKey)
	inp := &s3.CopyObjectInput{
		CopySource:                aws.String(fmt.Sprintf("%s/%s", sourceBucket, sourceKey)),
		Bucket:                    &destBucket,
		Key:                       &destKey,
		ACL:                       types.ObjectCannedACLPrivate,
		ExpectedBucketOwner:       nil,
		ExpectedSourceBucketOwner: nil,
	}

	if _, err := s.Client.CopyObject(ctx, inp); err != nil {
		return err
	}

	return nil
}

func (s *S3) DeleteObject(ctx context.Context, bucket, key string) error {
	info("DeleteObject %s/%s", bucket, key)
	inp := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	if _, err := s.Client.DeleteObject(ctx, inp); err != nil {
		return err
	}

	return nil
}
