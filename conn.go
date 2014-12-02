package gomes

import (
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sns"
	"log"
)

func authenticateAws() (auth aws.Auth, err error) {
	auth, err = aws.EnvAuth()
	return
}

func connectSNS(auth *aws.Auth) (snsConn *sns.SNS, err error) {
	snsConn, err = sns.New(*auth, aws.APNortheast)
	if err != nil {
		log.Println(err)
	}
	return
}
