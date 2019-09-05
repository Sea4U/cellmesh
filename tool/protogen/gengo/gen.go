package gengo

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
	"strings"
)

func GenCode(ctx *gen.Context,temp string) error {

	gen := codegen.NewCodeGen("cmgo").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(FuncMap).
		ParseTemplate(temp, ctx)
	if strings.Contains(temp,"go") {
		gen.FormatGoCode()
	}

	if gen.Error() != nil {
		fmt.Println(string(gen.Code()))
		return gen.Error()
	}

	return gen.WriteOutputFile(ctx.OutputFileName).Error()
}
