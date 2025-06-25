package yaml

import (
    "gopkg.in/yaml.v3"
    "bytes"
)

// Formatter provides methods to format YAML data.
type Formatter struct{}

// NewFormatter creates a new instance of Formatter.
func NewFormatter() *Formatter {
    return &Formatter{}
}

// Format takes an input data structure and returns a formatted YAML string.
func (f *Formatter) Format(data interface{}) (string, error) {
    var buf bytes.Buffer
    encoder := yaml.NewEncoder(&buf)
    encoder.SetIndent(2) // Set indentation for better readability
    if err := encoder.Encode(data); err != nil {
        return "", err
    }
    return buf.String(), nil
}