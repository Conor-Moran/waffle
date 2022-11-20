package utils

import (
	"fmt"
	"io/fs"
	"log"
)

func IfErr(err error, thenDo func(err error)) {
	if err != nil {
		thenDo(err)
	}
}

func IfErrLogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FilePath(dirPath string, entry fs.DirEntry) string {
	return fmt.Sprintf("%s/%s", dirPath, entry.Name())
}
