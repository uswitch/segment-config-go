package main

import (
	// "encoding/json"
  // "io/ioutil"
  // "os"
  "fmt"
  "github.com/uswitch/segment-config-go/segment"
  // "reflect"
)


func main() {

  workspace := "uswitch-sandbox"
  accessToken := "t5dwvaonRp-UqFZJX0iOStP1Cj8_Ea25z8FuxDQWL8U.sHdro1niwYG3oqrnVxSgBUEYNdNPUNVvNiRRUutCHbc"
  client, err := segment.NewClient(&accessToken, &workspace)

  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(client)
  
  // jsonFile, err := os.Open("segment/trackingplan.json")

  // if err != nil {
  //   fmt.Println(err)
  // }

  // fmt.Println("successfully opened json file")
  // defer jsonFile.Close()
  // byteValue, _ := ioutil.ReadAll(jsonFile)

  // var result map[string]interface{}
  // json.Unmarshal([]byte(byteValue), &result)
  
  // tp, err := client.CreateTrackingPlan(result)

  // fmt.Println(tp)
}
