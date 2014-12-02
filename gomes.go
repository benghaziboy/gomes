package gomes

import (
	"fmt"
	"log"
)

func Ping() {
	fmt.Println("PING")
}

func init() {
	auth, err := authenticateAws()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(auth)

	snsConn, err := connectSNS(&auth)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(snsConn)

}
