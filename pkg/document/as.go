package document

import (
	"roofix/ent"
)

func AsShortInfo(doc *Info) *InfoShort {
	if doc == nil {
		return nil
	}

	return &InfoShort{
		ID:          doc.ID,
		Key:         doc.Key,
		Folder:      doc.Folder,
		Name:        doc.Filename,
		ContentType: doc.ContentType,
		ContentSize: doc.ContentSize,
		Ready:       doc.Ready,
		CreatedAt:   doc.CreatedAt,
	}
}

func AsDocShortInfo(doc *ent.Document) *InfoShort {
	if doc == nil {
		return nil
	}

	return &InfoShort{
		ID:          doc.ID,
		Key:         doc.Key,
		Folder:      doc.Folder,
		Name:        doc.Filename,
		ContentType: doc.ContentType,
		ContentSize: doc.ContentSize,
		Ready:       doc.Ready,
		CreatedAt:   doc.CreatedAt,
	}
}

func AsDocInfo(doc *ent.Document) *Info {
	if doc == nil {
		return nil
	}

	return &Info{
		ID:          doc.ID,
		Key:         doc.Key,
		Folder:      doc.Folder,
		Section:     doc.Section,
		Name:        doc.Name,
		Filename:    doc.Filename,
		ContentType: doc.ContentType,
		ContentSize: doc.ContentSize,
		Ready:       doc.Ready,
		CreatedAt:   doc.CreatedAt,
		UpdatedAt:   doc.UpdatedAt,
	}
}
