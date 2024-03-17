/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/bootstrap/database"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/bootstrap/server"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/graph"
	handlerimpl "github.com/willjrcom/fullcycle-clean-architecture-golang/internal/infra/handler"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/infra/repository"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/usecase"
)

// restAndGraphqlCmd represents the restAndGraphql command
var restAndGraphqlCmd = &cobra.Command{
	Use:   "restAndGraphql",
	Short: "GraphQL and REST API server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restAndGraphql called")
		ctx := context.Background()

		serverInstance := server.NewServerChi()

		// Load database
		db, err := database.NewPostgreSQLConnection(ctx)

		if err != nil {
			panic(err)
		}

		orderRepository := repository.NewOrderRepositoryBun(db)
		orderUseCase := usecase.NewService(orderRepository)

		// REST API
		orderHandler := handlerimpl.NewHandlerOrder(orderUseCase)
		serverInstance.AddHandler(orderHandler)

		// GraphQL
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Service: orderUseCase}}))

		serverInstance.Router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		serverInstance.Router.Handle("/query", srv)

		// Start server
		if err := serverInstance.StartServer(":8080"); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(restAndGraphqlCmd)
}
