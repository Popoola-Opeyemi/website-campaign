package utils

import (
	"errors"

	"github.com/fatih/structs"
	"github.com/serenize/snaker"
)

// Response wraps the payload sent to every request
type Response struct {
	Store  map[string]interface{} `json:"store"`
	Errors map[string]string      `json:"errors"`
	Done   bool                   `json:"done"`
}

// SetData populates the Data field for response
func (s *Response) SetData(value interface{}) error {
	store := map[string]interface{}{}
	if s.Store == nil {
		s.Store = map[string]interface{}{}
	}
	if !structs.IsStruct(value) {
		return errors.New("response: required parameter should be a struct")
	}
	stx := structs.New(value)
	stx.FillMap(store)
	for k, v := range store {
		s.Store[snaker.CamelToSnake(k)] = v
	}
	s.Done = true
	return nil
}

// Set sets a singly named value in response's data
func (s *Response) Set(name string, value interface{}) {
	if s.Store == nil {
		s.Store = map[string]interface{}{}
	}
	s.Store[name] = value
	s.Done = true
}

// SetError sets uncovers the error and populates the error
// value of the response
func (s *Response) SetError(err error) {
	if s.Errors == nil {
		s.Errors = map[string]string{}
	}
	if errmsg, ok := err.(ErrMsg); ok {
		for k, v := range errmsg {
			s.Errors[k] = v
		}
		return
	}
	s.Errors["message"] = err.Error()
	return
}
