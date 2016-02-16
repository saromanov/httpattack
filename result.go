package httpattack

import (
   "net/http"
)

func provideResponse(method string, client *http.Client) error {
	resp, err := http.NewRequest()
}