package main

import s "strings"
import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func getClientIPall(rw http.ResponseWriter, req *http.Request) {

	sip1 := req.Header.Get("X-Forwarded-For")
	sip2 := req.Header.Get("http_client_ip")
	sip3 := req.Header.Get("proxy-client-iP")
	sip4 := req.Header.Get("wl-proxy-client-ip")
	sip5 := req.Header.Get("http_x_forwarded_for")
	sip6 := req.Header.Get("http_via")
	sip7 := req.Header.Get("X-Real-IP")
	sip8 := req.Header.Get("X_FORWARDED_FOR")
	sip9 := req.RemoteAddr

	io.WriteString(rw, "1: "+sip1)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "2: "+sip2)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "3: "+sip3)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "4: "+sip4)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "5: "+sip5)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "6: "+sip6)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "7: "+sip7)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "8: "+sip8)
	io.WriteString(rw, "\r\n")
	io.WriteString(rw, "9: "+sip9)
	io.WriteString(rw, "\r\n")

}

func getClientIP(rw http.ResponseWriter, req *http.Request) {
	// sip := s.Split(req.RemoteAddr, ":")
	// io.WriteString(rw, sip[0])
	var myip []string
	sip := req.Header.Get("x-forwarded-for")
	if sip == "" {
		sip = req.Header.Get("http_client_ip")
	}
	if sip == "" {
		sip = req.Header.Get("Proxy-Client-IP")
	}
	if sip == "" {
		sip = req.Header.Get("WL-Proxy-Client-IP")
	}
	if sip == "" {
		sip = req.Header.Get("HTTP_X_FORWARDED_FOR")
	}
	if sip == "" {
		sip = req.RemoteAddr
	}
	if sip != "" {
		myip = s.Split(sip, ":")
	}

	io.WriteString(rw, myip[0])

}

func getClientIP2(rw http.ResponseWriter, req *http.Request) {
	sip := req.Header.Get("x-forwarded-for")
	if sip == "" {
		sip = req.RemoteAddr
	}
	io.WriteString(rw, sip)
}

func main() {
	http.HandleFunc("/getClientIPall", getClientIPall)
	http.HandleFunc("/getClientIP", getClientIP)
	http.ListenAndServe(":8080", nil)
}

func get_external() string {
	resp, _ := http.Get("http://myexternalip.com/raw")
	defer resp.Body.Close()
	retdata, _ := ioutil.ReadAll(resp.Body)

	return string(retdata)

}

func get_internal() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	os.Exit(0)
}
