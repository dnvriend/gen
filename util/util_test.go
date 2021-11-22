package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"IntList", "IntList"},
		{"*IntList", "IntList"},
		{"int", "Int"},
		{"*int", "Int"},
		{"string", "String"},
		{"*string", "String"},
		{"s3.Bucket", "S3Bucket"},
		{"*s3.Bucket", "S3Bucket"},
		{"iam.Policy", "IamPolicy"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, TypeName(test.input))
	}
}

func TestFileName(t *testing.T) {
	tests := []struct {
		input         string
		expected      string
		containerType string
	}{
		{"int", "int_list.go", "list"},
		{"*int", "int_list.go", "list"},
		{"string", "string_list.go", "list"},
		{"*string", "string_list.go", "list"},
		{"s3.Bucket", "s3bucket_list.go", "list"},
		{"*s3.Bucket", "s3bucket_list.go", "list"},
		{"iam.Policy", "iampolicy_list.go", "list"},
		{"*iam.Policy", "iampolicy_list.go", "list"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, FileName(test.input, test.containerType))
	}
}
