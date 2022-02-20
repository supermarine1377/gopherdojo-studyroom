/* This command convert jpg image to png, gif, tif, etc.

Command options

- Directory name

You must place jpg images in this directory. All jpg image in this direcotory is to be converted.

- Extenstion

A extension name you want images to converted.

Example

./main images png

*/

package main

import (
	"log"
	"os"
	"supermarine1377/processes"
)

func main() {
	context, err := processes.HandleArgs(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	images, err := processes.GetImages(context)
	if err != nil {
		log.Println(err)
	}
	for _, image := range images {
		if err := processes.Convert(image, context.Extension); err != nil {
			log.Println(err)
		}
		log.Printf("finished converting %s", image.FileName)
	}
}
