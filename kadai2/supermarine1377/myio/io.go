package myio

type Reader interface {
	Read() ([]Myimage, error)
}

type Writer interface {
	Write(mi []Myimage, context Context) error
}

type Myimage struct {
	FileName  string
	Path      string
	Extension string
}

type Context struct {
	DirName   string
	Extension string
}

func Read(r Reader) ([]Myimage, error) {
	return r.Read()
}

func Write(w Writer, mi []Myimage, c Context) error {
	return w.Write(mi, c)
}
