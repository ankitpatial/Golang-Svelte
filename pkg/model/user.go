package model

type InputUserProfile struct {
	Picture       *string `json:"picture,omitempty"`
	FirstName     *string `json:"firstName,omitempty"`
	LastName      *string `json:"lastName,omitempty"`
	Phone         *string `json:"phone,omitempty"`
	OldPwd        *string `json:"oldPwd,omitempty"`
	NewPwd        *string `json:"newPwd,omitempty"`
	ConfirmNewPwd *string `json:"confirmNewPwd,omitempty"`
}

type UpdateProfileResponse struct {
	Picture *string `json:"picture,omitempty"`
}
