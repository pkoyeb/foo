package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	for {
		for port := 9000; port <= 9999; port++ {
			url := fmt.Sprintf("http://127.0.0.1:%d/clusters", port)
			fmt.Printf("Scanning port: %d", port)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(fmt.Sprint("Unnable to reach admin port : %s", err.Error()))
				time.Sleep(time.Second)
				continue
			}

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(fmt.Sprintf("Response from port %d : %s", port, body))

		}
	}
}
