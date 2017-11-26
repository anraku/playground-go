// curlをGoを使って実装
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	//fmt.Println("url: " + url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer resp.Body.Close()

	byte_resp, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte_resp))
}
