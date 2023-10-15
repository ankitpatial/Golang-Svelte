package queue

import (
	"context"
	"roofix/pkg/util/log"
	"roofix/template"
	"testing"
)

type Message struct {
	To      []string
	Subject string
	Tpl     template.Name
	Modal   map[string]interface{}
}

func TestUnmarshal(t *testing.T) {
	data := "{\"ctxData\":{\"user\":{\"id\":\"6de39dbf-ac6d-4fa9-bae2-ccf62027a245\",\"firstName\":\"Ankit\",\"lastName\":\"Patial\",\"email\":\"ankit@simsaw.com\",\"role\":\"ADMIN\"},\"ip\":\"61.12.74.6\",\"apiUserID\":\"\"},\"data\":{\"To\":[\"Luke Leak \\u003cluke@jwhomesolutions.com\\u003e\"],\"Subject\":\"Job Update\",\"Tpl\":\"job-update.html\",\"Modal\":{\"jobAddress\":\"5721  Violet Crown Place, Fort Worth\\u003cbr/\\u003e Texas 76126\",\"message\":\"Estimate request is denied with reason: 'ss'\",\"name\":\"Luke Leak\",\"pathname\":\"\"}}}"
	var msg Message
	var err error
	if _, msg, err = Unmarshal[Message](context.Background(), data); err != nil {
		log.Error(err)
		return
	}

	t.Log(msg)

}
