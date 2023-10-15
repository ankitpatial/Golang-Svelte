/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package document

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/ent"
	"roofix/ent/document"
	"roofix/pkg/enum"
	"roofix/pkg/util/str"
)

// ByFolderDirSection entries in DB (folderKey is MD5 of  DocFolder/<someID | app-level-data >)
func ByFolderDirSection(ctx context.Context, folder enum.DocFolder, dir *string, section *enum.DocSection) (res []*Info, err error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Document.
		Query().
		Where(document.Ready(true), document.FolderEQ(enum.DocFolder(folder)))

	if dir != nil {
		qry.Where(document.Dir(*dir))
	}

	if section != nil {
		qry.Where(document.SectionEQ(enum.DocSection(section.String())))
	}

	err = qry.Where(func(d *sql.Selector) {
		d.Select(
			d.C(document.FieldID),
			d.C(document.FieldKey),
			d.C(document.FieldSection),
			d.C(document.FieldName),
			d.C(document.FieldReady),
			fmt.Sprintf("%s as filename", d.C(document.FieldFilename)),
			fmt.Sprintf("%s as contentSize", d.C(document.FieldContentSize)),
			fmt.Sprintf("%s as contentSize", d.C(document.FieldContentSize)),
			fmt.Sprintf("%s as createdAt", d.C(document.FieldCreatedAt)),
			fmt.Sprintf("%s as updatedAt", d.C(document.FieldUpdatedAt)),
		).OrderBy(sql.Asc(d.C(document.FieldCreatedAt)))
	}).
		Select().
		Scan(ctx, &res)

	return res, err
}

func Exists(ctx context.Context, folder enum.DocFolder, dir *string, section *enum.DocSection, fileName *string) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := docQry(db.Document, folder, dir, section, fileName)
	return qry.Exist(ctx)
}

func Get(ctx context.Context, folder enum.DocFolder, dir *string, section *enum.DocSection, fileName *string) (*ent.Document, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := docQry(db.Document, folder, dir, section, fileName)
	return qry.Select(document.FieldBucket, document.FieldKey).First(ctx)
}

func docQry(cl *ent.DocumentClient, folder enum.DocFolder, dir *string, section *enum.DocSection, fileName *string) *ent.DocumentQuery {
	q := cl.Query().Where(document.Ready(true), document.FolderEQ(folder))

	// dir check
	d := str.Val(dir)
	if d != "" {
		q.Where(document.Dir(d))
	}
	// section check
	if section != nil {
		q.Where(document.SectionEQ(*section))
	}
	// filename check
	n := str.Val(fileName)
	if n != "" {
		q.Where(document.NameEqualFold(n))
	}

	return q
}
