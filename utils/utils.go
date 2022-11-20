package utils

import "log"

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
