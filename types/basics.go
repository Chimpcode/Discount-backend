package types

import (
	_ "github.com/asdine/storm"
	"time"
)


type Offer struct {
	Id            string `storm:"id"`
	Name          string
	ThumbnailPath string
	Description   string
	Stock         int32
	Master        string
	Address		  string
	CreatedAt time.Time  `storm:"index"`
}

type User struct {
	Id       string `storm:"id"`
	Type     string
	Name     string
	Lastname string
	Email    string
	CreatedAt time.Time  `storm:"index"`

}
