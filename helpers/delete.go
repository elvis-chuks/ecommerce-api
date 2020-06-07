package helpers

import (
  "fmt"
  "bytes"
  "mime/multipart"
  "net/http"
  "io/ioutil"
)

func Delete() {

  url := "https://api.imgur.com/3/image/{{imageHash}}"
  method := "DELETE"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
  }


  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Authorization", "Bearer {{accessToken}}")

  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}