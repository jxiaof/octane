# Octane Performance Analyzer

## 项目概述
`octane` 是一个主机综合性能测试工具，旨在评估交付给用户的主机性能。该工具采用 **Go + Python 混合架构**，支持单项测试、中心化数据上报和专业化场景测试。

## 特性
- **多项性能测试**: 支持 CPU、内存、存储、GPU 和网络等多项性能测试。
- **快速测试模式**: 提供快速性能测试命令，便于用户快速获取系统性能概览。
- **辛烷值评级**: 通过辛烷值算法对系统性能进行评级，帮助用户了解系统的性能等级。
- **专业化测试**: 针对游戏、AI、服务器等场景提供专业化测试选项。
- **数据上报**: 支持将测试结果上传至中心化服务器，便于数据分析和比较。

## 技术栈
- **Go**: 主要框架和命令行工具
- **Python**: 测试脚本和性能评估
- **SQLite**: 数据存储
- **YAML**: 配置文件格式

## 安装
1. 克隆项目：
   ```
   git clone https://octane.git
   cd octane
   ```
2. 安装 Go 依赖：
   ```
   go mod tidy
   ```
3. 安装 Python 依赖：
   ```
   pip install -r scripts/python/requirements.txt
   ```

## 使用
- 查看系统信息:
  ```
  octane info
  ```
- 运行完整测试套件:
  ```
  octane test
  ```
- 快速性能测试:
  ```
  octane boost
  ```
- 查看辛烷值评级:
  ```
  octane rating
  ```

## 贡献
欢迎任何形式的贡献！请查看 [CONTRIBUTING.md](docs/CONTRIBUTING.md) 以获取更多信息。

## 许可证
该项目采用 MIT 许可证，详细信息请查看 [LICENSE](LICENSE)。