package models

type User struct {
	Fullname        string `json:"fullname"`
	Contactno       string `json:"contactno"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	PWord           string `json:"password"`
	OldPassword     string `json:"oldpassword"`
	NewPassword     string `json:"newpassword"`
	ConfirmPassword string `json:"confirmpassword"`
}
