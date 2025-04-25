package main

import (
	"fmt"
	"os"

	"github.com/plume-protocol/plume-db/tools/cmd/plumedb/benchmark"
	"github.com/plume-protocol/plume-db/tools/cmd/plumedb/operations"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "plumedb",
		Short: "A tool to generate raw key value data from a node as well as benchmark different backends",
	}

	rootCmd.AddCommand(
		benchmark.GenerateCmd(),
		benchmark.DBWriteCmd(),
		benchmark.DBRandomReadCmd(),
		benchmark.DBIterationCmd(),
		benchmark.DBReverseIterationCmd(),
		operations.DumpDbCmd(),
		operations.PruneCmd(),
		operations.DumpIAVLCmd(),
		operations.StateSizeCmd(),
		operations.ReplayChangelogCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
