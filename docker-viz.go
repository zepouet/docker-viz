package main

import (
	"github.com/spf13/cobra"
	"github.com/Treeptik/docker-viz/server"
)

func main() {
	var port int

	var DockerVizCmd = &cobra.Command{
		Use:   "server",
		Short: "docker-viz server is a web server for data visualization on Docker.",
		Long: `docker-viz server is a web server who return a data visualization for different informations on Docker containers and images.
		       Complete documentation and source code is available at https://github.com/Treeptik/docker-viz`,

		Run: func(cmd *cobra.Command, args []string) {
			server.startServer(port);
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(DockerVizCmd)
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	rootCmd.Execute()
}