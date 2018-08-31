package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	fmt.Println("Blocking?")
	fmt.Println(res)

	if err != nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	result := string(page)[0:500]
	fmt.Println(result)

	scaner := bufio.NewScanner(strings.NewReader(result))

	scaner.Split(bufio.ScanWords)

	i := 0
	for scaner.Scan() {
		fmt.Println("-", i, "-", scaner.Text())
		i++
	}

	if err := scaner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
}
