package methods

import (
	"fmt"
	"regexp"
	"strings"
)

func PrepareDBScript(script, dbName string) (string, error) {
	createDataBase := fmt.Sprintf("CREATE DATABASE %s", dbName)
	if !strings.Contains(script, createDataBase) {
		var sb strings.Builder
		_, err := sb.WriteString(createDataBase)
		if err != nil {
			return "", err
		}
		_, err = sb.WriteString(script)
		if err != nil {
			return "", err
		}
		return sb.String(), nil
	}
	pattern := regexp.MustCompile("(?i)(CREATE\\s+DATABASE.*?;|\\c.*?;|USE\\s+\\w+;)")
	return pattern.ReplaceAllString(script, createDataBase), nil
}
