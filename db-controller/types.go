package db

import "time"

type Location struct {
	Longitude string `json:"longitude" form:"longitude"`
	Latitude string `json:"latitude" form:"latitude"`
}

type Company struct {
	Id string `json:"id" storm:"id"`
	CreatedAt time.Time `json:"created_at" storm:"index"`

	Name string `json:"name"`
	Categories []string `json:"categories"`

	PostsCount int `json:"posts_count"`

	ActivePosts []string `json:"active_posts"`
	Subscribes []string `json:"registered"`


}

type Post struct {
	Id string `json:"id" storm:"id"`

	By string `json:"by" form:"by"`

	CreatedAt time.Time `json:"created_at" storm:"index"`

	Title string `json:"title" fako:"title" form:"title"`
	Image string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
	Address string `json:"address" fako:"street_address" form:"address"`

	//Promotion string `json:"promotion"`

	Stock int `json:"stock" form:"stock"`


}

type User struct {
	Id string `json:"id" storm:"id"`

	Group string `json:"group"`
	CreatedAt time.Time `json:"created_at" storm:"index"`


	FullName string `json:"full_name" fako:"full_name"`
	Age int `json:"age"`
	Gender string `json:"gender" fako:"gender"`

	LoginType string `json:"login_type"`

	Email string `json:"email" fako:"email_address"`

	Username string `json:"username" fako:"username"`
	Password string `json:"password" fako:"simple_password"`

	FacebookAccount string `json:"facebook_account"`

	LastLocation string `json:"last_location"`

	FollowPosts []string `json:"follow_posts"`


}

type Thumbnail struct {
	Id string `json:"id" storm:"id"`

	Name string `json:"name"`
	LocalPath string `json:"local_path"`

	Size int32 `json:"size"`

	CreatedAt time.Time `json:"created_at" storm:"index"`
}
