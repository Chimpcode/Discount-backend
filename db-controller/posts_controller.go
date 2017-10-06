package db

import "github.com/asdine/storm/q"

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
	err := posts.One("Id", id, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetAllPosts() (map[string]Post, error) {
	posts := Db.From("posts")
	allPosts := make(map[string]Post)
	var postsArray []Post
	err := posts.All(&postsArray)
	if err != nil  {
		return nil, err
	}
	for _, post := range postsArray {
		allPosts[post.Id] = post
	}
	return allPosts, nil
}

func DeleteAllPosts() error {
	posts := Db.From("posts")
	err := posts.Select(q.True()).Delete(new(Post))
	return err
}

func DetelePostById(id string) error {
	posts := Db.From("posts")
	err:= posts.Select(q.Eq("Id", id)).Delete(new(Post))
	return err
}

func GetMiniPostById(id string) (*PostMinified, error){
	miniPost := new(PostMinified)
	post, err := GetPostById(id)
	if err != nil {
		return miniPost, err
	}
	miniPost, err = PostToMiniPost(post)

	return miniPost, err
}