package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	i := 0
	for {
		sendPost()
		i++
		fmt.Println(i)
	}
}

func sendPost() {

	url := "http://localhost:9094/api/email"
	method := "POST"

	payload := strings.NewReader(`{
    "from_mail": "from@mail.ro",
    "to_mail": "to@mail.ro",
    "mail_cc": "cc@mail.ro",
    "body": "body mare....",
    "subject": "subiectu"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
