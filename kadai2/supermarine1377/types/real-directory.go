package types

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"supermarine1377/myio"
)

type RealDirectory struct {
	Name string
}

func (rd RealDirectory) Read() ([]myio.Myimage, error) {
	var images []myio.Myimage
	err := filepath.WalkDir(rd.Name, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		var image myio.Myimage
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

func (rd RealDirectory) Write(mi []myio.Myimage, context myio.Context) error {
	for _, image := range mi {
		if err := convert(image, context.Extension); err != nil {
			log.Println(err)
			return err
		}
		log.Printf("finished converting %s", image.FileName)
	}
	return nil
}

// convert a Myimage (see types package) passed as a argumaent.
func convert(arg myio.Myimage, extension string) error {
	log.Printf("started to convert %s to %s", arg.FileName, extension)

	if arg.Extension != "jpg" {
		err := fmt.Errorf("%s is not .jpg", arg.FileName)
		return err
	}
	reader, err := myopen(arg.FileName, arg.Path)
	if err != nil {
		return err
	}

	m, err := mydecode(arg.FileName, reader)
	if err != nil {
		return err
	}

	if err := myencode(m, arg.FileName, extension); err != nil {
		return err
	}

	return nil
}

func myopen(name string, path string) (*os.File, error) {
	log.Printf("opening %s", name)
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return reader, err
}

func mydecode(name string, reader *os.File) (image.Image, error) {
	log.Printf("decoding %s", name)
	m, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func myencode(m image.Image, name string, extension string) error {
	log.Printf("encoding %s to %s", name, extension)
	new, err := newname(name, extension)
	if err != nil {
		return err
	}
	c, err := os.Create(new)
	defer c.Close()
	if err != nil {
		return err
	}
	jpeg.Encode(c, m, &jpeg.Options{Quality: 100})

	return nil
}

func newname(old string, extenstion string) (string, error) {
	log.Printf("renaming %s", old)
	var result string
	var position = strings.LastIndex(old, ".")
	if position == -1 {
		err := errors.New("This name has no dot")
		return result, err
	}
	result = old[0:position] + "." + extenstion
	return result, nil
}
