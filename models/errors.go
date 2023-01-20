package models

import (
	"fmt"
	"time"
)

type ErrorEndpoint struct {
	User  string `json:"user,omitempty"`
	Path  string `json:"path"`
	Err   error  `json:"error"`
	Cause error  `json:"cause"`
}

func (ee *ErrorEndpoint) Error() string {

	result := fmt.Sprintf("[%v] ", time.Now().String())
	if ee.User != "" {
		result += ee.User + ": "
	}

	if ee.Path != "" {
		result += ee.Path + ": "
	}

	return fmt.Sprintf("%v\nerr %v cause %v\n",
		result, ee.Err.Error(), ee.Cause.Error())
}
