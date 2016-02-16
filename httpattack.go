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