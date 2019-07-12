package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/k8sland/lab2/scheduler/party"
	"github.com/spf13/cobra"
)

const app = "PartyScheduler"

var (
	// Version set via build tags
	Version = ""

	rootCmd = &cobra.Command{
		Use:   strings.ToLower(app),
		Short: "Schedules pods based on costumes",
		Long:  "Schedules pods based on costumes",
		Run:   listen,
	}
)

func init() {
	rootCmd.Version = Version
}

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listen(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go setupSignal(cancel)

	party.NewScheduler().Run(ctx)
}

func setupSignal(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-signalChan:
			log.Printf("Shutdown signal received, exiting...")
			cancel()
			os.Exit(0)
		}
	}
}
