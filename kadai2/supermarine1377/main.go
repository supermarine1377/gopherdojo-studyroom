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
	"supermarine1377/myio"
	"supermarine1377/processes"
)

func main() {
	context, err := processes.HandleArgs(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rw, err := processes.GetDirectory(context)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	images, err := myio.Read(rw)
	if err != nil {
		log.Println(err)
	}
	if err := myio.Write(rw, images, context); err != nil {
		log.Println(err)
	}
}
