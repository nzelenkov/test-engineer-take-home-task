package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Parse arguments
	prot := flag.String("prot", "http", "Protocol of substrate API")
	host := flag.String("host", "localhost", "Host of substrate API")
	port := flag.String("port", "9933", "Port of substrate API")
	timeout := flag.Int("timeout", 10, "Timeout to disconnect from substrate API")
	flag.Parse()
	url := *prot + "://" + *host + ":" + *port
	log.Println("Substrate API url:", url)

	// Build the request data for getBlock
	reqBody, err := json.Marshal(map[string]string{
		"id":      "1",
		"jsonrpc": "2.0",
		"method":  "chain_getBlock",
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Loop the request until the block height reaches 10
	i := 0
	for i < 10 {
		// Build request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
		if err != nil {
			log.Fatal("Error reading request. ", err)
			os.Exit(2)
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{Timeout: time.Duration(time.Second * time.Duration(*timeout))}

		// Send request
		response, err := client.Do(req)
		if err != nil {
			log.Fatal("Error reading response. ", err)
			os.Exit(3)
		}
		defer response.Body.Close()

		// Check response data is valid
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Response data is invalid. ", err)
			os.Exit(4)
		}

		// Parse response to find block number
		var result map[string]interface{}
		json.Unmarshal([]byte(responseData), &result)
		blockN := result["result"].(map[string]interface{})["block"].(map[string]interface{})["header"].(map[string]interface{})["number"]

		// Convert Hex encoded value to int
		i2, err := strconv.ParseInt(strings.TrimLeft(blockN.(string), "0x"), 16, 16)
		if err != nil {
			log.Fatal("Error parsing the block number. ", err)
			os.Exit(5)
		}
		if i2 > 10 {
			log.Fatal("Block height out of boundry [0-10], current height is ", i2)
			os.Exit(6)
		}
		i = int(i2)
	}
	log.Println("Block height reached ", i)
	os.Exit(0)
}
