package gomes

import (
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sns"
)

var (
	awsAuth *aws.Auth
	awsSns  *sns.SNS
)

func authenticateAws() error {
	a, err := aws.EnvAuth()
	if err != nil {
		return err
	}

	awsAuth = &a

	return nil
}

func connectSNS(auth *aws.Auth) error {
	s, err := sns.New(*auth, aws.USEast)
	if err != nil {
		return err
	}

	awsSns = s

	return nil
}

func initAws() error {
	err := authenticateAws()
	if err != nil {
		return err
	}

	err = connectSNS(awsAuth)
	if err != nil {
		return err
	}

	return nil
}
