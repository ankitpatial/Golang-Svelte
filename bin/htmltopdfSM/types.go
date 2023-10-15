package htmltopdfSM

import (
	"fmt"
	"roofix/pkg/enum"
	"roofix/template"
)

type InputRenderHtml struct {
	Tpl   template.Name          `json:"tpl" validate:"required"`
	Title string                 `json:"title" validate:"required"`
	Data  map[string]interface{} `json:"data" validate:"required"`
	Dest  Destination            `json:"dest" validate:"required"`
}

type Destination struct {
	Bucket   string          `json:"bucket" validate:"required"`
	Folder   enum.DocFolder  `json:"folder" validate:"required"`
	Dir      string          `json:"dir"`
	Section  enum.DocSection `json:"section" validate:"required"`
	FileName string          `json:"fileName" validate:"required"`
}

type OutputRenderHtml struct {
	Bucket string      `json:"bucket" validate:"required"`
	Key    string      `json:"key" validate:"required"`
	Dest   Destination `json:"dest" validate:"required"`
	Html   string      `json:"html" validate:"required"`
}

type OutputRenderPDF struct {
	Bucket      string          `json:"bucket"`
	Key         string          `json:"Key"`
	Folder      enum.DocFolder  `json:"folder"`
	Dir         string          `json:"dir"`
	Section     enum.DocSection `json:"section"`
	FileName    string          `json:"fileName"`
	ContentType string          `json:"contentType"`
	ContentSize int64           `json:"contentSize"`
}

func (d *Destination) ObjectKey() string {
	dir := d.Folder.String()
	if d.Dir != "" {
		dir = fmt.Sprintf("%s/%s", dir, d.Dir)
	}
	return fmt.Sprintf("%s/%s", dir, d.FileName)
}
