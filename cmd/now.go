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
	noEntry bool

	// nowCmd represents the now command
	nowCmd = &cobra.Command{
		Use:   "now",
		Short: "Record the current time",
		Long: `Record the current time.

This is the same behavior for clkin when invoked by default.`,
		Run: func(cmd *cobra.Command, args []string) {
			recordNow()
		},
	}
)

func init() {
	rootCmd.AddCommand(nowCmd)

	nowCmd.PersistentFlags().BoolVar(&noEntry, "no-entry", false, "do not enter time in time log")
}

func recordNow() {
	now := time.Now()
	nowString := getTimeString(now)

	fmt.Println("Current time is: ", nowString)

	if noEntry {
		return
	}

	f, err := os.OpenFile(timeLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cobra.CheckErr(err)

	defer f.Close()

	_, err = f.WriteString(nowString + "\n")
	cobra.CheckErr(err)
}
