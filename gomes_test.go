package gomes

import (
	"github.com/crowdmob/goamz/aws"
	. "gopkg.in/check.v1"
)

var _ = Suite(&AwsSuite{})

func (s *AwsSuite) TestAwsAuthentication(c *C) {
	err := authenticateAws()
	c.Assert(err, IsNil)
	c.Assert(auth.AccessKey, Equals, awsAccessKey)
	c.Assert(auth.SecretKey, Equals, awsSecretKey)
}

func (s *AwsSuite) TestSnsEndpoint(c *C) {
	err := authenticateAws()
	c.Assert(err, IsNil)

	err = connectSNS(auth)
	c.Assert(err, IsNil)
	c.Assert(snsConn.Auth, Equals, *auth)
	c.Assert(snsConn.Region, Equals, aws.USEast)
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
