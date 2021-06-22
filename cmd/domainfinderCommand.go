package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var wordlist_path string
var domain string

func init() {
	rootCmd.AddCommand(domainfindercmd)
	domainfindercmd.Flags().StringVarP(&wordlist_path, "wordlist", "w", "", "Use this flag to set the wordlist path")
	domainfindercmd.Flags().StringVarP(&domain, "domain", "d", "", "Use this flag to set the domain to be enumerated")
}

var domainfindercmd = &cobra.Command{
	Use: "subdomainfinder",
	Short: "Use this tool for subdomain find",
	Long: "Tool for subdomain enumeration using a wordlist",
	Run: func (cmd *cobra.Command, args []string) {
		res := SubdomainFinder(wordlist_path, domain)
		fmt.Println(res)
	},
}
