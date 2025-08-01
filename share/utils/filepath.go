package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// GetProjectRootByGoList 获取当前项目根目录
func GetProjectRootByGoList() (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}

// RuntimeCallerSkipFile 返回调用者的文件
func RuntimeCallerSkipFile(skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	return file
}

// RuntimeCallerSkipFilePath 返回调用者的文件路径
func RuntimeCallerSkipFilePath(skip int) string {
	return filepath.Dir(RuntimeCallerSkipFile(skip))
}

// NLevelUp 返回 n 级上级目录
func NLevelUp(path string, levels int) string {
	for i := 0; i < levels; i++ {
		path = filepath.Dir(path)
	}

	return path
}

// GetExecutablePath 获取可执行文件路径
func GetExecutablePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("error getting executable path: %w", err)
	}

	return exePath, nil
}

// GetPathFileNameNoExt 获取文件名，不包含扩展名
func GetPathFileNameNoExt(filePath string) string {
	base := filepath.Base(filePath)
	return strings.TrimSuffix(base, filepath.Ext(base))
}
