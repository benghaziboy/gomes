package gomes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

const (
	mockCreatePlatformEndpoint = `
    <CreatePlatformEndpointResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
      <CreatePlatformEndpointResult>
        <EndpointArn>arn:aws:sns:us-west-2:123456789012:endpoint/GCM/gcmpushapp/5e3e9847-3183-3f18-a7e8-671c3a57d4b3</EndpointArn>
      </CreatePlatformEndpointResult>
      <ResponseMetadata>
        <RequestId>6613341d-3e15-53f7-bf3c-7e56994ba278</RequestId>
      </ResponseMetadata>
    </CreatePlatformEndpointResponse>`
)

func createMockServer() *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		m := parseAwsRequest(b)

		switch m["Action"] {
		case "CreatePlatformEndpoint":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, mockCreatePlatformEndpoint)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "TestResponse")
	}

	return httptest.NewServer(http.HandlerFunc(handler))
}

func parseAwsRequest(b []byte) map[string]string {
	m := make(map[string]string)
	q := strings.Split(string(b), "&")
	for _, v := range q {
		qp := strings.Split(v, "=")
		m[qp[0]] = qp[1]
	}

	return m
}
