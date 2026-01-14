package tools

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func DirFinder(url string, wordlist string, output string) {
	results := []string{}
	results = append(results, fmt.Sprintf("Directory brute-forcing: %s", url))
	results = append(results, fmt.Sprintf("Wordlist: %s", wordlist))
	results = append(results, "")
	
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Printf("Error opening wordlist: %v\n", err)
		return
	}
	defer file.Close()
	
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dir := strings.TrimSpace(scanner.Text())
		if dir == "" {
			continue
		}
		
		targetURL := strings.TrimRight(url, "/") + "/" + strings.TrimLeft(dir, "/")
		resp, err := client.Get(targetURL)
		
		if err == nil {
			result := fmt.Sprintf("%s - Status: %d", targetURL, resp.StatusCode)
			results = append(results, result)
			fmt.Println(result)
			resp.Body.Close()
		}
	}
	
	if output != "" {
		writeToFile(output, results)
	}
}
