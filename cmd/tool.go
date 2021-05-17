package cmd

import "strings"

func lintString(s string) string {
	if s == "" {
		return s
	}
	s = strings.Replace(s, "id", "ID", -1)
	sslice := strings.Split(s, "")
	length := len(sslice)
	var nslice []string
	for i := range sslice {
		if sslice[i] == "_" && i+1 < length {
			sslice[i+1] = strings.ToUpper(sslice[i+1])
		} else {
			if i == 0 {
				sslice[i] = strings.ToUpper(sslice[i])
			}
			nslice = append(nslice, sslice[i])
		}
	}
	return strings.Join(nslice, "")
}
