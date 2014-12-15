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
	c.Assert(*arn, Equals, "arn:aws:sns:us-west-2:123456789012:endpoint/GCM/gcmpushapp/5e3e9847-3183-3f18-a7e8-671c3a57d4b3")
}
