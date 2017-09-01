package db

import (
	"../storage-engine"
	"github.com/boltdb/bolt"
	"fmt"
)

func AddImagesPointerToPath(name string, size int32) error {
	Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("images"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	img := storage.Thumbnail{UUID: storage.GetUUID(), Name: name, LocalPath: "images/"+name, Size: size}
	err := Db.Set("images",img.UUID, img)
	if err != nil {
		return err
	}
	return nil
}