package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

func main() {
	wasim1, err := os.ReadFile("module1.wasm")
	if err != nil {
		log.Fatalln(err)
	}

	wasim2, err := os.ReadFile("module2.wasm")
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	config := wazero.NewCompileConfig()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntimeWithConfig(
		wazero.NewRuntimeConfig().WithWasmCore2(),
	)

	mcfg := wazero.NewModuleConfig().WithEnv(
		"FOO", "BAR2",
	).WithStderr(
		os.Stderr,
	).WithStdout(
		os.Stdout,
	).WithSysNanotime()

	compiled1, err := r.CompileModule(context.Background(), wasim1, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer compiled1.Close(ctx)

	// Compile WebAssembly that requires its own "env" module.
	compiled2, err := r.CompileModule(context.Background(), wasim2, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer compiled2.Close(ctx)

	debugmodule("m1", compiled1)
	debugmodule("m2", compiled2)

	ns1 := r.NewNamespace(ctx)

	wasienv, err := wasi_snapshot_preview1.NewBuilder(r).Instantiate(ctx, ns1)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "wasi module"))
	}
	defer wasienv.Close(ctx)

	m1, err := ns1.InstantiateModule(
		ctx,
		compiled1,
		mcfg.WithName("env"),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "m1 module"))
	}
	defer m1.Close(ctx)

	m2, err := ns1.InstantiateModule(
		ctx,
		compiled2,
		mcfg.WithName("m2"),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "m2 module"))
	}
	defer m2.Close(ctx)

}

func debugmodule(name string, m wazero.CompiledModule) {
	log.Println("module debug", name, m.Name())
	for _, imp := range m.ExportedFunctions() {
		paramtypestr := typeliststr(imp.ParamTypes()...)
		resulttypestr := typeliststr(imp.ResultTypes()...)
		log.Println("exported", imp.Name(), "(", paramtypestr, ")", resulttypestr)
	}

	for _, imp := range m.ImportedFunctions() {
		paramtypestr := typeliststr(imp.ParamTypes()...)
		resulttypestr := typeliststr(imp.ResultTypes()...)
		log.Println("imported", imp.Name(), "(", paramtypestr, ")", resulttypestr)
	}
}

func typeliststr(types ...api.ValueType) string {
	typesstr := []string(nil)
	for _, t := range types {
		typesstr = append(typesstr, api.ValueTypeName(t))
	}

	return strings.Join(typesstr, ", ")
}
