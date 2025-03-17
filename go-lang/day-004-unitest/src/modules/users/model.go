package users

import (
	z "github.com/Oudwins/zog"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var userSchema = z.Struct(z.Schema{
	"ID":    z.Int().Required(),
	"Name":  z.String().Required(),
	"Email": z.String().Email().Required(),
	"Age":   z.Int().Required(),
})

func (u *User) Validate() z.ZogIssueMap {
	return userSchema.Validate(u)
}

var users = make(map[string]User)
