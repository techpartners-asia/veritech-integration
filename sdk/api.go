package veritechSdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type API struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

var (
	CommandEndpoit = API{
		Url:    "https://erp.teso.mn:8080/erp-services/RestWS/runJson",
		Method: http.MethodPost,
	}
)

func (v *Sdk) Send(body interface{}) (response []byte, err error) {

	// assert body to map
	bodyMap, ok := body.(map[string]interface{})
	if !ok {
		return nil, errors.New("body is not a map")
	}
	bodyMap["username"] = v.Username
	bodyMap["password"] = v.Password
	bodyMap["command"] = v.Command
	bodyMap["languageCode"] = "mn"
	bodyMap["ignorepermission"] = "1"

	var requestByte []byte
	var requestBody *bytes.Reader

	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(CommandEndpoit.Method, CommandEndpoit.Url, requestBody)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}

	return
}
