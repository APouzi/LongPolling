package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	// ctx,cancel := context.WithCancel(context.Background())
	
	
	// ch := make(chan string)
	
	
	http.HandleFunc("/",longpoll)
	http.ListenAndServe(":8080", nil)
	
}

func longpoll(w http.ResponseWriter, r *http.Request){
	ctx,cancel := context.WithCancel(r.Context())
	ch := make(chan string)
	go func( ch chan string, ctxIn context.Context, cancelIn context.CancelFunc){
		
		time.Sleep(time.Second * 4)
		ch <- "message sent!"
		cancelIn()
		
	}(ch, ctx, cancel)

	for{
		select{
			case v1 := <-ch:
				w.Write([]byte(v1))
				return
				// cancel()
			case <-ctx.Done():
				// wg.Done()
				log.Fatal("we out!")
				return
			}
	}
}