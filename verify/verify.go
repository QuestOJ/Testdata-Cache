package verify

import (
	"errors"

	"github.com/QuestOJ/testdata-cache/typedef"
)

func Verify(judgerName string, password string, config typedef.Config) (bool, error) {
	request := typedef.Request{}

	request.URL = config.Server
	request.Method = "POST"
	request.NotRedirect = true
	request.Data = map[string]string{
		"token":       config.ClientID,
		"secret":      config.ClientSecret,
		"judger_name": judgerName,
		"password":    password,
	}

	response, err := HTTPRequest(request)

	if err != nil {
		return false, err
	}

	if string(response.ResponseBody) == "ok" {
		return true, nil
	}

	return false, errors.New(string(response.ResponseBody))
}
