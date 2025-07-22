package tool

import (
	"testing"

	"github.com/Js41313/Futuer-2/pkg/constant"
)

func TestExtractVersionNumber(t *testing.T) {
	versionNumber := ExtractVersionNumber(constant.Version)
	t.Log(versionNumber)
}
