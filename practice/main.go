package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://wiki.one.int.sap/wiki/pages/viewpage.action?pageId=3286771169"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Print(content)
}
