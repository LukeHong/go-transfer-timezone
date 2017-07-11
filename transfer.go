package main

import (
    "fmt"
    "time"
    "strings"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("data.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    rows := strings.Split(string(data), "\r\n")

    var result string
    for i := range rows {
        temp := strings.Split(rows[i], "\t")
        t, _ := time.Parse("2006-01-02T15:04:05Z", temp[0])
        t = t.Add(time.Hour * time.Duration(-8))

        result += fmt.Sprintf("{\"time\":\"%d\", %s\r\n", t.Unix(), temp[1][1:])
    }
    ioutil.WriteFile("processed.txt", []byte(result), 0644)
}
