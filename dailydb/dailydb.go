package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "encoding/json"
  "strings"
  "net/http"
)

type AdRuleResponse struct {
  AdServer string
}

var AdRules map[string]string

func init() {
  AdRules = make(map[string]string)
}

func main() {

  file, err := ioutil.ReadFile("data.json")
  if err != nil {
    log.Fatalf("Can't read data.json file")
  }

  LoadRulesFromJson(string(file))

  fmt.Println("Starting server on port 8021")

  http.HandleFunc("/", HandleServiceRequest)
  fmt.Println(http.ListenAndServe(":8021", nil))
}

// listen to request and serve back a json response with the right rule
func HandleServiceRequest(w http.ResponseWriter, req *http.Request) {

  respData := AdRuleResponse{"not-found"} // default response
  log.Printf("Request: %s \n", req.URL)

  key := req.URL.Query().Get("geo")
  if key != "" {
      res := GetRule(key)
      respData = AdRuleResponse{res}
  }

  js, err := json.Marshal(respData)
  if err != nil {
    log.Printf("JSON Error: %s", err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)

}


func LoadRulesFromJson(content string) {

  decoder := json.NewDecoder(strings.NewReader(content))
    // read open bracket
  	_, err := decoder.Token()
  	if err != nil {
  		log.Fatal(err)
  	}

  type Rule struct {
    AdServer, GEO string
  }

  for decoder.More() {
    var r Rule
    err := decoder.Decode(&r)
    if err != nil {
      log.Fatalf("error: %s", err)
    }
    //fmt.Printf("Rule %s %s \n", r.AdServer, r.GEO)
    AdRules[r.GEO] = r.AdServer
  }
}


func GetRule(key string) string {
  res := AdRules[key]
  if res == "" {
    res = "default.ads-daily.net"
  }
  return res
}
