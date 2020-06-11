package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main(){
dbpath :="testdb"
db,err:=leveldb.OpenFile(dbpath,nil)
	if err!=nil {
		log.Fatal(err)
	}
}