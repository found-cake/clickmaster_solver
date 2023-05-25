package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func fetchData(u string, i int) {
	data := url.Values{}
	data.Set("n", strconv.Itoa(i))

	resp, err := http.PostForm(u, data)
	if err != nil {
		println("errorId is", i)
		return
	}
	defer resp.Body.Close()

	var result string
	if resp.StatusCode == http.StatusOK {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)
		result = buf.String()
	} else {
		result = "nope! click other"
	}

	if result != "nope! click other" {
		fmt.Printf("%d: %s\n", i, result)
	}
}

func main() {
	u := "http://#/level1/p16/data.php" //혹시 몰라서 url 비공개
	n := 0

	for i := 0; i < 105000; i++ {
		if i%1000 == 0 {
			println("wait", n)
			time.Sleep(time.Second / 2)
			n++
		}
		fetchData(u, i)
	}

}
