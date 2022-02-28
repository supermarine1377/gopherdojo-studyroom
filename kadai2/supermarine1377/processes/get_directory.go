package processes

import (
	"io/ioutil"
	"supermarine1377/myio"
	"supermarine1377/types"
)

func GetDirectory(context myio.Context) (types.RealDirectory, error) {
	var rd types.RealDirectory
	if _, err := ioutil.ReadDir(context.DirName); err != nil {
		return rd, err
	}
	rd.Name = context.DirName
	return rd, nil
}
