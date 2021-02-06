package ftapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Error contains information about any non 2xx status code response.
type Error struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Details []interface{} `json:"details"`
	Body    string
}

type errorReply struct {
	Error *Error `json:"error"`
}

func (e *Error) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("ftapi: Got HTTP response code %d with body %v", e.Code, e.Body)
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "ftapi: Error %d: ", e.Code)
	if e.Message != "" {
		fmt.Fprintf(&buf, "%s", e.Message)
	}
	return buf.String()
}

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	errData, err := ioutil.ReadAll(res.Body)
	if err == nil {
		errType := new(errorReply)
		err = json.Unmarshal(errData, errType)
		if err == nil && errType.Error != nil {
			if errType.Error.Code == 0 {
				errType.Error.Code = res.StatusCode
			}
			errType.Error.Body = string(errData)
			return errType.Error
		}
	}
	return &Error{
		Code: res.StatusCode,
		Body: string(errData),
	}
}
