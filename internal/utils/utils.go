package utils

import "os"

func ReadSqlFile(path string) (string, error) {

	sql_query, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(sql_query), nil
}
