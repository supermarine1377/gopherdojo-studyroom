package types

import (
	"io/fs"
	"supermarine1377/myio"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDirEntry struct {
	name string
}

func (m mockDirEntry) Name() string {
	return m.name
}

func (m mockDirEntry) IsDir() bool {
	return false
}

func (m mockDirEntry) Type() fs.FileMode {
	return 0
}

func (m mockDirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

func TestGetImage(t *testing.T) {
	var testcases = []struct {
		name string
		in   struct {
			d    fs.DirEntry
			path string
		}
		out struct {
			expected myio.Myimage
			err      error
		}
	}{
		{
			name: "1st",
			in: struct {
				d    fs.DirEntry
				path string
			}{
				d:    mockDirEntry{name: "test.jpg"},
				path: ".",
			},
			out: struct {
				expected myio.Myimage
				err      error
			}{
				expected: myio.Myimage{
					FileName:  "test.jpg",
					Path:      ".",
					Extension: "jpg",
				},
				err: nil,
			},
		},
		{
			name: "2st",
			in: struct {
				d    fs.DirEntry
				path string
			}{
				d:    mockDirEntry{name: "test.png"},
				path: ".",
			},
			out: struct {
				expected myio.Myimage
				err      error
			}{
				expected: myio.Myimage{
					FileName:  "",
					Path:      "",
					Extension: "",
				},
				err: nil,
			},
		},
	}
	assert := assert.New(t)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getImage(tc.in.d, tc.in.path)
			assert.Equal(tc.out.expected, result)
			if err != nil {
				assert.EqualError(tc.out.err, err.Error())
			}
		})
	}
}

func TestNewName(t *testing.T) {
	var testcases = []struct {
		name string
		in   struct {
			old       string
			extension string
		}
		out struct {
			result string
			err    error
		}
	}{
		{
			name: "1st",
			in: struct {
				old       string
				extension string
			}{
				old:       "test.jpg",
				extension: "png",
			},
			out: struct {
				result string
				err    error
			}{
				result: "test.png",
				err:    nil,
			},
		},
	}
	assert := assert.New(t)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := newname(tc.in.old, tc.in.extension)
			assert.Equal(tc.out.result, result)
			if err != nil {
				assert.EqualError(tc.out.err, err.Error())
			}
		})
	}
}
