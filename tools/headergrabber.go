package tools

import (
	"fmt"
	"net/http"
	"time"
)

func HeaderGrabber(url string, output string) {
	results := []string{}
	results = append(results, fmt.Sprintf("HTTP Header Analysis: %s", url))
	results = append(results, "")
	
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	results = append(results, fmt.Sprintf("Status Code: %d %s", resp.StatusCode, resp.Status))
	results = append(results, "")
	results = append(results, "Headers:")
	
	fmt.Printf("Status Code: %d %s\n", resp.StatusCode, resp.Status)
	fmt.Println("\nHeaders:")
	
	for key, values := range resp.Header {
		for _, value := range values {
			result := fmt.Sprintf("%s: %s", key, value)
			results = append(results, result)
			fmt.Println(result)
		}
	}
	
	// Security header analysis
	results = append(results, "")
	results = append(results, "Security Analysis:")
	fmt.Println("\nSecurity Analysis:")
	
	securityHeaders := map[string]string{
		"X-Frame-Options":           "Protects against clickjacking",
		"X-Content-Type-Options":    "Prevents MIME type sniffing",
		"Strict-Transport-Security": "Enforces HTTPS",
		"Content-Security-Policy":   "Prevents XSS attacks",
		"X-XSS-Protection":          "XSS filter",
	}
	
	for header, description := range securityHeaders {
		if value := resp.Header.Get(header); value != "" {
			result := fmt.Sprintf("[+] %s: %s (%s)", header, value, description)
			results = append(results, result)
			fmt.Println(result)
		} else {
			result := fmt.Sprintf("[-] %s: MISSING (%s)", header, description)
			results = append(results, result)
			fmt.Println(result)
		}
	}
	
	if output != "" {
		writeToFile(output, results)
	}
}
