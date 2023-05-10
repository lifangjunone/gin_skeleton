package cmd

import (
	"fmt"
	"gin_skeleton/global"
	"github.com/spf13/cobra"
	"os"
)

var (
	printVersion bool
)

var RootCmd = &cobra.Command{
	Use:   "gin-ske",
	Short: "gin-skeleton",
	Long:  "Go gin skeleton, custom struct",
	RunE: func(cmd *cobra.Command, args []string) error {
		if printVersion {
			fmt.Println(global.Version)
			return nil
		}
		return cmd.Help()
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "V", false, "if you need to print version")
	RootCmd.PersistentFlags().StringVarP(&global.ConfigFilePath, "conf", "f", "/home/lifangjun/Desktop/go_study/gin_skeleton/conf/env/dev.toml", "config file path")
}
