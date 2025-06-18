package csvloader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

// ParserFunc is a function that parses a row into type T.
type ParserFunc[T any] func([]string) (T, error)

// CSVDriver holds parsed entries keyed by identifier.
type CSVDriver[T any] struct {
	data map[string]T
}

// NewCSVDriver loads a CSV from path, validates headers, parses rows, and returns a CSVDriver.
func NewCSVDriver[T any](path string, keyExtractor func(T) string, parser ParserFunc[T]) (*CSVDriver[T], error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	headers, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("cannot read headers: %w", err)
	}

	expected, err := extractCSVHeaders(new(T))
	if err != nil {
		return nil, err
	}
	if !equalStringSlices(headers, expected) {
		return nil, fmt.Errorf("header mismatch:\nexpected: %v\ngot: %v", expected, headers)
	}

	data := make(map[string]T)
	for {
		row, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading row: %w", err)
		}

		entry, err := parser(row)
		if err != nil {
			return nil, fmt.Errorf("invalid row %v: %w", row, err)
		}

		data[keyExtractor(entry)] = entry
	}

	return &CSVDriver[T]{data: data}, nil
}

// Get returns the entry for a given key (and a bool indicating presence).
func (d *CSVDriver[T]) Get(key string) (T, bool) {
	val, ok := d.data[key]
	return val, ok
}

func extractCSVHeaders[T any](val *T) ([]string, error) {
	t := reflect.TypeOf(*val)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("expected a struct type")
	}
	var tags []string
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("csv")
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags, nil
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
