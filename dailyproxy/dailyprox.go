package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

type AdServerRule struct {
	AdServer string
}

var CachedAdRules map[string]string

func init() {
	CachedAdRules = make(map[string]string)
}

func main() {

	fmt.Println("Starting DailyMotion proxy server.")

	// listen for connections forver
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to list on: %s", err)
	}
	// endless loop to accept connections
	for {
		conn, err := ln.Accept()
		if err == nil {
			go handleConnection(conn)
		} else {
			log.Fatalf("failed to accept connection %s", err)
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		req, err := http.ReadRequest(reader)
		if err != nil {
			if err == io.EOF {
				log.Printf("Failed to read request: %s", err)
			}
			return
		}

		// find geo location from ip
		ip_address := req.URL.Query().Get("ip")
		geo := FindGeoByIpAddress(ip_address)

		// get the rule for service or cache
		rule := GetAdRuleFromService(geo)
		backendUrl := fmt.Sprintf("%s:80", rule)

		// sending the request to backend
		if be, err := net.Dial("tcp", backendUrl); err == nil {
			be_reader := bufio.NewReader(be)
			if err := req.Write(be); err == nil {
				// read the response from the backend
				if resp, err := http.ReadResponse(be_reader, req); err == nil {
					resp.Close = true
					if err := resp.Write(conn); err == nil {
						log.Printf("%s: %d", req.URL.Path, resp.StatusCode)
					}
					conn.Close()
				}
			}
		}

	}
}

func FindGeoByIpAddress(ip string) string {
	result := "Europe"
	switch {
	case strings.HasPrefix(ip, "127"):
		result = "Local"
	case strings.HasPrefix(ip, "28"):
		result = "France"
	case strings.HasPrefix(ip, "200"):
		result = "USA"
	}
	return result
}

func GetAdRuleFromService(geo string) string {

	// check first in the local cache if we already saved the rule
	if CachedAdRules[geo] != "" {
		return CachedAdRules[geo]
	}

	url := fmt.Sprintf("http://localhost:8021/?geo=%s", geo)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Get url error: %s", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read response from url %s", err)
	}

	rule := AdServerRule{}
	json.Unmarshal([]byte(body), &rule)

	// add rule to CachedAdRules
	CachedAdRules[geo] = rule.AdServer

	return rule.AdServer
}
