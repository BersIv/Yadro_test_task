package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "yadro_test_task/gen/go/hostconfig"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var rootCmd = &cobra.Command{
	Use:   "hostconfig",
	Short: "hostconfig - edit dns and hostname",
}
var hostname, dnsServer string

var changeHostnameCmd = &cobra.Command{
	Use:   "change-hostname",
	Short: "Change hostname",
	Run: func(cmd *cobra.Command, args []string) {
		client, conn, err := newClient()
		if err != nil {
			log.Fatalf("Could not create new channel: %v", err)
		}
		defer conn.Close()
		req := &pb.ChangeHostnameRequest{Name: hostname}
		res, err := client.ChangeHostname(context.Background(), req)
		if err != nil {
			log.Fatalf("Error setting hostname: %v", err)
		}
		fmt.Printf("Response:\nstatus: %v\nnew name:%v\n", res.Status, res.NewName)
	},
}

var listDnsCmd = &cobra.Command{
	Use:   "list-dns",
	Short: "List all DNS servers",
	Run: func(cmd *cobra.Command, args []string) {
		client, conn, err := newClient()
		if err != nil {
			log.Fatalf("Couldn't create new channel: %v", err)
		}
		defer conn.Close()
		req := &pb.ListDNSServersRequest{}
		res, err := client.ListDNSServers(context.Background(), req)
		if err != nil {
			log.Fatalf("Couldn't list DNS servers: %v", err)
		}
		fmt.Println("DNS servers:")
		for _, server := range res.DnsServers {
			fmt.Println(server)
		}
	},
}

var addDNSServer = &cobra.Command{
	Use:   "add-dns",
	Short: "Add new DNS server",
	Run: func(cmd *cobra.Command, args []string) {
		client, conn, err := newClient()
		if err != nil {
			log.Fatalf("Couldn't create new channel: %v", err)
		}
		defer conn.Close()
		req := &pb.AddDNSServerRequest{DnsServer: dnsServer}
		res, err := client.AddDNSServer(context.Background(), req)
		if err != nil {
			log.Fatalf("Couldn't add DNS server: %v", err)
		}
		fmt.Printf("Response:\nstatus: %v\n", res.Status)

	},
}

func newClient() (pb.HostConfigClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewHostConfigClient(conn)
	return client, conn, nil
}

func init() {
	changeHostnameCmd.Flags().StringVarP(&hostname, "hostname", "n", "", "New hostname")
	changeHostnameCmd.MarkFlagRequired("hostname")
	addDNSServer.Flags().StringVarP(&dnsServer, "dnsServer", "a", "", "New dns server")
	addDNSServer.MarkFlagRequired("dnsServer")
	rootCmd.AddCommand(changeHostnameCmd, listDnsCmd, addDNSServer)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
