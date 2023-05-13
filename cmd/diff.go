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

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	fromLine  int
	toLine    int
	enterDiff bool

	// diffCmd represents the diff command
	diffCmd = &cobra.Command{
		Use:   "diff",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: func(cmd *cobra.Command, args []string) error {
			err := cobra.NoArgs(cmd, args)
			if err != nil {
				return err
			}

			err = validateDiffFlags(cmd, args)
			if err != nil {
				return err
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("diff called")
		},
	}
)

func init() {
	diffCmd.PersistentFlags().IntVarP(&fromLine, "from", "f", 0, "file line-number of starting time (default to last entry)")
	diffCmd.PersistentFlags().IntVarP(&toLine, "to", "t", 0, "file line-number of ending time (default to current time)")
	diffCmd.PersistentFlags().BoolVar(&enterDiff, "enter", false, "enter elapsed time in time log")
}

func validateDiffFlags(cmd *cobra.Command, args []string) error {
	fromChanged := cmd.Flag("from").Changed
	toChanged := cmd.Flag("to").Changed

	if (fromChanged && fromLine <= 0) || (toChanged && toLine <= 0) {
		return fmt.Errorf("invalid value: --from and --to flags must be positive integers")
	}
	if fromLine > toLine {
		return fmt.Errorf("invalid value: --from cannot be an entry after --to")
	}

	return nil
}
