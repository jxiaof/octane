package executor

import (
    "os/exec"
    "path/filepath"
)

// Executor 结构体用于执行Python脚本
type Executor struct {
    ScriptPath string // Python脚本的路径
}

// NewExecutor 创建一个新的Executor实例
func NewExecutor(scriptPath string) *Executor {
    return &Executor{
        ScriptPath: scriptPath,
    }
}

// Run 执行Python脚本并返回输出
func (e *Executor) Run(args ...string) (string, error) {
    cmd := exec.Command("python3", append([]string{e.ScriptPath}, args...)...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}

// GetScriptPath 获取脚本的绝对路径
func (e *Executor) GetScriptPath() (string, error) {
    absPath, err := filepath.Abs(e.ScriptPath)
    if err != nil {
        return "", err
    }
    return absPath, nil
}