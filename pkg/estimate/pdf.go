package estimate

import (
	"context"
	"fmt"
	"roofix/bin/htmltopdfSM"
	"roofix/config"
	"roofix/ent"
	"roofix/pkg/document"
	"roofix/pkg/enum"
	"roofix/pkg/nearmap"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"roofix/pkg/util/parse"
	"roofix/template"
)

type pdfTplData struct {
	Company  string `json:"company"`
	Address  string `json:"address"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Area     string `json:"area"`
	Pitch    string `json:"pitch"`
	Material string `json:"material"`
	Redeck   bool   `json:"redeck"`
	Amount   string `json:"amount"`
}

func genPdf(ctx context.Context, estID string, inp *Input, p *nearmap.Price) {
	// document destination info
	dest := htmltopdfSM.Destination{
		Bucket:   config.DataBucket(),
		Folder:   enum.FolderEstimates,
		Dir:      estID,
		Section:  enum.SectionPricingPDF,
		FileName: "Estimate.pdf",
	}
	// pdf template data
	data := pdfTplData{
		Company:  inp.SalesRep.CompanyName,
		Address:  fmt.Sprintf("%s, %s", inp.HomeOwner.StreetNumber, inp.HomeOwner.StreetName),
		City:     inp.HomeOwner.City,
		State:    inp.HomeOwner.State,
		Zip:      inp.HomeOwner.Zip,
		Phone:    inp.SalesRep.Mobile,
		Email:    inp.SalesRep.Email,
		Area:     fmt.Sprintf("%.2f SQ", p.TotalSq),
		Pitch:    fmt.Sprintf("%.0f/12", p.Pitch),
		Material: p.Material,
		Redeck:   p.IsRedeck,
		Amount:   num.FormatMoney(p.Amt),
	}

	if !config.IsDevEnv() {
		htmltopdfSM.Exec(ctx, dest, template.EstimatePDF, "Estimate", parse.StructToMap(data))
		// create document
		doc, err := document.Save(ctx, dest.Bucket, dest.ObjectKey(), &document.Input{
			Folder:   dest.Folder,
			Dir:      dest.Dir,
			Section:  dest.Section,
			Name:     dest.FileName,
			FileName: dest.FileName,
		}, false)
		if err != nil {
			log.Error(err)
			return
		}

		// link document to estimate
		if err = ent.GetClient().Estimate.UpdateOneID(estID).SetPdfID(doc.ID).Exec(ctx); err != nil {
			log.Error(err)
		}
	}
}
