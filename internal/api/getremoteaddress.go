package api

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"hospital-app/config"

	"github.com/labstack/echo/v4"
)

func GetRemoteAddress(c echo.Context) error {

	// fetch and verify the address
	_, ipAddress := TestIPAddress(GetOutboundIP())

	address := ""

	if ipAddress != "" {
		address = "http://" + ipAddress
	}

	return c.JSON(http.StatusOK, echo.Map{"address": address})
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}

func TestIPAddress(ipAddress string) (bool, string) {

	if ipAddress == "" {
		return false, ""
	}
	url := fmt.Sprintf("%v:%v", ipAddress, config.ServerPort)
	timeout := time.Duration(500 * time.Millisecond)
	_, err := net.DialTimeout("tcp", url, timeout)

	if err == nil {
		return true, url
	}

	return false, url
}
