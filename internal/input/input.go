package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if err != nil {
		return "", err
	}
	return text, nil
}

func ReadUrl() (string, error) {
	fmt.Print("Enter URL: ")
	url, err := readInput()
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	return url, nil
}
