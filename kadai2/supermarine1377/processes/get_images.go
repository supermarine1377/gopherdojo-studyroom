package processes

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"supermarine1377/types"
)

// implements getting image as a slice of Myimage (see types package).
func GetImages(context types.Context) ([]types.Myimage, error) {
	dirName := context.DirName
	if _, err := ioutil.ReadDir(dirName); err != nil {
		return nil, err
	}
	var images []types.Myimage
	err := filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		var image types.Myimage
		name := d.Name()

		match, err := regexp.MatchString("..+jpg", name)
		if err != nil {
			return err
		}
		if !match {
			return nil
		}
		image.FileName = name
		image.Path = path
		image.Extension = "jpg"

		images = append(images, image)

		return nil
	})

	return images, err
}
