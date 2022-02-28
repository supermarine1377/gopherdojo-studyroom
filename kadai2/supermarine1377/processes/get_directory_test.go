package processes_test

import (
	"supermarine1377/myio"
	"supermarine1377/processes"
	"supermarine1377/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirectory(t *testing.T) {
	var testcases = []struct {
		name string
		in   myio.Context
		out  struct {
			types.RealDirectory
			error
		}
	}{
		{
			"1st",
			myio.Context{DirName: "../images", Extension: "png"},
			struct {
				types.RealDirectory
				error
			}{
				types.RealDirectory{Name: "images"},
				nil,
			},
		},
	}

	var assert = assert.New(t)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := processes.GetDirectory(tc.in)
			var isFailed bool
			if assert.NotEqual(tc.out.RealDirectory, result) {
				isFailed = true
			}
			if err != nil && !assert.EqualError(tc.out.error, err.Error()) {
				isFailed = true
			}
			if isFailed {
				t.Logf("Failed test case, %s", tc.name)
			}
		})
	}
}
