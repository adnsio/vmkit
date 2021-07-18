// Spin up Linux VMs with QEMU and Apple virtualization framework
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	"github.com/adnsio/vmkit/pkg/engine"
	"github.com/spf13/cobra"
)

func newImagesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "List images",
		Use:   "images",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runImages(); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runImages() error {
	eng, err := engine.New(&engine.NewOptions{
		Writer: os.Stderr,
	})
	if err != nil {
		return err
	}

	imgs := eng.ListImages()

	tableRows := make([][]string, 0, len(imgs))
	for _, image := range imgs {
		pulled := "No"

		imgPulled, err := image.Pulled()
		if err != nil {
			return err
		}

		if imgPulled {
			pulled = "Yes"
		}

		tableRows = append(tableRows, []string{
			image.Name,
			image.Version,
			image.Description,
			pulled,
		})
	}

	writeTable(&writeTableOptions{
		writer: os.Stdout,
		header: []string{
			"Name",
			"Version",
			"Description",
			"Pulled",
		},
		rows: tableRows,
	})

	return nil
}