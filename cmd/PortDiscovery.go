package cmd

import (
	"fmt"
	"log"
	"net"
)

type PortStatus struct {
	port int
	isopen bool
}

func PortDiscover(host string, PortRange ...int) ([]PortStatus, error){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Initialize the Slice of port responses
	ports := []PortStatus{}
	isok := false
	for _, prt := range PortRange {
		xnt := net.JoinHostPort(host,fmt.Sprintf("%v", prt))
		conn, err := net.Dial("tcp4", xnt)
		if err != nil {
			isok = false
			pd := PortStatus{
				port: prt,
				isopen: isok,
			}
			ports = append(ports, pd)
			continue
		}
		if conn != nil {
			defer conn.Close()
			isok = true
			pd := PortStatus{
				port: prt,
				isopen: isok,
			}
			ports = append(ports, pd)
		}
	}
	return ports, nil
}