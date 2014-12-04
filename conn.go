package gomes

import (
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sns"
	"log"
)

var (
	auth    *aws.Auth
	snsConn *sns.SNS
)

func authenticateAws() error {
	a, err := aws.EnvAuth()
	if err != nil {
		return err
	}

	auth = &a

	return nil
}

func connectSNS(auth *aws.Auth) error {
	s, err := sns.New(*auth, aws.USEast)
	if err != nil {
		return err
	}

	snsConn = s

	return nil
}

func init() {
	err := authenticateAws()
	if err != nil {
		log.Println(err)
		return
	}

	err = connectSNS(auth)
	if err != nil {
		log.Println(err)
		return
	}
}
