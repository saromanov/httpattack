package httpattack

import (
  //"time"
)

type HttpAction struct {
	Addr   string
	// GET, REMOVE, POST, DELETE, UPDATE
	Method string
	// Breif description of this action
	Title  string
	//Accept
	Accept  string
	Params interface{}
}