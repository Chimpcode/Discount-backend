package db

type Company struct {
	Id string `json:"id"`

	Name string `json:"name"`
	Categories []string `json:"categories"`

	PostsCount int `json:"posts_count"`

	ActivePosts []string `json:"active_posts"`
	Subscribes []string `json:"registered"`
}

type Post struct {
	Id string `json:"id"`

	By string `json:"by"`

	CreatedAt string `json:"created_at"`

	Title string `json:"title"`
	Image string `json:"image"`
	Description string `json:"description"`
	Address string `json:"address"`
	Distance int `json:"distance"`

	Promotion string `json:"promotion"`

	Stock int `json:"stock"`


}

type User struct {
	Id string `json:"id"`

	Group string `json:"group"`

	FullName string `json:"full_name"`
	Age int `json:"age"`
	Sex string `json:"sex"`

	LoginType string `json:"login_type"`

	Email string `json:"email"`

	Username string `json:"username"`
	Password string `json:"password"`

	FacebookAccount string `json:"facebook_account"`

	LastLocation string `json:"last_location"`

	FollowPosts []string `json:"follow_posts"`
}

type Thumbnail struct {
	Id int32 `storm:"id,increment"`
	UUID string
	Name string
	LocalPath string
	Size int32
}
