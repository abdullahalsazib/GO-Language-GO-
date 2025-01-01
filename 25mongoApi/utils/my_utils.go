package utils

import "log"

func CheckTheError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
