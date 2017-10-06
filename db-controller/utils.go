package db

import "github.com/satori/go.uuid"

func GetUUID() string {
	return uuid.NewV4().String()
}

func PostToMiniPost(post *Post) (*PostMinified, error) {
	mPost := new(PostMinified)

	mPost.Id = post.Id
	mPost.Location = post.Location
	mPost.Title = post.Title
	mPost.Stock = post.Stock

	return mPost, nil
}
