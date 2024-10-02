package memcache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name     string
		payload  interface{}
		expected string
	}{
		{
			name:     "string to string",
			payload:  "test1",
			expected: "test1",
		},
		{
			name:     "byte array to string",
			payload:  []byte("test2"),
			expected: "test2",
		},
		{
			name:     "int to string",
			payload:  "1",
			expected: "1",
		},
		{
			name:     "int64 to string",
			payload:  "1",
			expected: "1",
		},
		{
			name:     "float64 to string",
			payload:  "1.1",
			expected: "1.1",
		},
		{
			name:     "nil to string",
			payload:  nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := GetString(tt.payload)
			assert.Equal(t, tt.expected, actual)

			t.Parallel()
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name     string
		payload  interface{}
		expected int
	}{
		{
			name:     "int to int",
			payload:  1,
			expected: 1,
		},
		{
			name:     "int32 to int",
			payload:  int32(32),
			expected: 32,
		},
		{
			name:     "int64 to int",
			payload:  int64(64),
			expected: 64,
		},
		{
			name:     "num string to int",
			payload:  "128",
			expected: 128,
		},
		{
			name:     "nil to int",
			payload:  nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := GetInt(tt.payload)
			assert.Equal(t, tt.expected, actual)

			t.Parallel()
		})
	}
}

func TestGetInt64(t *testing.T) {
	tests := []struct {
		name     string
		payload  interface{}
		expected int64
	}{
		{
			name:     "int to int64",
			payload:  1,
			expected: 1,
		},
		{
			name:     "int32 to int64",
			payload:  int32(32),
			expected: 32,
		},
		{
			name:     "int64 to int64",
			payload:  int64(64),
			expected: 64,
		},
		{
			name:     "num string to int64",
			payload:  "128",
			expected: 128,
		},
		{
			name:     "nil to int64",
			payload:  nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := GetInt64(tt.payload)
			assert.Equal(t, tt.expected, actual)

			t.Parallel()
		})
	}
}

func TestGetFloat64(t *testing.T) {
	tests := []struct {
		name     string
		payload  interface{}
		expected float64
	}{
		{
			name:     "float64 to float64",
			payload:  float64(1.11),
			expected: 1.11,
		},
		{
			name:     "float32 to float64",
			payload:  float32(1.11),
			expected: 1.11,
		},
		{
			name:     "int64 to float64",
			payload:  int64(1),
			expected: 1,
		},
		{
			name:     "int to float64",
			payload:  int(1),
			expected: 1,
		},
		{
			name:     "num string to float64",
			payload:  "1.128",
			expected: 1.128,
		},
		{
			name:     "nil to float64",
			payload:  nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := GetFloat64(tt.payload)
			assert.Equal(t, tt.expected, actual)

			t.Parallel()
		})
	}
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		name     string
		payload  interface{}
		expected bool
	}{
		{
			name:     "bool to bool",
			payload:  true,
			expected: true,
		},
		{
			name:     "string to bool",
			payload:  "true",
			expected: true,
		}, {
			name:     "nil to bool",
			payload:  nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := GetBool(tt.payload)
			assert.Equal(t, tt.expected, actual)

			t.Parallel()
		})
	}
}
