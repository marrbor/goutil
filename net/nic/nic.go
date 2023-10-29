package nic

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

// GetIP returns IP address of this reporter
func GetIP() (string, error) {
	// gather ip addresses.
	as, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	// seek ip address that is not a loop back.
	for _, a := range as {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil // got it.
			}
		}
	}
	return "", fmt.Errorf("effective adress not found")
}

// GetInterface returns NIC for given IP.
func GetInterface(ip string) (*net.Interface, error) {
	// gather interfaces.
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// seek network interface bind for given IP
	for _, i := range ifs {
		if a, err := i.Addrs(); err == nil {
			for _, addr := range a {
				if strings.Contains(addr.String(), ip) {
					return &i, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("NIC not found")
}

// GetMacAddress returns mac address for given NIC.
func GetMacAddress(nic *net.Interface) (string, error) {
	// extract the hardware information base on the interface name capture above
	i, err := net.InterfaceByName(nic.Name)
	if err != nil {
		return "", err
	}
	hwa := i.HardwareAddr
	mac := hwa.String()
	return mac, nil
}

// RegexpMacAddressType express the regular expression that detect mac address strings.
var RegexpMacAddressType = regexp.MustCompile("^([[:xdigit:]]{2}[:.-]?){5}[[:xdigit:]]{2}$")

// ValidateMacAddress returns whether given string suit for mac address or not.
func ValidateMacAddress(mac string) bool {
	return RegexpMacAddressType.Match([]byte(mac))
}
