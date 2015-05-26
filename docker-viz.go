package main

import (
	"github.com/spf13/cobra"
	"github.com/Treeptik/docker-viz/server"
)

func main() {
	var port int

	var rootCmd = &cobra.Command{
		Use:   "docker-viz",
		Short: "docker-viz server is a web server for data visualization on Docker.",
		Long: `docker-viz server is a web server who return a data visualization for different informations on Docker containers and images.`,

		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer(port);
		},
	}
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	rootCmd.Execute()
}