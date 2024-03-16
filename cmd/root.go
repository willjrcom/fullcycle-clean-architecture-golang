/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "clean architecture",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println("httpserver called")
		ctx := context.Background()

		serverInstance := server.NewServerChi()

		// Load database
		db, err := database.NewPostgreSQLConnection(ctx)

		if err != nil {
			panic(err)
		}

		orderRepository := repository.NewOrderRepositoryBun(db)
		orderUseCase := usecase.NewService(orderRepository)
		orderHandler := handlerimpl.NewHandlerOrder(orderUseCase)
		serverInstance.AddHandler(orderHandler)

		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Service: orderUseCase}}))

		serverInstance.Router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		serverInstance.Router.Handle("/query", srv)

		if err := serverInstance.StartServer(":8080"); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
