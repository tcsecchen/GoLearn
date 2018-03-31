package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler));
	http.Handle("/", http.FileServer(http.Dir(".")));

	err := http.ListenAndServe(":8089", nil);
	
	if err != nil {
		panic("ListenAndServe: " + err.Error());
	}

}

func echoHandler(ws *websocket.Conn){
	for{
		msg := make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive: %s\n", msg[:n])
		
		for a:= 1;a<=20;a++{
			send_msg := "[" + string(msg[:n]) + "]"
			m, err := ws.Write([]byte(send_msg))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Send: %s\n", msg[:m])
			time.Sleep(time.Second);  
		}		
		
	}
	
}