package server

import (
	"github.com/501miles/logger"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

var url = "http://www.evan0.xyz:8501/v1"
var registerUrl = url + "/agent/service/register"

func RegisterToConsul() {
	param := map[string]interface{}{
		"ID": "redis2",
		"Name": "redis2",
		"Address": "110.120.119.911",
		"Port": 12306,
		"Meta": map[string]string{
			"data": "A",
			"info": "B",
		},
		"Check": map[string]string{
			"Name":                           "check-name",
			"CheckID":                        "check-id",
			"DeregisterCriticalServiceAfter": "30s",
			"Interval":                       "5s",
			"Timeout":                        "1s",
			"Tcp":                            "127.0.0.1:9999",
			"Status":                         "passing",
		},
	}
	s, _ := jsoniter.MarshalToString(param)
	logger.Info(s)
	logger.Info(registerUrl)
	payload := strings.NewReader(s)
	req, err := http.NewRequest("PUT", registerUrl, payload)
	if err != nil {
		logger.Error(err)
	}
	req.Header.Set("X-Consul-Token", "7f85db13-c45f-f619-3acc-756d2d9af9cf")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(err)
	}
	logger.Info(res.StatusCode)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	logger.Info(string(body))
}
