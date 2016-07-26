package services

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

import (
	"testing"
)

func TestGetServices(t *testing.T) {
	svc, err := GetServices()
	if err != nil {
		t.Fatalf("Did not expect to get an error")
	}
	if !(svc[22].Proto == "tcp" || svc[22].Proto == "udp" || svc[22].Name == "ssh") {
		t.Fatalf("Expected to find SSH, got %v", svc[22])
	}
}
