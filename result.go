package httpattack

import (
   "net/http"
)

type Response struct {
	Result  string
}

func provideResponse(method string, client *http.Client) error {
	/*_, err := http.NewRequest(method)
	if err != nil {
		return
	}*/
	return nil
}