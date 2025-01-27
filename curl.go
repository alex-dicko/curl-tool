package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	url := args[0]

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println(string(body))
}
