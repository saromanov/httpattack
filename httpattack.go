package httpattack

import (
   "net/http"
   "fmt"
   "sync"
)

// Attack provides testing of http endpoints
type Attack struct {
	// Number of concurrent users
	User  int
	// Workers per each user
	Workers int
	//your client
	Client http*Client
	// Scenarios
}

func New()*Attack{
	attack := new(Attack)
	return attack
}

func (attack* Attack) Run() error {
	for i := 0; i < attack.Users; i++ {
		go run()
	}
}

func (attack *Attack) run(workers *sync.WaitGroup){
	workers.Add(1)
	defer workers.Done()
}

func buildHttpRequest(httpAction HttpAction, sessionMap map[string]string) *http.Request {
    var req *http.Request
    var err error
    if httpAction.Body != "" {
        reader := strings.NewReader(SubstParams(sessionMap, httpAction.Body))
        req, err = http.NewRequest(httpAction.Method, SubstParams(sessionMap, httpAction.Url), reader)
    } else {
        req, err = http.NewRequest(httpAction.Method, SubstParams(sessionMap, httpAction.Url), nil)
    }
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Add("Accept", httpAction.Accept)

    return req
}