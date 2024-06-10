package components

import (
	"regexp"
	"strings"
)

func PrepareScripts(scripts string) map[string]string {
	prepared_scripts := make(map[string]string)
	scripts = strings.ReplaceAll(scripts, "\r", " ")
	scripts = strings.ReplaceAll(scripts, "\n", " ")
	scripts = strings.ReplaceAll(scripts, "\t", " ")

	reg, _ := regexp.Compile(`\d+\.\s*.*?;`)
	subtaskRule, _ := regexp.Compile(`^\d+`)
	scriptRule, _ := regexp.Compile(`(select|insert|update|delete|create\s(unique\s){0,1}index).*?;`)
	matches := reg.FindAllString(scripts, -1)

	for i := 0; i < len(matches); i++ {
		subtaskNumber := subtaskRule.FindString(matches[i])
		matchesScript := scriptRule.FindString(strings.ToLower(matches[i]))
		if len(matchesScript) > 1 {
			prepared_scripts[subtaskNumber] = strings.ReplaceAll(matchesScript, ";", "")
		}
	}

	return prepared_scripts
}

func ParseBannedWords(input string) (map[string][]string, map[string][]string, error) {
	bannedWords := make(map[string][]string)
	admissionWords := make(map[string][]string)

	reg, _ := regexp.Compile(`\d+\.\s*.*?;`)
	bannedWordRule, _ := regexp.Compile(`Запрет:(.*?)!`)
	admissionWordRule, _ := regexp.Compile(`Допуск:(.*?)!`)
	subtaskRule, _ := regexp.Compile(`^\d+`)

	bannedWordList := reg.FindAllString(input, -1)

	for i := 0; i < len(bannedWordList); i++ {
		subtaskNumber := subtaskRule.FindString(bannedWordList[i])
		matchesBanned := bannedWordRule.FindStringSubmatch(bannedWordList[i])
		matchesAdmission := admissionWordRule.FindStringSubmatch(bannedWordList[i])
		if len(matchesBanned) > 1 {
			bannedWords[subtaskNumber] = strings.Fields(matchesBanned[1])
		}
		if len(matchesAdmission) > 1 {
			admissionWords[subtaskNumber] = strings.Fields(matchesAdmission[1])
		}
	}
	return bannedWords, admissionWords, nil
}

func RemoveEmpty(slice *[]string) {
	i := 0
	p := *slice
	for _, entry := range p {
		if strings.Trim(entry, " ") != "" {
			p[i] = entry
			i++
		}
	}
	*slice = p[0:i]
}

func ParseIndexName(inputedScript string) string {
	reg, err := regexp.Compile(`create\s+(?:unique\s+)?index\s+(\w+)\s+on`)
	if err != nil {
		return ""
	}
	matches := strings.Split(reg.FindString(inputedScript), " ")
	length := len(matches)
	if length >= 4 {
		return matches[length-1]
	}
	return ""
}
