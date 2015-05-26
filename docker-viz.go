package main

import (
	"github.com/spf13/cobra"
	"github.com/Treeptik/docker-viz/server"
	"github.com/Treeptik/docker-viz/dockertype"
	"fmt"
)

func main() {
	var port int
	version := "0.7.0"

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "docker-viz and docker engine version",
		Long: `docker-viz and docker engine version`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Docker-Viz version: " + version)
			fmt.Println("Docker Engine version: " + dockertype.DockerVersion())
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "docker-viz",
		Short: "docker-viz server is a web server for data visualization on Docker.",
		Long: `docker-viz server is a web server who return a data visualization for different informations on Docker containers and images.`,

		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer(port);
		},
	}
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "docker-viz server port")
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}