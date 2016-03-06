package httpattack

import (
	"log"
	"net/http"
	"time"
	//"sync"
)

// Attack provides testing of http endpoints
type Attack struct {
	// Number of concurrent users
	User int
	// Workers per each user
	Workers int
	// File for store results
	Outputfile string
	//your client
	Client *http.Client
	// Scenarios
	Scenarios *Schenario
	// Actions config. Should be in json format
	ActionsConfig  string
}

// New initialize new http attack item
func New() *Attack {
	attack := new(Attack)
	return attack
}

// Run provides main method for running items
func (attack *Attack) Run() error {
	for i := 0; i < attack.User; i++ {
		go attack.run()
	}

	return nil
}

// inner call for run
func (attack *Attack) run() {
	now := time.Now()
	req, err := buildHttpRequest(action, map[string]string{})
	if err != nil {
		return
	}

	sub := time.Now().Sub(now)
}

// buildHTTPRequest is a helpful method for building request
func buildHTTPRequest(httpAction HttpAction, sessionMap map[string]string) (*http.Request, error) {
	var req *http.Request
	var err error
	if httpAction.Method != "" {
		//reader := strings.NewReader(SubstParams(sessionMap, httpAction.Body))
		req, err = http.NewRequest(httpAction.Method, httpAction.Addr, nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(httpAction.Method, httpAction.Addr, nil)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", httpAction.Accept)
	req.Header.Add("Allowd", "")

	return req, nil
}

func buildResponse(req *http.Request) string {
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}
