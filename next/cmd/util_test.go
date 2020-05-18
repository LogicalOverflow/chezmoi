package cmd

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeCleanAbsSlashPath(t *testing.T) {
	type testCase struct {
		dir      string
		file     string
		expected string
	}
	testCases := []testCase{
		{
			dir:      "/home/user",
			file:     "file",
			expected: "/home/user/file",
		},
		{
			dir:      "/home/user",
			file:     "/tmp/file",
			expected: "/tmp/file",
		},
	}
	if runtime.GOOS == "windows" {
		testCases = append(testCases,
			testCase{
				dir:      `C:\Users\user`,
				file:     "file",
				expected: `C:/Users/user/file`,
			},
			testCase{
				dir:      `C:\Users\user`,
				file:     `dir/file`,
				expected: `C:/Users/user/dir/file`,
			},
			testCase{
				dir:      `C:\Users\user`,
				file:     `D:\Users\user\file`,
				expected: `D:/Users/user/file`,
			},
		)
	}
	for _, tc := range testCases {
		t.Run(strings.Join([]string{tc.dir, tc.file}, "_"), func(t *testing.T) {
			assert.Equal(t, tc.expected, makeCleanAbsSlashPath(tc.dir, tc.file))
		})
	}
}

func TestUpperSnakeCaseToCamelCaseMap(t *testing.T) {
	actual := upperSnakeCaseToCamelCaseMap(map[string]string{
		"BUG_REPORT_URL": "",
		"ID":             "",
	})
	assert.Equal(t, map[string]string{
		"bugReportURL": "",
		"id":           "",
	}, actual)
}
