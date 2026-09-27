package cmd

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func addFileFlag(cmd *cobra.Command) {
	cmd.Flags().String("file", "", "JSON request file, or - for stdin")
}

func readBody(cmd *cobra.Command) ([]byte, error) {
	path, _ := cmd.Flags().GetString("file")
	if path == "" {
		return nil, errors.New("provide a JSON object with --file PATH or --file - for stdin")
	}
	var reader io.Reader
	if path == "-" {
		reader = cmd.InOrStdin()
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		reader = file
	}
	content, err := io.ReadAll(io.LimitReader(reader, 1<<20+1))
	if err != nil {
		return nil, err
	}
	if len(content) > 1<<20 {
		return nil, errors.New("JSON body exceeds 1 MiB")
	}
	return content, nil
}

func addYesFlag(cmd *cobra.Command) { cmd.Flags().Bool("yes", false, "confirm this action") }

func requireYes(cmd *cobra.Command) error {
	yes, _ := cmd.Flags().GetBool("yes")
	if !yes {
		return errors.New("pass --yes to confirm this action")
	}
	return nil
}
