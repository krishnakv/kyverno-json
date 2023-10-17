package template

import (
	"context"

	jpfunctions "github.com/jmespath-community/go-jmespath/pkg/functions"
	"github.com/kyverno/kyverno-json/pkg/engine/template/functions"
	kyvernofunctions "github.com/kyverno/kyverno-json/pkg/engine/template/kyverno"
)

func GetFunctions(ctx context.Context) []jpfunctions.FunctionEntry {
	var funcs []jpfunctions.FunctionEntry
	funcs = append(funcs, jpfunctions.GetDefaultFunctions()...)
	funcs = append(funcs, functions.GetFunctions()...)
	funcs = append(funcs, functions.GetFunctions()...)
	for _, function := range kyvernofunctions.GetFunctions() {
		funcs = append(funcs, function.FunctionEntry)
	}
	return funcs
}
