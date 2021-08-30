package pkg

import (
	"io/ioutil"
	"path"
)

type Video struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Format string `json:"format"`
}

func Videos(name string, format string, regex string) []Video {
	files, _ := ioutil.ReadDir("videos")
	if files == nil {
		return nil
	}
	var videos []Video

	for index, file := range files {
		fileName := file.Name()
		video := Video{
			ID:     index,
			Name:   path.Base(fileName),
			Size:   file.Size(),
			Format: path.Ext(file.Name()),
		}
		videos = append(videos, video)
	}
	return videos
}