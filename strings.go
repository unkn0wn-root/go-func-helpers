package strings

import (
    "strings"
)

func ContainsString(list []string, target string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}
	return false
}

func HasPrefix(list []string, prefix string) bool {
	for _, str := range list {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

// Return a new *[]string with empty lines removed from *[]string.
// Original []string not modified.
func StrArrayPtrRemoveEmpty(stringList *[]string) *[]string {
	var sl []string
	var l int = len(*stringList)
	for i := 0; i < l; i++ {
		if (*stringList)[i] != "" {
			sl = append(sl, (*stringList)[i])
		}
	}

    return &sl
}

func BoolString(b bool) string {
	if b {
		return "true"
	}

    return "false"
}

func BoolStatus(b bool) string {
	if b {
		return "ok"
	}

    return "fail"
}

func BoolYesNo(b bool) string {
	if b {
		return "yes"
	}

    return "no"
}
