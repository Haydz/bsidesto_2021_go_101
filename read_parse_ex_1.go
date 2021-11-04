package main

import (
   "bufio"
   "fmt"
   "log"
   "os"
   "strings"
)

func main() {
 

   file, err := os.Open("access.log")
   if err != nil {
       log.Fatalf("failed to open")

   }
   scanner := bufio.NewScanner(file)
   scanner.Split(bufio.ScanLines) // don't necessarily need (default is new line)

   var text []string

   for scanner.Scan() {
       text = append(text, scanner.Text())
   }

   file.Close()

   for _, each_ln := range text {
    //    fmt.Println(each_ln)
       if strings.Contains(each_ln, "bsides") {
           fmt.Println("Found the line with bsides:", each_ln)
       }
   }
}



