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
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	humanRead bool
	timeLog   string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "clkin",
		Short: "Simple time tracking with a time log file",
		Long: `Simple time tracking with a time log file.

CLKIN records timestamps into a file and allows operations, 
such as finding the elapsed time between records.`,
		Version: "0.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			now := time.Now()

			var nowString string
			if humanRead {
				nowString = now.Format(time.RFC1123)
			} else {
				nowString = now.String()
			}

			fmt.Println("Current time is: ", nowString)

			f, err := os.OpenFile(timeLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			cobra.CheckErr(err)

			defer f.Close()

			_, err = f.WriteString(nowString + "\n")
			cobra.CheckErr(err)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&humanRead, "human", false, "use human-readable time format")
	rootCmd.PersistentFlags().StringVarP(&timeLog, "timelog", "l", ".clkin.log", "file path to time log")
}
