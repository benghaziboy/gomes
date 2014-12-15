package gomes

import (
	"github.com/crowdmob/goamz/aws"
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type AwsSuite struct{}

var (
	awsAccessKey = "THISISANACCESSKEY"
	awsSecretKey = "THISISASECRETKEY"
	deviceToken  = "FE66489F304DC75B8D6E8200DFF8A456E8DAEACEC428B427E9518741C92C6660"
)

func (s *AwsSuite) SetUpTest(c *C) {
	_, err := Conn.Exec(`DELETE FROM push_tokens;`)
	c.Assert(err, IsNil)

	testServer := createMockServer()
	c.Assert(err, IsNil)
	aws.USEast.SNSEndpoint = testServer.URL

	os.Setenv("AWS_ACCESS_KEY", awsAccessKey)
	os.Setenv("AWS_SECRET_KEY", awsSecretKey)
}
