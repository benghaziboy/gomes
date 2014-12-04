package gomes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crowdmob/goamz/sns"
	"os"
	"strconv"
)

var (
	MissingEnabledAttribute = errors.New("Missing 'Enabled' attribute from the aws response")
	Apn                     = os.Getenv("SNS_APN_ARN")
)

type PushToken struct {
	Uid     string
	Arn     string
	ArnType string
	Token   string
	Sandbox bool
}

type PushBody struct {
	Default     string   `json:"default"`
	Apns        ApnsBody `json:"APNS,omitempty"`
	ApnsSandbox ApnsBody `json:"APNS_SANDBOX,omitempty"`
}

type ApnsBody struct {
	Aps ApsBody `json:"aps,omitempty"`
}

type ApsBody struct {
	Alert string  `json:"alert,omitempty"`
	Badge *string `json:"badge,omitempty"`
	Sound *string `json:"sound,omitempty"`
}

func (pt *PushToken) SendMessage(alert string, badge, sound *string) (*sns.PublishResponse, error) {
	rawBody := PushBody{Default: alert}

	if pt.Arn != "" {
		if pt.Sandbox == true {
			rawBody.ApnsSandbox = ApnsBody{
				Aps: ApsBody{
					Alert: alert,
					Badge: badge,
					Sound: sound,
				},
			}
		}
	}

	b, err := json.Marshal(&rawBody)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	options := sns.PublishOptions{
		Message:          string(b),
		MessageStructure: "json",
		TargetArn:        pt.Arn,
	}

	pr, err := snsConn.Publish(&options)
	return pr, err
}

func (pt *PushToken) IsEnabled() (bool, error) {
	resp, err := snsConn.GetEndpointAttributes(pt.Arn)
	if err != nil {
		return false, err
	}

	for _, v := range resp.Attributes {
		if v.Key == "Enabled" {
			enabled, err := strconv.ParseBool(v.Value)
			return enabled, err
		}
	}

	return false, MissingEnabledAttribute
}

func New(uid, token string) (*PushToken, error) {
	arn, err := NewArn(token)
	if err != nil {
		return nil, err
	}

	pushtoken := PushToken{
		Uid:   uid,
		Arn:   *arn,
		Token: token,
	}
	fmt.Println(pushtoken)

	return &pushtoken, nil
}

func NewArn(token string) (*string, error) {
	opts := sns.PlatformEndpointOptions{
		PlatformApplicationArn: Apn,
		Token: token,
	}

	resp, err := snsConn.CreatePlatformEndpoint(&opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &resp.EndpointArn, nil
}
