// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/Code-Hex/battery-server/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type BatteryServer struct {
	port    string
	path    string
	logMode bool
	command *cobra.Command
}

func New() *BatteryServer {
	btserver := &BatteryServer{
		port: "4000",
		command: &cobra.Command{
			Use:   "battery-server",
			Short: "battery-server",
			Long:  `これはバッテリー情報を表示するサーバーだよー`,
		},
	}

	btserver.command.RunE = btserver.RunServer

	// Add flags for server mode
	btserver.command.Flags().StringVarP(&btserver.path, "path", "", "bt-server.log", "log output path")
	btserver.command.Flags().StringVarP(&btserver.port, "port", "p", "4000", "port number")
	btserver.command.Flags().BoolVarP(&btserver.logMode, "logger", "", false, "logger mode")

	// Add sub commands
	btserver.command.AddCommand(CliNew().CliCmdNew())
	btserver.command.AddCommand(CliNew().CliCmdNew())

	return btserver
}

func (b *BatteryServer) Execute() {
	if err := b.command.Execute(); err != nil {
		os.Exit(-1)
	}
}

func (b *BatteryServer) RunServer(cmd *cobra.Command, args []string) error {
	e := echo.New()
	if b.logMode {
		log, err := os.Create(b.path)
		if err != nil {
			return errors.Wrapf(err, "Could not run the battery server")
		}

		middleware.DefaultLoggerConfig.Output = log
		e.Use(middleware.Logger())
	}

	e.GET("/", route.HealthCheck)
	e.GET("/battery", route.ShowBattery)

	e.Run(standard.New(":" + b.port))

	return nil
}
