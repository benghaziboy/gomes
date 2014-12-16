package gomes

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

const (
	mockCreatePlatformEndpoint = `
    <CreatePlatformEndpointResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
      <CreatePlatformEndpointResult>
        <EndpointArn>arn:aws:sns:us-east-2:123456789012:endpoint/GCM/gcmpushapp/5e3e9847-3183-3f18-a7e8-671c3a57d4b3</EndpointArn>
      </CreatePlatformEndpointResult>
      <ResponseMetadata>
        <RequestId>6613341d-3e15-53f7-bf3c-7e56994ba278</RequestId>
      </ResponseMetadata>
    </CreatePlatformEndpointResponse>`

	mockGetEndpointAttributes = `
    <GetEndpointAttributesResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
      <GetEndpointAttributesResult>
        <Attributes>
          <entry>
            <key>Enabled</key>
            <value>true</value>
          </entry>
          <entry>
            <key>CustomUserData</key>
            <value>UserId=01234567</value>
          </entry>
          <entry>
            <key>Token</key>
            <value>APA91bGi7fFachkC1xjlqT66VYEucGHochmf1VQAr9k...jsM0PKPxKhddCzx6paEsyay9Zn3D4wNUJb8m6HZrBEXAMPLE</value>
          </entry>
        </Attributes>
      </GetEndpointAttributesResult>
      <ResponseMetadata>
        <RequestId>6c725a19-a142-5b77-94f9-1055a9ea04e7</RequestId>
      </ResponseMetadata>
    </GetEndpointAttributesResponse>
    `

	mockSendMessageResponse = `
    <PublishResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
      <PublishResult>
        <MessageId>567910cd-659e-55d4-8ccb-5aaf14679dc0</MessageId>
      </PublishResult>
      <ResponseMetadata>
        <RequestId>d74b8436-ae13-5ab4-a9ff-ce54dfea72a0</RequestId>
      </ResponseMetadata>
    </PublishResponse>
    `
)

type Handler func(http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createMockServer() *httptest.Server {
	router := mux.NewRouter()
	router.Handle("/", Handler(handleMockRequests))
	return httptest.NewServer(router)
}

func handleMockRequests(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, mockGetEndpointAttributes)

	} else {
		m, err := parseAwsRequestBody(r)
		if err != nil {
			return err
		}

		switch m["Action"] {
		case "CreatePlatformEndpoint":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, mockCreatePlatformEndpoint)
		case "Publish":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, mockSendMessageResponse)
		}
	}

	return nil
}

func parseAwsRequestBody(r *http.Request) (map[string]string, error) {
	m := make(map[string]string)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	q := strings.Split(string(b), "&")
	for _, v := range q {
		qp := strings.Split(v, "=")
		m[qp[0]] = qp[1]
	}

	return m, nil
}
