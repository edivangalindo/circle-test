package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No tokens detected. Hint: cat tokens.txt | circle-test")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		token := scanner.Text()

		err := testCircleCIToken(token)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}

func testCircleCIToken(token string) error {
	fmt.Println("Testing token ->", token)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://circleci.com/api/v1.1/me", nil)

	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Circle-Token", token)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var prettyJson bytes.Buffer
	err = json.Indent(&prettyJson, body, "", "\t")

	if err != nil {
		return err
	}

	fmt.Println(string(prettyJson.Bytes()))

	return nil
}
