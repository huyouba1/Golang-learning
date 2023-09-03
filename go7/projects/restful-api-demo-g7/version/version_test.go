package version_test

import (
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/version"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}
