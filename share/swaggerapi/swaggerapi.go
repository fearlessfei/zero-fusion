package swaggerapi

import (
	"fmt"
	"net/http"
	"path/filepath"
	"zero-fusion/share/constant"
	"zero-fusion/share/utils"

	"github.com/zeromicro/go-zero/rest"
)

// SwaggerRunOptions swagger run options
func SwaggerRunOptions(runOptions []rest.RunOption) []rest.RunOption {
	projectRoot, err := utils.GetProjectRootByGoList()
	if err != nil {
		panic(err)
	}
	currentDir := utils.RuntimeCallerSkipFilePath(2)

	swaggerUIPath := filepath.Join(projectRoot, constant.SwaggerUIPath)
	swaggerPath := filepath.Join(utils.NLevelUp(currentDir, 2), constant.SwaggerPath)
	runOptions = append(runOptions, rest.WithFileServer(constant.SwaggerStaticPath, http.Dir(swaggerUIPath)))
	runOptions = append(runOptions, rest.WithFileServer(constant.SwaggerJsonPath, http.Dir(swaggerPath)))

	fmt.Println("swaggerUIPath", swaggerUIPath)
	fmt.Println("swaggerPath", swaggerPath)

	return runOptions
}
