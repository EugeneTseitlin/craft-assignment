package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"merge-csv/internal"
)

var (
	input      []string
	joinColumn string
	output     string

	rootCmd = &cobra.Command{
		Use:   "merge-csv",
		Short: "merge-csv is a tool to merge csv files",
		Run: func(cmd *cobra.Command, args []string) {
			files := make([][][]string, len(input))
			for i, path := range input {
				file, err := internal.ReadCsvFile(path)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				files[i] = file
			}

			merged, err := internal.MergeFiles(files, joinColumn)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			err = internal.WriteCsvFile(output, merged)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.Flags().StringArrayVar(&input, "input", []string{}, "input csv files")
	rootCmd.Flags().StringVar(&joinColumn, "join-column", "", "join column name")
	rootCmd.Flags().StringVar(&output, "output", "", "output csv file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

