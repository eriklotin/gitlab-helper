package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type JSON struct {
	Token string `json:"token"`
}

func checkErr(err error) {
	if nil != err {
		panic(err)
	}
}

func Init() JSON {
	var jsonDto JSON

	jsonStr, err := os.ReadFile("config.json")

	if os.IsNotExist(err) {
		fmt.Print("Enter your GitLab access token (read api): ")
		reader := bufio.NewReader(os.Stdin)
		token, _, err := reader.ReadLine()
		checkErr(err)
		jsonDto = JSON{string(token)}
		jsonStr, err = json.Marshal(&jsonDto)
		checkErr(err)
		err = os.WriteFile("config.json", jsonStr, 0644)
		checkErr(err)

		return jsonDto
	} else {
		checkErr(err)
	}

	err = json.Unmarshal(jsonStr, &jsonDto)
	checkErr(err)

	return jsonDto
}
