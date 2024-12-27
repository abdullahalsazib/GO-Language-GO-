package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Define the URL
	url := "https://ngcol.gov.bd/Home/getStudentsList"

	// Define the payload (replace with actual data)
	payload := map[string]interface{}{
		"key1": "value1", // Replace "key1" and "value1" with actual keys and values
		"key2": "value2",
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers (replace with actual headers if needed)
	req.Header.Set("Content-Type", "application/json")
	// Add other headers if required (e.g., Authorization, User-Agent)
	// req.Header.Set("Authorization", "Bearer YOUR_TOKEN")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Check the status code
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Data Retrieved Successfully!")
		fmt.Println(string(body)) // Print the response body
	} else {
		fmt.Printf("Failed to retrieve data: %d\n", resp.StatusCode)
		fmt.Println(string(body)) // Print the error response body for debugging
	}
}

