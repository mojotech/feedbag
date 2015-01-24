package main

import "log"

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
	TimeStamp
}

type UserList []User

func (u *UserList) List() error {
	_, err := dbmap.Select(u, "SELECT * FROM USERS ORDER BY CreatedAt DESC")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (u *User) Create() error {
	// run the UpdateTime ethod on the user model
	u.UpdateTime()

	// run the DB insert function
	err := dbmap.Insert(u)
	if err != nil {
		return err
	}

	return nil
}
