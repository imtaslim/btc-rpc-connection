package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

func main() {
    data, err := json.Marshal(map[string]interface{}{
        "method": "getreceivedbyaddress",
        "id":     "curltext",
        "params": []interface{}{"tb1qsu7p3znu22ea28w7vamrh7fpcnfgrrg386rlka"},
	"jsonrpc":"1.0",
    })
    if err != nil {
        log.Fatalf("Marshal: %v", err)
    }
    resp, err := http.Post("http://codemen:123456@143.244.162.98:8332/",
        "text/plain", strings.NewReader(string(data)))
    if err != nil {
        log.Fatalf("Post: %v", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("ReadAll: %v", err)
    }
    result := make(map[string]interface{})
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    log.Println(result)
}
