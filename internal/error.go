package internal

import "log"

func FatalError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
