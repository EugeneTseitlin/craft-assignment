package internal

import (
	"testing"
)

func TestReadCsvFileReadsValuesProperly(t *testing.T) {
	path := "testdata/incremental.csv"
	expected := [][]string{
		{"header1", "header2", "header3"},
		{"value1", "value2", "value3"},
		{"value4", "value5", "value6"},
	}

	records, err := ReadCsvFile(path)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(records) != len(expected) {
		t.Errorf("Unexpected number of records. Expected: %d, Got: %d", len(expected), len(records))
	}

	for i, record := range records {
		if len(record) != len(expected[i]) {
			t.Errorf("Unexpected number of fields in record %d. Expected: %d, Got: %d", i, len(expected[i]), len(record))
		}

		for j, field := range record {
			if field != expected[i][j] {
				t.Errorf("Unexpected field value at (%d, %d). Expected: %s, Got: %s", i, j, expected[i][j], field)
			}
		}
	}
}

func TestReadCsvFileNonexistingFile(t *testing.T) {
	nonexistingFilePath := "testdata/unicorn.csv"
	_, err := ReadCsvFile(nonexistingFilePath)
	if err == nil {
		t.Error("Expected an error for nonexisting file")
	}
}

func TestReadCsvFileEmpty(t *testing.T) {
	emptyPath := "testdata/empty.csv"
	emptyExpected := [][]string{}
	emptyRecords, err := ReadCsvFile(emptyPath)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(emptyRecords) != len(emptyExpected) {
		t.Errorf("Unexpected number of records. Expected: %d, Got: %d", len(emptyExpected), len(emptyRecords))
	}
} 

