package gomes

import (
	"github.com/crowdmob/goamz/aws"
	. "gopkg.in/check.v1"
)

var _ = Suite(&AwsSuite{})

func (s *AwsSuite) TestAwsAuthentication(c *C) {
	err := authenticateAws()
	c.Assert(err, IsNil)
	c.Assert(awsAuth.AccessKey, Equals, awsAccessKey)
	c.Assert(awsAuth.SecretKey, Equals, awsSecretKey)
}

func (s *AwsSuite) TestSnsEndpoint(c *C) {
	err := authenticateAws()
	c.Assert(err, IsNil)

	err = connectSNS(awsAuth)
	c.Assert(err, IsNil)
	c.Assert(awsSns.Auth, Equals, *awsAuth)
	c.Assert(awsSns.Region, Equals, aws.USEast)
}

func (s *AwsSuite) TestNewArn(c *C) {
	err := initAws()
	c.Assert(err, IsNil)

	arn, err := NewArn(deviceToken)
	c.Assert(err, IsNil)
	c.Assert(*arn, Equals, respGcmArn)
}

func (s *AwsSuite) TestNewPushToken(c *C) {
	err := initAws()
	c.Assert(err, IsNil)

	pt, err := New("dogman", deviceToken)
	c.Assert(err, IsNil)
	c.Assert(pt.Uid, Equals, "dogman")
	c.Assert(pt.Arn, Equals, respGcmArn)
	c.Assert(pt.ArnType, Equals, "APNS_SANDBOX")
	c.Assert(pt.Token, Equals, deviceToken)
}

func (s *AwsSuite) TestPushTokenEnabled(c *C) {
	err := initAws()
	c.Assert(err, IsNil)

	pt, err := New("dogman", deviceToken)
	c.Assert(err, IsNil)

	enabled, err := pt.IsEnabled()
	c.Assert(err, IsNil)
	c.Assert(enabled, Equals, true)
}

func (s *AwsSuite) TestPushTokenSendMessage(c *C) {
	err := initAws()
	c.Assert(err, IsNil)

	pt, err := New("dogman", deviceToken)
	c.Assert(err, IsNil)

	alert := "Werner Herzog's Fun House"
	resp, err := pt.SendMessage(alert, nil, nil, nil)
	c.Assert(err, IsNil)
	c.Assert(resp.ResponseMetadata.RequestId, Equals, "d74b8436-ae13-5ab4-a9ff-ce54dfea72a0")
}
