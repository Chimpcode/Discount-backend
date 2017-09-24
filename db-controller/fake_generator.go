package db

import (
	"github.com/wawandco/fako"
	"../storage-engine"
	"time"
	"math/rand"
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

func GetFakePost(boundA, boundB Location) *Post {
	imageFake := "placeholder.jpeg"
	var post Post

	post.Id = storage.GetUUID()
	post.CreatedAt = time.Now()
	post.Image = imageFake
	post.Location = *GetRandomLocation(boundA, boundB)

	fako.Fill(&post)

	return &post
}

func GetRandomLocation(a, b Location) *Location {
	//-12.007243, -77.109899
	//-12.087220, -76.959972
	var x float64

	if a.Longitude > 0 {
		x = a.Longitude-((a.Longitude+b.Longitude)*rand.Float64())
	}else{
		x = ((a.Longitude+b.Longitude)*rand.Float64()) - a.Longitude
	}

	var y float64

	if b.Latitude > 0 {
		y = b.Latitude-((b.Latitude+a.Latitude)*rand.Float64())
	}else{
		y = ((b.Latitude+a.Latitude)*rand.Float64()) - b.Latitude
	}


	var location Location

	location.Longitude = x
	location.Latitude = y

	return &location

}
