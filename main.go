package main

import (
	"encoding/json"
	"io/ioutil"
	
  "fmt"
  "github.com/uswitch/segment-config-go/segment"
  "os"
  // "reflect"
)


func main() {
  client := segment.NewClient("t5dwvaonRp-UqFZJX0iOStP1Cj8_Ea25z8FuxDQWL8U.sHdro1niwYG3oqrnVxSgBUEYNdNPUNVvNiRRUutCHbc", "uswitch-sandbox")
  
  jsonFile, err := os.Open("segment/trackingplan.json")

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println("successfully opened json file")
  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)

  var result map[string]interface{}
  json.Unmarshal([]byte(byteValue), &result)
  
  tp, err := client.CreateTrackingPlan(result)

  fmt.Println(tp)

}
