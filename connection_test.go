package tools

import (
	"testing"
	"time"
)

func TestQuantity(t *testing.T) {
	testPattern := map[int]int{
		FLOAT32: 2,
		SINT16:  1,
		SINT32:  2,
		SINT64:  4,
		UINT16:  1,
		UINT32:  2,
		UINT64:  4}

	for k, v := range testPattern {
		if Quantity(k) != v {
			t.Errorf("key: %v value: %v", k, v)
		}
	}
}

func TestConnection(t *testing.T) {
	type value map[string]interface{}
	testPattern := map[string]value{
		"tcp 192.168.65.197:502 deviceid:1 maxretries:3 timeout:1000": {
			"connection": "192.168.65.197:502",
			"deviceid":   1,
			"maxretries": 3,
			"timeout":    time.Second},
		"http http://raspberryz.fritz.box:8080 maxretries:3 timeout:1000": {
			"connection": "http://raspberryz.fritz.box:8080",
			"maxretries": 3,
			"timeout":    time.Second},
		"http://fritz.box ain:116570149698 username:smarthome password:7Wl6UW5TsOr5Ba6uMbOO timeout:2000 maxretries:3": {
			"baseUrl":    "http://fritz.box",
			"ain":        "116570149698",
			"username":   "smarthome",
			"password":   "7Wl6UW5TsOr5Ba6uMbOO",
			"maxretries": 3,
			"timeout":    2 * time.Second},
		"http://wallbox:4000/currentdata timeout:1000 maxretries:3 cache:1000": {
			"connection": "http://wallbox:4000/currentdata",
			"maxretries": 3,
			"cache":      time.Second,
			"timeout":    time.Second},
		"https://wallbox": {
			"connection": "https://wallbox",
		},
		"http://www.orf.at": {
			"connection": "http://www.orf.at",
		},
		"http://www.orf.at:65000/a/b/c": {
			"connection": "http://www.orf.at:65000/a/b/c",
		},
		"ftp://www.orf.at": {
			"connection": "",
		},
		"1.1.1.1:1": {
			"connection": "1.1.1.1:1",
		},
		"999.999.999.999:0": {
			"connection": "999.999.999.999:0",
		},
		"1.1.1.1:1 p1:1 p2:abc p3:1": {
			"connection": "1.1.1.1:1",
			"p1":         1,
			"p2":         "abc",
			"p3":         time.Millisecond,
		},
		"address:4116 uint32 sf:-2": {
			"address": 4116,
			"sf":      -2,
			"format":  UINT32,
		},
		"address:4116 uint33 sf:-2": {
			"format": 0,
		},
	}

	for connString, m := range testPattern {
		for key, value := range m {
			switch value.(type) {
			case int:
				var v int
				if err := GetField(&v, connString, key); err != nil {
					t.Errorf("error: %v", err)
				}
				if !IsEqual(v, value) {
					t.Errorf("connection string: %v\nkey: %v is: %v should be: %v", connString, key, v, value)
				}
			case string:
				var v string
				if err := GetField(&v, connString, key); err != nil {
					t.Errorf("error: %v", err)
				}
				if !IsEqual(v, value) {
					t.Errorf("connection string: %v\nkey: %v is: %v should be: %v", connString, key, v, value)
				}
			case uint8:
				var v uint8
				if err := GetField(&v, connString, key); err != nil {
					t.Errorf("error: %v", err)
				}
				if !IsEqual(v, value) {
					t.Errorf("connection string: %v\nkey: %v is: %v should be: %v", connString, key, v, value)
				}
			case time.Duration:
				var v time.Duration
				if err := GetField(&v, connString, key); err != nil {
					t.Errorf("error: %v", err)
				}
				if !IsEqual(v, value) {
					t.Errorf("connection string: %v\nkey: %v is: %v should be: %v", connString, key, v, value)
				}
			}
		}
	}
}
