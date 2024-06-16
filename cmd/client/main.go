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

var changeHostnameCmd = &cobra.Command{
	Use:   "change-hostname",
	Short: "Change hostname",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, conn, err := newClient()
		if err != nil {
			log.Fatalf("Could not create new channel: %v", err)
		}
		defer conn.Close()
		req := &pb.ChangeHostnameRequest{Name: args[0]}
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
	Args:  cobra.ExactArgs(0),
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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, conn, err := newClient()
		if err != nil {
			log.Fatalf("Couldn't create new channel: %v", err)
		}
		defer conn.Close()
		req := &pb.AddDNSServerRequest{DnsServer: args[0]}
		res, err := client.AddDNSServer(context.Background(), req)
		if err != nil {
			log.Fatalf("Couldn't add DNS server: %v", err)
		}
		fmt.Printf("Response:\nstatus: %v\n", res.Status)
	},
}

var removeDNSServer = &cobra.Command{
	Use:   "remove-dns",
	Short: "Remove DNS server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, conn, err := newClient()
		if err != nil {
			log.Fatalf("Couldn't create new channel: %v", err)
		}
		defer conn.Close()
		req := &pb.RemoveDNSServerRequest{DnsServer: args[0]}
		res, err := client.RemoveDNSServer(context.Background(), req)
		if err != nil {
			log.Fatalf("Couldn't remove DNS server: %v", err)
		}
		fmt.Printf("Response:\nstatus: %v\n", res.Status)
		fmt.Println("DNS servers:")
		for _, server := range res.DnsServers {
			fmt.Println(server)
		}
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
	rootCmd.AddCommand(changeHostnameCmd, listDnsCmd, addDNSServer, removeDNSServer)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
