package gomes

import (
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sns"
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

func initAws() error {
	err := authenticateAws()
	if err != nil {
		return err
	}

	err = connectSNS(auth)
	if err != nil {
		return err
	}

	return nil
}
