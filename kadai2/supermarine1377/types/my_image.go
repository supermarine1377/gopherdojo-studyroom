/*
Package types implements Myimage type, which is used by other packages to represent a image file.
*/

package types

type File interface {
	Get()
}

type Myimage struct {
	FileName  string
	Path      string
	Extension string
}

type Fakeimage struct {
	FileName  string
	Path      string
	Extension string
}

func (mi Myimage) Get() {

}

func (fi Fakeimage) Get() {

}
