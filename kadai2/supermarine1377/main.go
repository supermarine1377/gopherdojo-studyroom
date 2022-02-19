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
	"io/ioutil"
	"log"
	"os"
	"supermarine1377/processes"
)

func main() {
	context, err := processes.HandleArgs(os.Args)
	if err != nil {
		log.Println("no arguments passed, exiting...")
		os.Exit(1)
	}
	dirName := context.DirName
	extension := context.Extension
	if _, err := ioutil.ReadDir(dirName); err != nil {
		log.Printf("no dir %s found, exiting...", dirName)
		os.Exit(1)
	}
	images, err := processes.GetImages(dirName)
	if err != nil {
		log.Println("error occured during reading images", err)
	}
	for _, image := range images {
		if err := processes.Convert(image, extension); err != nil {
			log.Println(err)
		}
		log.Printf("finished converting %s", image.FileName)
	}
}
