package gomes

import (
	"bytes"
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
}

type ApsBody struct {
	Alert string  `json:"alert,omitempty"`
	Badge *string `json:"badge,omitempty"`
	Sound *string `json:"sound,omitempty"`
}

func (pt *PushToken) SendMessage(alert string, badge, sound *string) (*sns.PublishResponse, error) {
	body := &bytes.Buffer{}

	if pt.ArnType == "APNS" || pt.ArnType == "APNS_SANDBOX" {
		err := json.NewEncoder(body).Encode(map[string]ApsBody{
			"aps": ApsBody{
				Alert: alert,
				Badge: badge,
				Sound: sound,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	payload, err := json.Marshal(map[string]string{
		pt.ArnType: string(body.String()),
	})
	if err != nil {
		return nil, err
	}

	options := sns.PublishOptions{
		Message:          string(payload),
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
