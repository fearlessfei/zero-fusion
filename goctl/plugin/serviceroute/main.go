package main

import (
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"
	goctlplugin "github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func main() {
	plugin, err := goctlplugin.NewPlugin()
	if err != nil {
		panic(err)
	}

	dir := plugin.Dir
	rootPkg, err := golang.GetParentPackage(dir)
	cfg, err := config.NewConfig("")
	if err != nil {
		panic(err)
	}

	err = genRoutes(dir, rootPkg, cfg, plugin.Api)
	if err != nil {
		panic(err)
	}

	// Delete the system-generated routes.go
	routesFilePath := filepath.Join(dir, handlerDir, routesFilename+".go")
	if fileExists(routesFilePath) {
		err = os.Remove(routesFilePath)
		if err != nil {
			panic(err)
		}
	}

	// Delete the system-generated service startup file
	// that was created when multiple route files were generated,
	// since multiple services now require their own startup files.
	serviceStartupFilePath := filepath.Join(dir, plugin.Api.Service.Name+".go")
	if fileExists(serviceStartupFilePath) {
		err = os.Remove(serviceStartupFilePath)
		if err != nil {
			panic(err)
		}
	}
}
