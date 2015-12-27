package filemanip

import "io/ioutil"

func LoadFiles(files map[string]*[]byte) error {
	var err error

	for file, bytes := range files {
		*bytes, err = ioutil.ReadFile(file)
		if err != nil {
			break
		}
	}

	return err
}

func ReadFiles(files map[string]*string) error {
	var err error
	var bytes []byte

	for file, content := range files {
		bytes, err = ioutil.ReadFile(file)
		*content = string(bytes)
		if err != nil {
			break
		}
	}

	return err
}
