package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "test",
		Short: "test issue with local flags at the root",
		Run: func(_ *cobra.Command, _ []string) {
			log.Print("hello")
		},
	}

	_ = rootCmd.LocalFlags().BoolP("enable-thing", "w", false, "enables thing only for the root command")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
