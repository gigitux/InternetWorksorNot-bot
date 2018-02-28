package main

import (
  "time"
  "fmt"
  "net/http"
  tb "gopkg.in/tucnak/telebot.v2"
)

const tokenBot string = "insert your token"
//insert id User
const idUser int64 = 1

func isInternetWork() bool {
  resp, err := http.Get("http://www.google.com/")
  if err != nil {
    return false
  }
  fmt.Println(resp)
  return true
}

func main() {
  for true == true {
    if (isInternetWork()) {
      b, errBot := tb.NewBot(tb.Settings {
  		Token:  tokenBot,
  		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
  	})
  	if errBot != nil {
  		fmt.Println(errBot)
  	}
    chat := tb.Chat{ID: idUser}
    b.Send(&chat, "Internet finally works")
    fmt.Println("Internet finally works")
    break
    } else {
      fmt.Println("Internet does not work yet :(")
      time.Sleep(1 * time.Minute)
    }
  }
}
