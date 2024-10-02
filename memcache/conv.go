package memcache

import (
	"fmt"
	"strconv"
)

// GetString convert interface to string.
func GetString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	default:
		if v != nil {
			return fmt.Sprint(result)
		}
	}

	return ""
}

// GetInt convert interface to int.
func GetInt(v interface{}) int {
	switch result := v.(type) {
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	default:
		if d := GetString(v); d != "" {
			val, _ := strconv.Atoi(d)
			return val
		}
	}

	return 0
}

// GetInt64 convert interface to int64.
func GetInt64(v interface{}) int64 {
	switch result := v.(type) {
	case int64:
		return result
	case int32:
		return int64(result)
	case int:
		return int64(result)
	default:
		if d := GetString(v); v != "" {
			val, _ := strconv.ParseInt(d, 10, 64)
			return val
		}
	}

	return 0
}

// GetFloat64 convert interface to float64.
func GetFloat64(v interface{}) float64 {
	switch result := v.(type) {
	case float64:
		return result
	default:
		if d := GetString(v); d != "" {
			val, _ := strconv.ParseFloat(d, 64)
			return val
		}
	}

	return 0
}

// GetBool convert interface to bool.
func GetBool(v interface{}) bool {
	switch result := v.(type) {
	case bool:
		return result
	default:
		if d := GetString(v); d != "" {
			val, _ := strconv.ParseBool(d)
			return val
		}
	}

	return false
}
