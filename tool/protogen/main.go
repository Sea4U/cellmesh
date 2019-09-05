package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/cellmesh/tool/protogen/gengo"
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/model"
	_ "github.com/davyxu/protoplus/msgidutil"
	"github.com/davyxu/protoplus/util"
	"os"
)

var (
	flagPackage = flag.String("package", "", "package name in source files")
	flagGoOut   = flag.String("cmgo_out", "", "cellmesh binding for golang")
	flagCsOut   = flag.String("cmcs_out","","cellmesh binding for cs")
	flagLuaOut  = flag.String("cmlua_out","","cellmesh binding for lua")
)

func main() {

	flag.Parse()

	var err error
	var ctx gen.Context
	ctx.DescriptorSet = new(model.DescriptorSet)
	ctx.DescriptorSet.PackageName = *flagPackage
	ctx.PackageName = *flagPackage
	if ctx.PackageName == "" {
		err = fmt.Errorf("package is empty")
		goto  OnError
	}
	err = util.ParseFileList(ctx.DescriptorSet)

	if err != nil {
		goto OnError
	}

	ctx.OutputFileName = *flagGoOut
	if ctx.OutputFileName != "" {
		err = gengo.GenCode(&ctx ,gengo.GoCodeTemplate)
		if err != nil {
			goto OnError
		}
	}
	ctx.OutputFileName = *flagCsOut
	if ctx.OutputFileName != "" {
		err = gengo.GenCode(&ctx,gengo.CsTemplate)
		if err != nil {
			goto  OnError
		}
	}

	ctx.OutputFileName = *flagLuaOut
	if ctx.OutputFileName != "" {
		err = gengo.GenCode(&ctx,gengo.LuaTemplate)
		if err != nil {
			goto  OnError
		}
	}


	return

OnError:
	fmt.Println(err)
	os.Exit(1)
}
