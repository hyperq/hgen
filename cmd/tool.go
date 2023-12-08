package cmd

import (
	"fmt"
	"os"
	"strings"
)

func replace(rs string) string {
	rs = strings.Replace(rs, "{{UpperTableName}}", UpperTableName, -1)
	rs = strings.Replace(rs, "{{TableName}}", TableName, -1)
	rs = strings.Replace(rs, "{{TableComment}}", TableComment, -1)
	rs = strings.Replace(rs, "{{TableColumns}}", GoStructs, -1)
	rs = strings.Replace(rs, "{{TsInterfaces}}", TsInterfaces, -1)
	return rs
}

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

func WriteFile(dirpath, filename, s string) (err error) {
	err = os.MkdirAll(dirpath, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath := dirpath + "/" + filename
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	_, err = f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.Sync()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
