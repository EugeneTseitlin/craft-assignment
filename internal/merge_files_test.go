package internal

import (
	"fmt"
	"testing"
)

// TestMergeFileTwoInputsMultipleColumns tests the case when the input files have multiple columns
func TestMergeFilesTwoInputsMultipleColumns(t *testing.T) {

	input1 := [][]string{
		{"id", "name", "age"},
		{"1", "John", "20"},
		{"2", "Jane", "21"},
		{"3", "Jack", "22"},
	}

	input2 := [][]string{
		{"id", "city", "country"},
		{"1", "Tokyo", "Japan"},
		{"2", "Warsaw", "Poland"},
		{"3", "Washington", "USA"},
	}

	expected := [][]string{
		{"id", "name", "age", "city", "country"},
		{"1", "John", "20", "Tokyo", "Japan"},
		{"2", "Jane", "21", "Warsaw", "Poland"},
		{"3", "Jack", "22", "Washington", "USA"},
	}

	result, err := MergeFiles([][][]string{input1, input2}, "id")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d rows, got %d", len(expected), len(result))
	}

	assertResultFile(t, result, expected)
}

// TestMergeFilesJoinColumnNotFirst tests the case when the join column is not the first column
func TestMergeFilesJoinColumnNotFirst(t *testing.T) {
	input1 := [][]string{
		{"name", "id", "age"},
		{"John", "1", "20"},
		{"Jane", "2", "21"},
		{"Jack", "3", "22"},
	}

	input2 := [][]string{
		{"city", "country", "id"},
		{"Tokyo", "Japan", "1"},
		{"Warsaw", "Poland", "2"},
		{"Washington", "USA", "3"},
	}

	expected := [][]string{
		{"id", "name", "age", "city", "country"},
		{"1", "John", "20", "Tokyo", "Japan"},
		{"2", "Jane", "21", "Warsaw", "Poland"},
		{"3", "Jack", "22", "Washington", "USA"},
	}

	result, err := MergeFiles([][][]string{input1, input2}, "id")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assertResultFile(t, result, expected)
}


// TestMergeFilesOrderByFirstInput ensures that the output is ordered by the first input file
func TestMergeFilesOrderByFirstInput(t *testing.T) {
	input1 := [][]string{
		{"name", "id", "age"},
		{"John", "1", "20"},
		{"Jane", "2", "21"},
		{"Jack", "3", "22"},
	}

	input2 := [][]string{
		{"city", "country", "id"},
		{"Warsaw", "Poland", "2"},
		{"Washington", "USA", "3"},
		{"Tokyo", "Japan", "1"},
	}

	expected := [][]string{
		{"id", "name", "age", "city", "country"},
		{"1", "John", "20", "Tokyo", "Japan"},
		{"2", "Jane", "21", "Warsaw", "Poland"},
		{"3", "Jack", "22", "Washington", "USA"},
	}

	result, err := MergeFiles([][][]string{input1, input2}, "id")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assertResultFile(t, result, expected)
}

func assertResultFile(t *testing.T, result [][]string, expected [][]string) {
	if len(result) != len(expected) {
		t.Errorf("expected %d rows, got %d", len(expected), len(result))
	}

	for i, row := range result {
		if len(row) != len(expected[i]) {
			t.Errorf("expected %d columns in row %d, got %d", len(expected[i]), i, len(row))
		}

		for j, column := range row {
			if column != expected[i][j] {
				t.Errorf("expected %s in row %d column %d, got %s", expected[i][j], i, j, column)
			}
		}
	}
}

// TestMergeFilesEmptyInputArray tests the case when the input array is empty
func TestMergeFilesEmptyInputArray(t *testing.T) {
	res, err := MergeFiles([][][]string{}, "id")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	fmt.Println(res)
}

// TestMergeFilesInputWithoutJoinColumn tests the case when the input files do not have the join column
func TestMergeFilesInputWithoutJoinColumn(t *testing.T) {
	input1 := [][]string{
		{"name", "age"},
		{"John", "20"},
		{"Jane", "21"},
		{"Jack", "22"},
	}

	input2 := [][]string{
		{"city", "country"},
		{"Tokyo", "Japan"},
		{"Warsaw", "Poland"},
		{"Washington", "USA"},
	}

	res, err := MergeFiles([][][]string{input1, input2}, "id")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	fmt.Println(res)
}
