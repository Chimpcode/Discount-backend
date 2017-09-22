package db

func SavePost(post *Post) (string, error) {
	posts := Db.From("posts")
	err := posts.Save(post)
	if err != nil {
		return "", err
	}
	return post.Id, nil
}

func GetPostById(id string) (*Post, error) {
	posts := Db.From("posts")
	var post Post
	err := posts.One("id", id, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}