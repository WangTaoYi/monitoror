package system

import (
	"net"

	"golang.org/x/net/icmp"
)

func IsRawSocketAvailable() bool {
	_, err := icmp.ListenPacket("ip4:icmp", "")
	return err == nil
}

// ListLocalhostIpv4 list IP of every local network interfaces
func ListLocalhostIpv4() ([]string, error) {
	return listLocalhostIpv4(net.Interfaces)
}
func listLocalhostIpv4(listInterfaces func() ([]net.Interface, error)) ([]string, error) {
	var ips []string
	ifaces, err := listInterfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok && ip.IP.To4() != nil {
				ips = append(ips, ip.IP.String())
			}
		}
	}

	return ips, nil
}
