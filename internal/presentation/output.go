package presentation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"unicode"
)

func Print(w io.Writer, result any, format string, columns []Column) error {
	if result == nil {
		if format == "json" {
			_, err := fmt.Fprintln(w, "{}")
			return err
		}
		_, err := fmt.Fprintln(w, "Done.")
		return err
	}
	encoded, err := json.Marshal(result)
	if err != nil {
		return err
	}
	if format == "json" {
		var out bytes.Buffer
		if err := json.Indent(&out, encoded, "", "  "); err != nil {
			return err
		}
		_, err := fmt.Fprintln(w, out.String())
		return err
	}
	var envelope struct {
		Data json.RawMessage `json:"data"`
		Meta json.RawMessage `json:"meta"`
	}
	if err := json.Unmarshal(encoded, &envelope); err != nil || len(envelope.Data) == 0 {
		return errors.New("Fluxa response has no data envelope")
	}
	var rows []map[string]json.RawMessage
	if envelope.Data[0] == '[' {
		if err := json.Unmarshal(envelope.Data, &rows); err != nil {
			return err
		}
	} else {
		var row map[string]json.RawMessage
		if err := json.Unmarshal(envelope.Data, &row); err != nil {
			return err
		}
		rows = []map[string]json.RawMessage{row}
	}
	if len(rows) == 0 {
		_, err := fmt.Fprintln(w, "No records.")
		return err
	}
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	header := make([]string, len(columns))
	for i, col := range columns {
		header[i] = col.Name
	}
	if _, err := fmt.Fprintln(tw, strings.Join(header, "\t")); err != nil {
		return err
	}
	for _, row := range rows {
		values := make([]string, len(columns))
		for i, col := range columns {
			values[i] = displayValue(row[col.Key])
		}
		if _, err := fmt.Fprintln(tw, strings.Join(values, "\t")); err != nil {
			return err
		}
	}
	if err := tw.Flush(); err != nil {
		return err
	}
	if len(envelope.Meta) > 0 {
		var meta struct {
			Page  int `json:"page"`
			Pages int `json:"pages"`
			Total int `json:"total"`
		}
		if json.Unmarshal(envelope.Meta, &meta) == nil {
			_, err := fmt.Fprintf(w, "Page %d/%d, %d total\n", meta.Page, meta.Pages, meta.Total)
			return err
		}
	}
	return nil
}

func displayValue(raw json.RawMessage) string {
	if len(raw) == 0 || string(raw) == "null" {
		return "-"
	}
	var value string
	if json.Unmarshal(raw, &value) == nil {
		return strings.Map(func(r rune) rune {
			if unicode.IsControl(r) {
				return ' '
			}
			return r
		}, value)
	}
	return string(raw)
}
