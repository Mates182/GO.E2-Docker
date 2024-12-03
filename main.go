package main

import (
	"fmt"
	"net"
	"net/http"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "Unknown"
}

func handler(w http.ResponseWriter, r *http.Request) {
	localIP := getLocalIP()

	fmt.Fprintf(w, "<h1>Welcome, this page is deployed on Docker</h1>")
	fmt.Fprintf(w, "<h1>By: Mateo Pillajo :D</h1>")
	fmt.Fprintf(w, "<h1>Local IP: %s</h1>", localIP)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
