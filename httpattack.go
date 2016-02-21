package httpattack

import (
	"log"
	"net/http"
	//"sync"
)

// Attack provides testing of http endpoints
type Attack struct {
	// Number of concurrent users
	User int
	// Workers per each user
	Workers int
	//your client
	Client *http.Client
	// Scenarios
	Scenarios *Schenario
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

func (attack *Attack) run() {
	/*workers.Add(1)
	defer workers.Done()*/
}

func buildHttpRequest(httpAction HttpAction, sessionMap map[string]string) (*http.Request, error) {
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

	return req, nil
}
