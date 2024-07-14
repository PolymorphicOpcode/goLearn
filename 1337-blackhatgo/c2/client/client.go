package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"c2/grpcapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		conn   *grpc.ClientConn
		err    error
		client grpcapi.AdminClient
	)

	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use insecure credentials for the connection
	conn, err = grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", 9090), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client = grpcapi.NewAdminClient(conn)

	var cmd = new(grpcapi.Command)
	cmd.In = os.Args[1]

	// Create a new background context for the command execution
	cmdCtx := context.Background()

	cmd, err = client.RunCommand(cmdCtx, cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Out)
}
