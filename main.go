package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	current_time := time.Now()
	t := current_time.Format("2006-January-02")
	file, err := os.Create(t)
	if err != nil {
		fmt.Println("Unable to open file:", err)
	}

	len, err := file.WriteString("Hello World")

	if err != nil {
		fmt.Println("Unable to write data:", err)
	}

	file.Close()

	fmt.Printf("%d character written successfully into file", len)
}
