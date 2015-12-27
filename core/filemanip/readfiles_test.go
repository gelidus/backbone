package filemanip

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadFiles(t *testing.T) {
	readfilesContent := []byte{}
	readfile_testContent := []byte{}

	err := LoadFiles(map[string]*[]byte{
		"readfiles.go": &readfilesContent,
		"readfiles_test.go": &readfile_testContent,
	})

	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, len(readfilesContent))
	assert.NotEqual(t, 0, len(readfile_testContent))
}

func TestReadFiles(t *testing.T) {
	readfilesContent := ""
	readfile_testContent := ""

	err := ReadFiles(map[string]*string{
		"readfiles.go": &readfilesContent,
		"readfiles_test.go": &readfile_testContent,
	})

	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, len(readfilesContent))
	assert.NotEqual(t, 0, len(readfile_testContent))
}