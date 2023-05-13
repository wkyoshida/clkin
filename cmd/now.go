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
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := recordNow(cmd, args)
			cobra.CheckErr(err)
		},
	}
)

func init() {
	nowCmd.PersistentFlags().BoolVar(&noEntry, "no-entry", false, "do not enter time in time log")
}

func recordNow(cmd *cobra.Command, args []string) (err error) {
	now := time.Now()
	nowString := timeToString(now)

	fmt.Println("Current time is: ", nowString)

	if noEntry {
		return nil
	}

	return timeLog.addEntry(nowString)
}
