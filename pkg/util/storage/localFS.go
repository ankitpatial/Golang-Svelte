/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package storage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"roofix/config"
	"roofix/pkg/util/dir"
	"roofix/pkg/util/file"
	"roofix/pkg/util/str"
	"time"
)

type LocalFS struct {
	HomeDir string
}

func (s *LocalFS) GetObject(_ context.Context, bucket, key string) ([]byte, error) {
	path := filepath.Join(s.HomeDir, bucket, key)
	info("GET %s", path)

	if err := dir.Create(filepath.Dir(path)); err != nil {
		return nil, err
	}

	if !file.Exists(path) {
		return nil, nil
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *LocalFS) Exists(_ context.Context, bucket, key string) bool {
	path := filepath.Join(s.HomeDir, bucket, key)
	return file.Exists(path)
}

func (s *LocalFS) PresignGetObject(_ context.Context, bucket, key string, _ time.Duration) (string, error) {
	qs := fmt.Sprintf("?bucket=%s&key=%s", bucket, key)
	path := fmt.Sprintf("%s/file-download%s", config.Read().Website.Url, qs)
	info("GetSignedUrl %s", path+qs)
	return path, nil
}

func (s *LocalFS) PresignPutObject(
	_ context.Context, bucket, key string, contentType *string, meta map[string]string, _ time.Duration,
) (string, error) {
	qs := fmt.Sprintf("?bucket=%s&key=%s", bucket, key)
	if ct := str.Val(contentType); ct != "" {
		qs += "&contentType=" + ct
	}
	for key, element := range meta {
		qs = fmt.Sprintf("%s&%s=%s", qs, key, element)
	}
	path := fmt.Sprintf("%s/file-upload%s", config.Read().Website.Url, qs)
	info("PresignPutObject %s", path)
	return path, nil
}

func (s *LocalFS) PutObject(_ context.Context, bucket, key string, data []byte, _ *string) error {
	path := filepath.Join(s.HomeDir, bucket, key)
	info("PutObject %s", path)

	if err := dir.Create(filepath.Dir(path)); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0740)
}

func (s *LocalFS) CopyObject(_ context.Context, sourceBucket, sourceKey, destBucket, destKey string) error {
	info("MoveObject from: %s/%s, to: %s/%s", sourceBucket, sourceKey, destBucket, destKey)

	return errors.New("not implemented")
}

func (s *LocalFS) DeleteObject(_ context.Context, bucket, key string) error {
	path := filepath.Join(s.HomeDir, bucket, key)
	info("DeleteObject %s", path)

	if file.Exists(path) {
		return os.Remove(path)
	}

	return nil
}
