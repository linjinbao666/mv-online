package pkg

import (
	"io/ioutil"
	"os"
	"path"
)

type Video struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Format string `json:"format"`
	Source string `json:"source"`
}

func VideoDelete(name string, format string, regex string, workingDir string) (string, error) {
	err := os.Remove(workingDir + "/videos/" + name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func Videos(name string, format string, regex string, workingDir string) []Video {
	files, _ := ioutil.ReadDir(workingDir + "/videos")
	if files == nil {
		return nil
	}
	var videos []Video

	for index, file := range files {
		fileName := file.Name()
		video := Video{
			ID:     index,
			Name:   path.Base(fileName),
			Size:   file.Size() >> 20,
			Format: path.Ext(file.Name()),
			Source: "本地",
		}
		videos = append(videos, video)
	}
	return videos
}
