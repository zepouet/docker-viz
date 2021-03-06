package main

import (
	"github.com/spf13/cobra"
	"github.com/Treeptik/docker-viz/server"
	"github.com/Treeptik/docker-viz/dockertype"
	"fmt"
)

func main() {
	var port int
	var debug bool
	version := "0.8.2"

	// create version commande
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "docker-viz and docker engine version",
		Long: `docker-viz and docker engine version`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Docker-Viz version: " + version)
			fmt.Println("Docker Engine version: " + dockertype.DockerVersion())
		},
	}

	// create default root command
	var rootCmd = &cobra.Command{
		Use:   "docker-viz",
		Short: "docker-viz server is a web server for data visualization on Docker.",
		Long: `docker-viz server is a web server who return a data visualization for different informations on Docker containers and images.`,

		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer(port, debug);
		},
	}
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "docker-viz server port")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Run docker-viz server in \"debug\" mode")
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}