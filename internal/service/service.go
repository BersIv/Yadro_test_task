package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	pb "yadro_test_task/gen/go/hostconfig"
)

type Server struct {
	pb.UnimplementedHostConfigServer
}

func (s *Server) ChangeHostname(ctx context.Context, in *pb.ChangeHostnameRequest) (*pb.ChangeHostnameResponse, error) {
	cmd := exec.Command("hostnamectl", "set-hostname", in.Name)
	err := cmd.Run()
	if err != nil {
		return &pb.ChangeHostnameResponse{Status: false}, err
	}
	hostname, err := os.Hostname()
	if err != nil {
		return &pb.ChangeHostnameResponse{Status: false}, err
	}
	if hostname != in.Name {
		return &pb.ChangeHostnameResponse{Status: false}, errors.New("hostname not changed")
	}
	return &pb.ChangeHostnameResponse{Status: true, NewName: hostname}, nil
}

func (s *Server) ListDNSServers(ctx context.Context, in *pb.ListDNSServersRequest) (*pb.ListDNSServersResponse, error) {
	resolvConf := "/etc/resolv.conf"
	content, err := os.ReadFile(resolvConf)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	dnsServers := []string{}
	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver ") {
			dnsServers = append(dnsServers, strings.TrimPrefix(line, "nameserver "))
		}
	}
	return &pb.ListDNSServersResponse{DnsServers: dnsServers}, nil
}

func (s *Server) AddDNSServer(ctx context.Context, in *pb.AddDNSServerRequest) (*pb.AddDNSServerResponse, error) {
	if in.DnsServer == "" {
		return &pb.AddDNSServerResponse{Status: false}, errors.New("dns server cannot be empty")
	}
	resolvConf := "/etc/resolv.conf"
	file, err := os.OpenFile(resolvConf, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return &pb.AddDNSServerResponse{Status: false}, err
	}
	defer file.Close()
	if _, err := file.WriteString(fmt.Sprintf("\nnameserver %s", in.DnsServer)); err != nil {
		return &pb.AddDNSServerResponse{Status: false}, err
	}

	return &pb.AddDNSServerResponse{Status: true}, nil
}
