package cmd

import (
	"gin_skeleton/conf"
	"gin_skeleton/global"
	"gin_skeleton/middlewares"
	"gin_skeleton/routers"
	"github.com/spf13/cobra"
)

var runServerCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Long:  "start server by command line",
	RunE: func(cmd *cobra.Command, args []string) error {
		// init engine
		engine := routers.InitEngine()

		// init config
		conf.InitConfig(global.ConfigFilePath)

		// init middlewares
		middlewares.InitMiddleware(engine)

		// init routers
		routers.InitRouter(engine)

		// start server
		engine.Run(conf.ConfObj.App.GetAddr())
		return nil
	},
}

func init() {
	RootCmd.AddCommand(runServerCmd)
}
