package cmd

import "io/ioutil"

func SaveFile(dir string, bytes []byte) error {
	return ioutil.WriteFile(dir, bytes, 0644)
}
