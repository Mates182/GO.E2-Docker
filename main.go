package main

import (
	"html/template"
	"net"
	"net/http"
)

type PageData struct {
	LocalIP string
}

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
	ip := getLocalIP()

	data := PageData{
		LocalIP: ip,
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
