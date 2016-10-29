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
	"fmt"

	"github.com/Code-Hex/battery-server/battery"
	"github.com/Code-Hex/battery-server/route"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Cli struct {
	command *cobra.Command
}

func CliNew() *Cli {
	return &Cli{}
}

func (c *Cli) CliCmdNew() *cobra.Command {
	return &cobra.Command{
		Use:   "cli",
		Short: "battery-server cli mode",
		Long:  `cli でバッテリー情報を表示するよー`,
		RunE:  c.GetBatteryInfoCli,
	}
}

func (c *Cli) GetBatteryInfoCli(cmd *cobra.Command, args []string) error {
	var b route.BTInfo
	var err error
	b.Percent, b.IsPowerd, err = battery.BatteryInfo()
	if err != nil {
		return errors.Wrapf(err, "Could not get battery info")
	}

	fmt.Println(b.Percent, b.IsPowerd)
	return nil
}
