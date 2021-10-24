package main

import (
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
)
func main() {
   resp, err := http.Get("https://www.bsidesto.ca/")
   if err != nil {
       log.Fatalln(err)
   }
   defer resp.Body.Close()

   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
       log.Fatalln(err)
   }
   body_string := string(body)
   fmt.Println(body_string)
}
