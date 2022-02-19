package processes_test

import (
	"errors"
	"supermarine1377/processes"
	"supermarine1377/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name string
	in   []string
	out  out
}

type out struct {
	result types.Context
	err    error
}

var testcases = []testcase{
	{
		"zero arguments",
		[]string{"./main"},
		out{
			types.Context{DirName: "", Extension: ""},
			errors.New("no arguments passed"),
		},
	},
	{
		"one argument",
		[]string{"./main", "images"},
		out{
			types.Context{DirName: "images", Extension: "png"},
			nil,
		},
	},
	{
		"two arguments",
		[]string{"./main", "images", "webp"},
		out{
			types.Context{DirName: "images", Extension: "webp"},
			nil,
		},
	},
}

func TestHandleArgs(t *testing.T) {
	asserts := assert.New(t)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("test name: %s", tc.name)
			t.Logf(tc.out.result.Extension)
			result, _ := processes.HandleArgs(tc.in)
			asserts.Equal(tc.out.result, result)
		})
	}
}
