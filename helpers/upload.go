package helpers

import (
  "fmt"
  "errors"
  "bytes"
  "mime/multipart"
  "net/http"
  "encoding/json"
//   "io/ioutil"
)

type Res map[string]interface{}

func Upload(file string) (string, error){
	url := "https://api.imgur.com/3/image"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image", file)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}


	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", "Client-ID eeda5c647156e3e")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()

	var r Res
	json.NewDecoder(res.Body).Decode(&r)
	if r["success"] == true {
		return r["data"].(map[string]interface{})["link"].(string), nil
	}
	return "", errors.New("No link")
}