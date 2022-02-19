package processes

import (
	"errors"
	"supermarine1377/types"
)

func HandleArgs(args []string) (types.Context, error) {
	var context types.Context
	if len(args) == 1 {
		err := errors.New("no arguments passed")
		return context, err
	}
	if len(args) == 2 {
		dirName := args[1]
		if dirName == "" {
			err := errors.New("dir name is empty string")
			return context, err
		}
		context.DirName = dirName
		context.Extension = "png"
		return context, nil
	}
	dirName := args[1]
	extension := args[2]
	context.DirName = dirName
	context.Extension = extension
	return context, nil
}
