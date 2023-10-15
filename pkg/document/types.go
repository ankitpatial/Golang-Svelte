package document

import (
	"roofix/pkg/enum"
	"time"
)

type Input struct {
	Folder      enum.DocFolder  `json:"folder" validate:"required"`
	Dir         string          `json:"dir"`
	Section     enum.DocSection `json:"section" validate:"required"`
	Name        string          `json:"name" validate:"required"`
	FileName    string          `json:"fileName"`
	ContentType *string         `json:"contentType"`
	ContentSize int64           `json:"contentSize"`
	Overwrite   bool            `json:"overwrite"`
}

type Info struct {
	ID          string            `json:"id"`
	Key         string            `json:"key"`
	Folder      enum.DocFolder    `json:"folder"`
	Section     enum.DocSection   `json:"section"`
	Name        string            `json:"name"`
	Filename    string            `json:"filename"`
	ContentType *string           `json:"contentType"`
	ContentSize int64             `json:"contentSize"`
	Ready       bool              `json:"ready"`
	UploadUrl   string            `json:"uploadUrl,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
	CreatedAt   time.Time         `json:"createdAt,omitempty"`
	UpdatedAt   time.Time         `json:"updatedAt,omitempty"`
}

type InfoShort struct {
	ID          string         `json:"id"`
	Key         string         `json:"key"`
	Folder      enum.DocFolder `json:"folder"`
	Name        string         `json:"name"`
	ContentType *string        `json:"contentType"`
	ContentSize int64          `json:"contentSize"`
	Ready       bool           `json:"ready"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
}
