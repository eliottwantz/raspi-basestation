package internal

import "log"

func FatalError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func NotFatalError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
