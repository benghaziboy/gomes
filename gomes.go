package gomes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crowdmob/goamz/sns"
	"os"
	"strconv"
	"strings"
)

var (
	MissingEnabledAttribute = errors.New("Missing 'Enabled' attribute from the aws response")
	SnsArn                  = os.Getenv("SNS_APP_ARN")
)

type PushToken struct {
	Uid     string
	Arn     string
	ArnType string
	Token   string
}

type Apns struct {
	Apns string `json:"APNS"`
}

type ApnsSandbox struct {
	ApnsSandbox string `json:"APNS_SANDBOX"`
}

type ApsType struct {
	ApsData `json:"aps,omitempty"`
}

type ApsData struct {
	Alert string  `json:"alert,omitempty"`
	Badge *string `json:"badge,omitempty"`
	Sound *string `json:"sound,omitempty"`
}

type Gcm struct {
	Gcm string `json:"GCM"`
}

type GcmType struct {
	GcmData `json:"data"`
}

type GcmData struct {
	Message string  `json:"message,omitempty"`
	Url     *string `json:"url,omitempty"`
}

func (pt *PushToken) SendMessage(alert string, badge, sound, url *string) (*sns.PublishResponse, error) {
	var message []byte
	var payload interface{}

	if pt.ArnType == "APNS" || pt.ArnType == "APNS_SANDBOX" {
		body, err := json.Marshal(ApsType{
			ApsData{
				Alert: alert,
				Badge: badge,
				Sound: sound,
			},
		})
		if err != nil {
			return nil, err
		}

		switch pt.ArnType {
		case "APNS":
			payload = Apns{string(body)}
		case "APNS_SANDBOX":
			payload = ApnsSandbox{string(body)}
		}
	}

	if pt.ArnType == "GCM" {
		body, err := json.Marshal(GcmType{
			GcmData{
				Message: alert,
				Url:     url,
			},
		})
		if err != nil {
			return nil, err
		}

		payload = Gcm{string(body)}
	}

	message, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	options := sns.PublishOptions{
		Message:          string(message),
		MessageStructure: "json",
		TargetArn:        pt.Arn,
	}

	pr, err := awsSns.Publish(&options)
	return pr, err
}

func (pt *PushToken) IsEnabled() (bool, error) {
	resp, err := awsSns.GetEndpointAttributes(pt.Arn)
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
		Uid:     uid,
		Arn:     *arn,
		ArnType: strings.Split(SnsArn, "/")[1],
		Token:   token,
	}

	return &pushtoken, nil
}

func NewArn(token string) (*string, error) {
	opts := sns.PlatformEndpointOptions{
		PlatformApplicationArn: SnsArn,
		Token: token,
	}

	resp, err := awsSns.CreatePlatformEndpoint(&opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &resp.EndpointArn, nil
}
