type httpattack

import (
  "time"
)

type HttpAction struct {
	Addr   string
	// GET, REMOVE, POST, DELETE, UPDATE
	Method string
	// Breif description of this action
	Title  string
	Params interface{}
}