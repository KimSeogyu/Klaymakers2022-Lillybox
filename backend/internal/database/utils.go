package database

import "log"

var pl = log.Println

// FatalWithError ...
func FatalWithError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
