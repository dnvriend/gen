package util

import (
	"fmt"
	"github.com/dnvriend/gen/typ"
)

// returns the type of type [T], which is
// based on the typename
// eg.
// List[int], gen list -t int => IntList
// List[string], gen list -t string => StringList
// List[s3.Bucket], gen list -t s3.Bucket => S3BucketList
func TypeName(typeName string) string {
	return typ.String(typeName).
		ReplaceAll("*", "").
		ReplaceAll(".", "").
		Title().
		Str()
}

func FileName(typeName string, containerType string) string {
	return typ.String(typeName).
		Map(func(s string) string {
			return fmt.Sprintf("%v_%v.go", TypeName(typeName), containerType)
		}).
		ToLower().
		Str()
}
