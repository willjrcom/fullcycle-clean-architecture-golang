/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/bootstrap/database"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/infra/pb"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/infra/repository"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc called")
		ctx := context.Background()

		// Load database
		db, err := database.NewPostgreSQLConnection(ctx)

		if err != nil {
			panic(err)
		}

		orderRepository := repository.NewOrderRepositoryBun(db)
		orderUseCase := usecase.NewService(orderRepository)

		// gRPC
		serviceGrpc := pb.NewServiceGrpc(orderUseCase)
		grpcServer := grpc.NewServer()
		pb.RegisterOrderServiceServer(grpcServer, serviceGrpc)
		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
