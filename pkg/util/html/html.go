package html

import (
	"bytes"
	"html/template"
	"roofix/config"
)

func ExecTpl(layout string, content string, data map[string]interface{}) (string, error) {
	/*
	 *  template name are those defined in .html file
	 * and these names plays a role template rendering
	 */
	tpl, err := template.New("layout").Parse(layout)
	if err != nil {
		return "", err
	}

	_, err = tpl.New("message").Funcs(template.FuncMap{"html": toHtml}).Parse(content)
	if err != nil {
		return "", err
	}

	data["assetURL"] = config.Read().Website.AssetUrl
	data["siteURL"] = config.Read().Website.Url
	data["siteName"] = config.Read().Website.Domain()

	buf := new(bytes.Buffer)
	err = tpl.ExecuteTemplate(buf, "layout", data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func toHtml(s string) template.HTML {
	return template.HTML(s)
}
