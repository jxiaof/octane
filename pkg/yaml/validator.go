package yaml

import (
    "errors"
    "gopkg.in/yaml.v3"
)

// ValidateYAML validates the given YAML data against the expected structure.
func ValidateYAML(data []byte, out interface{}) error {
    err := yaml.Unmarshal(data, out)
    if err != nil {
        return errors.New("invalid YAML format: " + err.Error())
    }
    return nil
}