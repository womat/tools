package tools

import "encoding/json"

func IsEqual(a interface{}, b interface{}) bool {
	expect, _ := json.Marshal(a)
	got, _ := json.Marshal(b)
	return string(expect) == string(got)
}

func In(s interface{}, pattern ...interface{}) bool {
	for _, p := range pattern {
		if s == p {
			return true
		}
	}
	return false
}
