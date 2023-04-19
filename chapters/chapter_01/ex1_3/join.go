package main

import "strings"

func SimpleJoin(args []string, sep string) string {
	res := ""
	for i, arg := range args {
		if i == 0 {
			res += arg
			continue
		}
		res += sep + arg
	}
	return res
}

func StdJoin(args []string, sep string) string {
	return strings.Join(args, sep)
}
