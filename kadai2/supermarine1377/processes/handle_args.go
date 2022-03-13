package processes

import (
	"errors"
	"supermarine1377/myio"
)

// Handle arguments passed when user excuses ./main
func HandleArgs(args []string) (*myio.Context, error) {
	var context myio.Context
	if len(args) == 1 {
		err := errors.New("no arguments passed")
		return nil, err
	}
	if len(args) == 2 {
		if args[1] == "" || args[1] == " " {
			err := errors.New("dir name is empty string")
			return nil, err
		}
		context.DirName = args[1]
		context.Extension = "png"
		return &context, nil
	}
	context.DirName = args[1]
	context.Extension = args[2]
	return &context, nil
}
