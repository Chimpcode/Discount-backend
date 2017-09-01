package storage


type Thumbnail struct {
	Id int32 `storm:"id,increment"`
	UUID string
	Name string
	LocalPath string
	Size int32
}
