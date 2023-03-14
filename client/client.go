package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type gitLabClient struct {
	token string
}

func (c gitLabClient) GetMyOpenedMRs() []mergeRequest {
	body := c.makeRequest("/merge_requests?scope=assigned_to_me&state=opened")

	var listOfMR []mergeRequest

	err := json.Unmarshal(body, &listOfMR)
	if nil != err {
		panic(err)
	}

	return listOfMR
}

func (c gitLabClient) makeRequest(path string) []byte {
	url := "https://gitlab.com/api/v4" + path
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("PRIVATE-TOKEN", c.token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return []byte(body)
}

func GetClient(token string) gitLabClient {
	return gitLabClient{token}
}
