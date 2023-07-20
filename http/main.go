package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lg := logWriter{}
	io.Copy(lg, res.Body)
	// io.Copy(os.Stdout, res.Body)
	// bs := make([]byte, 9999999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
