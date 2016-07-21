package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SERVICES_FILE = "/etc/services"

type Port struct {
	Proto string
	Name  string
}

func (p *Port) String() string {
	return fmt.Sprintf("{Proto:%v Name: %v}", p.Proto, p.Name)
}

type PortMap map[uint16]Port

func GetServices() (PortMap, error) {
	file, err := os.Open(SERVICES_FILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	port_map := make(PortMap)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			// ignore comments
			continue
		}
		line = strings.TrimSpace(line)
		split := strings.SplitN(line, "#", 2)
		fields := strings.Fields(split[0])
		if len(fields) < 2 {
			continue
		}
		name := fields[0]
		portproto := strings.SplitN(fields[1], "/", 2)
		port, err := strconv.ParseInt(portproto[0], 10, 32)
		if err != nil {
			panic(err)
		}
		proto := portproto[1]
		port_map[uint16(port)] = Port{
			Name:  name,
			Proto: proto,
		}
	}

	return port_map, nil
}
