package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"math/rand"
)

type CountDatas struct {
	Line_count int;
	File_count int;
	Issue_count int;
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler));
	http.Handle("/", http.FileServer(http.Dir(".")));

	err := http.ListenAndServe(":8089", nil);
	
	if err != nil {
		panic("ListenAndServe: " + err.Error());
	}

}

func echoHandler(ws *websocket.Conn){
	var countDatas CountDatas;
	for{
		countDatas.File_count = 10000;
		countDatas.Line_count = 10000;
		countDatas.Issue_count = 10000;		
		for a:= 1;a<=100;a++{			
			send_msg, send_err :=  json.Marshal(countDatas);
			if send_err != nil {
				log.Fatal(send_err)
			}
			_, err := ws.Write([]byte(send_msg))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Send: %s\n", send_msg)
			rand1 := rand.Intn(20);
			rand2 := rand.Intn(5);
			rand3 := rand.Intn(100);
			
			countDatas.Line_count += rand1;
			countDatas.File_count += rand2;
			countDatas.Issue_count += rand3;
			time.Sleep(2*time.Second);  
		}		
		
	}
	
}