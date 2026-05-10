package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func readWithMessage(s *bufio.Scanner, msg string) string {
	fmt.Fprint(os.Stdout, msg)
	s.Scan()
	return string(s.Bytes())
}

func getUrl(s *bufio.Scanner) string {
	return readWithMessage(s, "Url: ")
}

func getMethod(s *bufio.Scanner) string {
	return readWithMessage(s, "Method: ")
}

func getHeaders(s *bufio.Scanner) http.Header {
	headers := make(http.Header)
	for {
		name := readWithMessage(s, "Header name: ")
		if len(name) == 0 {
			return headers
		}

		value := readWithMessage(s, "Header value: ")
		headers[name] = append(headers[name], value)
	}
}

func getBody(s *bufio.Scanner) string {
	return readWithMessage(s, "Body: ")
}

func setLogger() {
	options := slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	url := getUrl(scanner)
	method := getMethod(scanner)
	headers := getHeaders(scanner)
	body := getBody(scanner)

	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		method,
		url,
		strings.NewReader(body),
	)
	req.Header = headers
	if err != nil {
		panic(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		panic("An error happened while making the request.")
	}

	fmt.Printf("\n%d\n", res.StatusCode)
	for header, values := range res.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", header, value)
		}
	}

	reader := bufio.NewReader(res.Body)
	bodyRes := make([]byte, 2048)
	var str string
	for {
		read, err := reader.Read(bodyRes)
		str += string(bodyRes[read:])
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("error while getting body")
		}
	}

	fmt.Printf("%s\n", str)

}
