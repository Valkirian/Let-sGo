package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var host string
var ports []int

var defaultports []int = []int{80,22,23,445,443}


func init() {
	rootCmd.AddCommand(portdiscoverycmd)
	portdiscoverycmd.Flags().StringVarP(&host, "host", "s", "localhost", "You must set the target")
	portdiscoverycmd.Flags().IntSliceVarP(&ports, "ports", "p", defaultports, "Set the port range to scan")
}

var portdiscoverycmd = &cobra.Command {
	Use: "portdiscovery",
	Short: "Do a Port Discoverty",
	Long: `Do a Port Discovery over TCP network against you host and ports`,
	Run: func (cmd *cobra.Command, args []string) {
		res, err := PortDiscover(host, ports...)
		if err != nil {
			fmt.Println("Error al lanzar el escaneo")
		}
		for _, status := range res {
			switch status.isopen {
			case true:
				fmt.Printf("Port %d is Open\n", status.port)
			case false:
				fmt.Printf("Port %d is Closed\n", status.port)
			}
		}
	},
}