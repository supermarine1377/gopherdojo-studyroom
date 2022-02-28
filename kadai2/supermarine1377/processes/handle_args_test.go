package processes_test

import (
	"errors"
	"supermarine1377/myio"
	"supermarine1377/processes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleArgs(t *testing.T) {
	var testcases = []struct {
		name string
		in   []string
		out  struct {
			result myio.Context
			err    error
		}
	}{
		{
			"zero arguments",
			[]string{"./main"},
			struct {
				result myio.Context
				err    error
			}{
				myio.Context{DirName: "", Extension: ""},
				errors.New("no arguments passed"),
			},
		},
		{
			"one argument",
			[]string{"./main", "images"},
			struct {
				result myio.Context
				err    error
			}{
				myio.Context{DirName: "images", Extension: "png"},
				nil,
			},
		},
		{
			"two arguments",
			[]string{"./main", "images", "webp"},
			struct {
				result myio.Context
				err    error
			}{
				myio.Context{DirName: "images", Extension: "webp"},
				nil,
			},
		},
		{
			"three arguments",
			[]string{"./main", "images", "webp", "hoge"},
			struct {
				result myio.Context
				err    error
			}{
				myio.Context{DirName: "images", Extension: "webp"},
				nil,
			},
		},
		{
			"include invalid argument",
			[]string{"./main", " "},
			struct {
				result myio.Context
				err    error
			}{
				myio.Context{DirName: "", Extension: ""},
				errors.New("dir name is empty string"),
			},
		},
	}

	asserts := assert.New(t)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := processes.HandleArgs(tc.in)
			var isFailed bool
			if !asserts.Equal(tc.out.result, result) {
				isFailed = true
			}
			if err != nil && !asserts.EqualError(tc.out.err, err.Error()) {
				isFailed = true
			}
			if isFailed {
				t.Errorf("Test Failed, %s", tc.name)
			}
		})
	}
}
