package gomodulesdepone

import (
	"log"
)

// GetData : return gomodulesdepone v1.0.1.
func GetData() string {
	log.Println("fix pug && become v1.0.1")
	return "Hi from gomodulesdepone v1.0.1"
}

// GetExtraData : return gomodulesdepone v1.1.1.
func GetExtraData() string {
	return "GetExtraData from gomodulesdepone v1.1.1"
}
