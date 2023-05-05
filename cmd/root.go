/*
Copyright © 2023 wkyoshida

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	humanRead bool
	timeLog   timeLogFile

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "clkin",
		Short: "Simple time tracking with a time log file",
		Long: `Simple time tracking with a time log file.

CLKIN records timestamps into a file and allows operations, 
such as finding the elapsed time between records.`,
		Version: "0.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			recordNow()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(openTimeLog)

	rootCmd.PersistentFlags().BoolVar(&humanRead, "human", false, "use human-readable time format")
	rootCmd.PersistentFlags().StringVarP(&timeLog.name, "timelog", "l", ".clkin.log", "file path to time log")

	rootCmd.AddCommand(diffCmd)
	rootCmd.AddCommand(nowCmd)
	rootCmd.AddCommand(versionCmd)

	cobra.OnFinalize(closeTimeLog)
}

func openTimeLog() {
	err := timeLog.open()
	cobra.CheckErr(err)
}

func closeTimeLog() {
	err := timeLog.close()
	cobra.CheckErr(err)
}
