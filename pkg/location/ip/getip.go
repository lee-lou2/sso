package ip

import (
	"net"
	"sso/pkg/http"
)

// GetLocalIP 내부 IP 조회
func GetLocalIP() []string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return []string{}
	}
	IPs := make([]string, 0)
	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				IPs = append(IPs, ipNet.IP.To4().String())
			}
		}
	}
	return IPs
}

// GetOutboundIP 외부 IP 조회
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "0.0.0.0"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

// GetPublicIP 공개 IP 조회
func GetPublicIP() string {
	resp, err := http.Request("GET", "https://ifconfig.me", nil)
	if err != nil || resp.Body == "" {
		return "0.0.0.0"
	}
	return resp.Body
}
