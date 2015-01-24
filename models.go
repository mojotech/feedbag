package main

type User struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Username    string      `json:"username"`
	AvatarUrl   string      `json:"avatar_url"`
	ProfileUrl  string      `json:"profile_url"`
	Email       string      `json:"email"`
	Joined      string      `json:"joined"`
	Raw         interface{} `json:"raw"`
	AccessToken string      `json:"-"`
}
