package main

import (

	"log"
)

func main(){

	store,err:=NewPostgressStore()
	if err!=nil{
		log.Fatal(err)
	}
	store.init()
	server:=NewApiServer(":8080",store)
	server.Run()
}