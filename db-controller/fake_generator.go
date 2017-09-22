package db

import (
	"github.com/wawandco/fako"
	"../storage-engine"
	"time"
)

func GetFakeCompany() *Company{
	var company Company

	company.Id = storage.GetUUID()
	company.CreatedAt = time.Now()
	company.Name = "Generic cupoint company"

	fako.Fill(&company)
	return &company
}

func GetFakeUsers(howMany int) []User {
	users := make([]User, 0)
	for i:=0;i<howMany;i++ {
		var user User
		user.Id = storage.GetUUID()
		user.CreatedAt = time.Now()
		fako.Fill(&user)
		users = append(users, user)
	}

	return users
}

func GetFakeUser() *User {
	var user User
	user.Id = storage.GetUUID()
	user.CreatedAt = time.Now()
	fako.Fill(&user)
	return  &user
}

func GetFakePost() *Post {
	imageFake := "placeholder.jpeg"
	var post Post

	post.Id = storage.GetUUID()
	post.CreatedAt = time.Now()
	post.Image = imageFake

	fako.Fill(&post)

	return &post
}
