package model

type InputSurveyProgress struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Order    int    `json:"order"`
}
