package gomes

import (
	"fmt"
	"github.com/crowdmob/goamz/sns"
	"os"
)

var (
	APN = os.Getenv("SNS_APN_ARN")
)

type PushToken struct {
	uid   string
	arn   string
	token string
}

func (pt *PushToken) IsEnabled() (bool, error) {
	resp, err := snsConn.GetEndpointAttributes(pt.arn)
	if err != nil {
		return false, err
	}

	fmt.Println(resp)

	return true, nil
}

// func New(uid string) (*PushToken, error) {

// }

func NewArn(token string) (*sns.CreatePlatformEndpointResponse, error) {
	opts := sns.PlatformEndpointOptions{
		PlatformApplicationArn: APN,
		Token: token,
	}

	resp, err := snsConn.CreatePlatformEndpoint(&opts)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return resp, nil
}

func PrintAuth() {
	fmt.Println(auth)
	fmt.Println(snsConn)
}
