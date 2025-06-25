package executor

import (
    "os/exec"
    "path/filepath"
)

// PythonExecutor 结构体用于执行Python脚本
type PythonExecutor struct {
    ScriptPath string // Python脚本的路径
}

// NewPythonExecutor 创建一个新的Python执行器
func NewPythonExecutor(scriptPath string) *PythonExecutor {
    return &PythonExecutor{
        ScriptPath: scriptPath,
    }
}

// Execute 执行Python脚本并返回输出
func (pe *PythonExecutor) Execute(args ...string) (string, error) {
    cmd := exec.Command("python3", append([]string{pe.ScriptPath}, args...)...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}

// GetScriptPath 获取Python脚本的绝对路径
func (pe *PythonExecutor) GetScriptPath() (string, error) {
    absPath, err := filepath.Abs(pe.ScriptPath)
    if err != nil {
        return "", err
    }
    return absPath, nil
}