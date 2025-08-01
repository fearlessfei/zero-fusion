package main

import (
	"zero-fusion/app/demo/api/scripts/common"
	demologic "zero-fusion/app/demo/api/scripts/internal/logic/demo"
)

func main() {
	common.RunScript(&demologic.DemoLogic{})
}
