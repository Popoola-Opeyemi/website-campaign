package utils

import (
	"os"

	"bitbucket.org/mayowa/helpers/path"
	"github.com/labstack/gommon/log"
)

// directorys ...
var directorys = ([]string{"static/uploads", "tmp"})

// InitBoot function to create all the folders needed ...
func InitBoot() {

	for k, pth := range directorys {
		if !path.Available(pth) {
			log.Errorf("path[%s]: %s not accessible", directorys[k], pth)
			log.Infof("Creating Folder %s", pth)
			if err := os.MkdirAll(pth, 0700); err != nil {
				log.Errorf("path[%s]: %s cannot create folder", k, pth)
			}
		} else {
			log.Infof("path[%s]: %s", directorys[k], pth)
		}
	}
}

// GetPath returns the upload path ...
func GetPath() map[string]string {
	path := make(map[string]string)
	path["upload"] = directorys[0]
	return path
}

// BaseURL returns the baseurl path ...
func BaseURL() map[string]string {
	url := make(map[string]string)
	url["base_url"] = "/static/uploads"
	return url
}
