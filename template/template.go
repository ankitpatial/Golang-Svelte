package template

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"roofix/config"
	"roofix/pkg/enum"
	"roofix/pkg/util/html"
	"roofix/pkg/util/log"
	"roofix/pkg/util/storage"
	"strings"
	"sync"
)

const (
	KindEmail Kind = "email"
	KindPDF   Kind = "pdf"

	EmailContactUs             Name = "contact-us.html"
	EmailEagleViewEstChange    Name = "eagle-est-change.html"
	EmailEagleViewEstNeed      Name = "eagle-est-need.html"
	EmailNewJobAssigned        Name = "job-assigned.html"
	EmailJobWillUnassign       Name = "job-will-unassign.html"
	EmailJobUnassigned         Name = "job-unassigned.html"
	EmailJobUnassignedAdmin    Name = "job-unassigned-admin.html"
	EmailNewUserAccount        Name = "new-user-account.html"
	EmailPartnerOnBoarding     Name = "partner-on-boarding.html"
	EmailPartnerOnBoardingDone Name = "partner-on-boarding-done.html"
	EmailPartnerNewUser        Name = "partner-new-user.html"
	EmailSetPassword           Name = "set-password.html"

	EstimatePDF Name = "estimate.html"
)

type Name string

func (e Name) String() string {
	return string(e)
}

type Kind string

func (e Kind) String() string {
	return string(e)
}

func Exec(ctx context.Context, kind Kind, name Name, data map[string]interface{}) (string, error) {
	// get html template
	var layout, content string
	if config.IsDevEnv() {
		wd, _ := os.Getwd()
		mod := fmt.Sprintf("%s%s", string(os.PathSeparator), config.AppName)
		idx := strings.Index(wd, mod)
		fileLayout := filepath.Join(wd[:idx], config.AppName, "template", kind.String(), "_layout.html")
		fileTpl := filepath.Join(wd[:idx], config.AppName, "template", kind.String(), name.String())

		l, err := os.ReadFile(fileLayout)
		if err != nil {
			return "", err
		}
		layout = string(l)

		c, err := os.ReadFile(fileTpl)
		if err != nil {
			return "", err
		}
		content = string(c)
	} else {
		var wg sync.WaitGroup
		var dir string
		switch kind {
		case KindEmail:
			dir = enum.DirEmailTemplates.String()
		case KindPDF:
			dir = enum.DirPdfTemplates.String()
		}

		wg.Add(2)
		go func() {
			defer wg.Done()
			d, err := storage.GetObject(ctx, config.DataBucket(), fmt.Sprintf("%s/%s", dir, "_layout.html"))
			if err != nil {
				log.Error(err)
				return
			}
			layout = string(d)
		}()
		go func() {
			defer wg.Done()
			d, err := storage.GetObject(ctx, config.DataBucket(), fmt.Sprintf("%s/%s", dir, name.String()))
			if err != nil {
				log.Error(err)
				return
			}
			content = string(d)
		}()
		wg.Wait()

		if layout == "" {
			return "", errors.New("failed to get _layout.html")
		}

		if content == "" {
			return "", errors.New("failed to get " + name.String())
		}
	}

	return html.ExecTpl(layout, fmt.Sprintf("{{define \"message\"}}%s{{end}}", content), data)
}
