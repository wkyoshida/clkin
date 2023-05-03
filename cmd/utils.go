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
	"bufio"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func getTimeString(timestamp time.Time) string {
	if humanRead {
		return timestamp.Format(time.RFC1123)
	} else {
		return timestamp.Format(time.RFC3339Nano)
	}
}

func getLogEntryTime(f *os.File, lineNumber int) time.Time {
	entry, err := readLogEntry(f, lineNumber)
	cobra.CheckErr(err)

	// attempt parsing both default and human-readable formats
	entryTime, err := time.Parse(time.RFC3339Nano, entry)
	if err != nil {
		entryTime, err = time.Parse(time.RFC1123, entry)
	}
	cobra.CheckErr(err)

	return entryTime
}

func readLogEntry(f *os.File, lineNumber int) (line string, err error) {
	scanner := bufio.NewScanner(f)
	var linesSeen int

	for scanner.Scan() {
		linesSeen++

		if linesSeen == lineNumber {
			return scanner.Text(), nil
		}
	}

	return line, scanner.Err()
}
