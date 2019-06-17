package tlib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		return nil
	}
	return fmt.Errorf("Problem in pkg.Mkdir. Could not create: %s\n", path)
}

func Rmdir(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		os.RemoveAll(path)
	}
}

func ConstructDir() func() {

	old, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get current dir: %s\n", err)
	}
	subDir := "junk"
	mockdir := filepath.Join("../test-fixtures", subDir)
	err = Mkdir(mockdir)
	if err != nil {
		log.Fatalf("ConstructDir Failed: %s\n", err)
	}
	os.Chdir(mockdir)

	return func() {
		os.Chdir(old)
		c, _ := os.Getwd()
		fmt.Printf("current: %s\n", c)

		err := os.Chdir("../test-fixtures")
		if err != nil {
			log.Fatalf("can't cd")
		}

		Rmdir(subDir)
		os.Chdir(old)

	}
}
