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
	"fmt"
	"io"
	"os"
	"time"
)

func timeToString(t time.Time) string {
	if humanRead {
		return t.Format(time.RFC1123)
	} else {
		return t.Format(time.RFC3339Nano)
	}
}

func stringToTime(s string) (t time.Time, err error) {
	// attempt parsing both default and human-readable formats
	t, err = time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t, err = time.Parse(time.RFC1123, s)
		if err != nil {
			return time.Time{}, err
		}
	}

	return t, nil
}

type timeLogFile struct {
	name      string
	file      *os.File
	size      int64
	offset    int64
	linesSeen int
}

func (f *timeLogFile) open() (err error) {
	f.file, err = os.OpenFile(f.name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	fileStat, err := f.file.Stat()
	if err != nil {
		return err
	}

	f.size = fileStat.Size()

	return nil
}

func (f *timeLogFile) close() (err error) {
	return f.file.Close()
}

func (f *timeLogFile) addEntry(timeString string) (err error) {
	bytes, err := f.file.WriteString(timeString + "\n")
	if err != nil {
		return err
	}

	f.size += int64(bytes)

	return nil
}

func (f *timeLogFile) readEntry(lineNumber int) (entry string, err error) {
	_, err = f.file.Seek(f.offset, io.SeekStart)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(f.file)

	for f.offset < f.size {
		lineBytes, err := reader.ReadBytes('\n')
		if err != nil {
			return "", err
		}

		lineLen := len(lineBytes)
		f.offset += int64(lineLen)
		f.linesSeen++

		if lineBytes[lineLen-1] == '\n' {
			drop := 1
			if lineLen > 1 && lineBytes[lineLen-2] == '\r' {
				drop = 2
			}

			lineBytes = lineBytes[:lineLen-drop]
		}

		if f.linesSeen == lineNumber {
			return string(lineBytes), nil
		}
	}

	return "", fmt.Errorf("entry %d not found", lineNumber)
}
