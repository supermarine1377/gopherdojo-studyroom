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
		if args[1] == "" {
			err := errors.New("dir name is empty string")
			return context, err
		}
		context.DirName = args[1]
		context.Extension = "png"
		return context, nil
	}
	context.DirName = args[1]
	context.Extension = args[2]
	return context, nil
}
