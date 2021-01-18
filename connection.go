package tools

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	_nil = iota
	SINT16
	SINT32
	SINT64
	UINT16
	UINT32
	UINT64
	FLOAT32
)

func GetField(v interface{}, connectionString, param string) (err error) {
	switch param {
	case "baseUrl", "connection":
		fields := strings.Fields(connectionString)
		for _, field := range fields {
			// check if connection string is valid
			if regexp.MustCompile(`^https?://.*$`).MatchString(field) || regexp.MustCompile(`^[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}:[\d]{1,5}$`).MatchString(field) {
				if x, ok := v.(*string); ok {
					*x = field
				}
				return
			}
		}
	case "format":
		fields := strings.Fields(connectionString)
		var i int
		for _, field := range fields {
			switch field {
			case "sint16":
				i = SINT16
			case "sint32":
				i = SINT32
			case "sint64":
				i = SINT64
			case "uint16":
				i = UINT16
			case "uint32":
				i = UINT32
			case "uint64":
				i = UINT64
			case "float32":
				i = FLOAT32
			default:
				continue
			}

			if x, ok := v.(*int); ok {
				*x = i
			}
			return
		}

	default:
		fields := strings.Fields(connectionString)
		for _, field := range fields {
			parts := strings.Split(field, ":")
			if parts[0] != param || len(parts) != 2 {
				continue
			}

			value := parts[1]

			switch x := v.(type) {
			case *string:
				*x = value
			case *int:
				*x, err = strconv.Atoi(value)
			case *uint16:
				var i int
				i, err = strconv.Atoi(value)
				*x = uint16(i)
			case *uint8:
				var i int
				i, err = strconv.Atoi(value)
				*x = uint8(i)
			case *time.Duration:
				var i int
				if i, err = strconv.Atoi(value); err == nil {
					*x = time.Duration(i) * time.Millisecond
				}
			}
			return
		}
	}

	return
}

func Quantity(format int) int {
	switch format {
	case SINT16, UINT16:
		return 1
	case SINT32, UINT32, FLOAT32:
		return 2
	case SINT64, UINT64:
		return 4
	}
	return 0
}
