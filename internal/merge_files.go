package internal

import "fmt"


func MergeFiles(files [][][]string, joinColumn string) ([][]string, error) {
	
	if len(files) == 0 {
		return nil, fmt.Errorf("attempt to merge empty array of files")
	}

	columns := []string{joinColumn}

	rowsByJoinKey := make(map[string][]string)
	var joinKeys []string

	for _, file := range files {
		if len(file) == 0 {
			return nil, fmt.Errorf("attempt to merge empty file")
		}

		fileColumns := file[0]		
		joinColumnIndex, err := IndexOf(fileColumns, joinColumn)
		if err != nil {
			return nil, fmt.Errorf("join column not found in the file. error: %v", err) 
		}
		
		fileDataColumns := ExcludeIndex(fileColumns, joinColumnIndex)
		columns = append(columns, fileDataColumns...)
		
		for _, fileRow := range file[1:] {
			joinKey := fileRow[joinColumnIndex]
			dataRow := ExcludeIndex(fileRow, joinColumnIndex)
			_, ok := rowsByJoinKey[joinKey]
			if !ok {
				joinKeys = append(joinKeys, joinKey)
				initRow := make([]string, 0, len(fileRow))
				initRow = append(initRow, joinKey)
				initRow = append(initRow, dataRow...)
				rowsByJoinKey[joinKey] = initRow
			} else {
				rowsByJoinKey[joinKey] = append(rowsByJoinKey[joinKey], dataRow...)
			}
		}
	}

	dataRows := make([][]string, 0, len(joinKeys))
	for _, value := range joinKeys {
		dataRows = append(dataRows, rowsByJoinKey[value])
	}
		
	return append([][]string{columns}, dataRows...), nil
}

