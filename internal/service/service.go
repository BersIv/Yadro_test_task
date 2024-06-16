package service

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
	if ok, _ := regexp.MatchString("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$", in.DnsServer); !ok {
		return &pb.AddDNSServerResponse{Status: false}, errors.New("mistake in ip adress")
	}
	resolvConf := "/etc/resolv.conf"
	file, err := os.OpenFile(resolvConf, os.O_RDWR, 0666)
	if err != nil {
		return &pb.AddDNSServerResponse{Status: false}, err
	}
	defer file.Close()
	exists, err := dnsServerExists(file, in.DnsServer)
	if err != nil {
		return &pb.AddDNSServerResponse{Status: false}, err
	}
	if exists {
		return &pb.AddDNSServerResponse{Status: false}, errors.New("ip already exists")
	}
	if _, err := file.WriteString(fmt.Sprintf("nameserver %s\n", in.DnsServer)); err != nil {
		return &pb.AddDNSServerResponse{Status: false}, err
	}

	return &pb.AddDNSServerResponse{Status: true}, nil
}

func (s *Server) RemoveDNSServer(ctx context.Context, in *pb.RemoveDNSServerRequest) (*pb.RemoveDNSServerResponse, error) {
	resolvConf := "/etc/resolv.conf"
	content, err := os.ReadFile(resolvConf)
	if err != nil {
		return &pb.RemoveDNSServerResponse{Status: false}, err
	}
	if ok, _ := regexp.MatchString("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$", in.DnsServer); !ok {
		return &pb.RemoveDNSServerResponse{Status: false}, errors.New("mistake in ip adress")
	}
	pattern := fmt.Sprintf("\\b%s\\b", regexp.QuoteMeta(in.DnsServer))
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return &pb.RemoveDNSServerResponse{Status: false}, err
	}
	lines := strings.Split(string(content), "\n")
	dnsServers := []string{}
	newLines := []string{}
	var exists bool
	for _, line := range lines {
		if !reg.MatchString(line) {
			newLines = append(newLines, line)
			if strings.HasPrefix(line, "nameserver ") {
				dnsServers = append(dnsServers, strings.TrimPrefix(line, "nameserver "))
			}
			continue
		}
		exists = true
	}
	if !exists {
		return &pb.RemoveDNSServerResponse{Status: false}, errors.New("ip not exists in conf")
	}
	newConf := strings.Join(newLines, "\n")
	err = os.WriteFile(resolvConf, []byte(newConf), 0666)
	if err != nil {
		return &pb.RemoveDNSServerResponse{Status: false}, err
	}

	return &pb.RemoveDNSServerResponse{Status: true, DnsServers: dnsServers}, nil
}

func dnsServerExists(file *os.File, dnsServer string) (bool, error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "nameserver ") {
			dns := strings.TrimSpace(strings.TrimPrefix(line, "nameserver"))
			if dns == dnsServer {
				return true, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}
