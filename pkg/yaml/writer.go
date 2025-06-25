package yaml

import (
    "io/ioutil"
    "gopkg.in/yaml.v3"
)

// WriteYAML 将数据写入YAML文件
func WriteYAML(filePath string, data interface{}) error {
    // 将数据编码为YAML格式
    output, err := yaml.Marshal(data)
    if err != nil {
        return err
    }

    // 将YAML数据写入指定文件
    return ioutil.WriteFile(filePath, output, 0644)
}