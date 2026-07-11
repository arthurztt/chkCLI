package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	timeout int
	Version string
)

var rootCmd = &cobra.Command{
	Use:   "chk [urls...]",
	Short: "Checa o status de endpoints HTTP",
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func SetVersion(v string) {
	Version = v
	rootCmd.Version = v
}

func init() {
	rootCmd.Flags().IntVarP(&timeout, "timeout", "t", 5, "Aborta a requisição caso demora mais que o timeout definido. (default 5)")
}

func run(cmd *cobra.Command, args []string) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	for _, url := range args {
		start := time.Now()
		resp, err := client.Get(url)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("%-8s	%s   	 %s\n", err, elapsed.Round(time.Millisecond), url)
			continue
		}
		resp.Body.Close()

		fmt.Printf("%-3d		%-8s %s\n", resp.StatusCode, elapsed.Round(time.Millisecond), url)
	}
}
