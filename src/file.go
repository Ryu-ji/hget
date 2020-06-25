package main

import (
	"fmt"
	"log"
	"os"
)

func write(file string, bytes []byte) {

	if file == "" {
		log.Println("不正なファイル名です。")
		return
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {

		log.Println(err)
		return
	}

	defer f.Close()

	fmt.Fprintf(f, "%s", bytes)

}
