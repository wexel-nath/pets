package database

import (
	"database/sql"
)

func ScanRowToMap(row *sql.Row, colNames []string) (map[string]interface{}, error) {
	colPointers := colPointers(len(colNames))
	err := row.Scan(colPointers...)
	if err != nil {
		return nil, err
	}

	m := colPointersToMap(colNames, colPointers)

	return m, nil
}

func ScanRowsToMap(rows *sql.Rows, colNames []string) ([]map[string]interface{}, error) {
	var s []map[string]interface{}

	for rows.Next() {
		colPointers := colPointers(len(colNames))
		err := rows.Scan(colPointers...)
		if err != nil {
			return nil, err
		}

		m := colPointersToMap(colNames, colPointers)
		s = append(s, m)
	}

	return s, nil
}

func colPointers(columnCount int) []interface{} {
	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	columns := make([]interface{}, columnCount)
	columnPointers := make([]interface{}, columnCount)
	for i := range columns {
		columnPointers[i] = &columns[i]
	}
	return columnPointers
}

func colPointersToMap(columnNames []string, columnPointers []interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i, colName := range columnNames {
		val := columnPointers[i].(*interface{})
		m[colName] = *val
	}
	return m
}
