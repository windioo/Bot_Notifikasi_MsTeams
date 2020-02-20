package main

import(
  "log"
  "net/http"
  "encoding/json"
  "bytes"
  "os"
  "github.com/joho/godotenv"
)

type Payload struct {
	Text string `json:"Text"`
}

func main(){
      var param = os.Args[1:]
      data := Payload{param[0]}

      payloadBytes, _ := json.Marshal(data)
      err:= godotenv.Load("config.env")

    	if  err != nil {
        log.Panic("Error loading .env file")
      }
      Url:= os.Getenv("Url")
      body := bytes.NewReader(payloadBytes)
      req, err := http.NewRequest("POST", Url, body)
      if err != nil {

        log.Panic(err)
      }

      req.Header.Set("Content-Type", "application/json")

      resp, err := http.DefaultClient.Do(req)
      if err != nil {

        log.Panic(err)
      }

      defer resp.Body.Close()
}
