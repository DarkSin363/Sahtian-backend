package model

type User struct {
	ID                    int64  `json:"id" bson:"id"`
	IsAdmin               bool   `json:"is_admin" bson:"is_admin"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_menu" bson:"added_to_attachment_menu"`
	AllowsWriteToPm       bool   `json:"allows_write_to_pm" bson:"allows_write_to_pm"`
	FirstName             string `json:"first_name" bson:"first_name"`
	Visible               bool   `json:"visible" bson:"visible"`
	AvatarURL             string `json:"avatar_url" bson:"avatar_url"`
	IsBot                 bool   `json:"is_bot" bson:"is_bot"`
	IsPremium             bool   `json:"is_premium" bson:"is_premium"`
	LastName              string `json:"last_name" bson:"last_name"`
	Username              string `json:"username" bson:"username"`
	LanguageCode          string `json:"language_code" bson:"language_code"`
}

type UserParams struct {
	FirstLogin int64 `json:"first_login" bson:"first_login"`
	LastLogin  int64 `json:"last_login" bson:"last_login"`
}
