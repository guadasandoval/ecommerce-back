package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// Data struct
type Data struct {
	data map[string]interface{}
}

// Config struct
type Config struct {
	version string
	name    string
	data    map[string]interface{}
}

// NewConfig - Create a config
func NewConfig() *Config {
	var config Config
	return &config
}

// ReadConfigFile - Read the config file
func (c *Config) ReadConfigFile(configFile string) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("ReadConfigFile: %w", err)
	}

	dataJSON := make(map[string]interface{})
	err = json.Unmarshal(data, &dataJSON)
	if err != nil {
		return fmt.Errorf("ReadConfigFile: %w", err)
	}

	var ok bool
	var tmp interface{}

	tmp, ok = dataJSON["version"]
	if !ok {
		return fmt.Errorf("'version' is not defined")
	}
	c.version = tmp.(string)

	tmp, ok = dataJSON["name"]
	if !ok {
		return fmt.Errorf("'name' is not defined")
	}
	c.name = tmp.(string)

	tmp, ok = dataJSON["config"]
	if !ok {
		return fmt.Errorf("'config' is not defined")
	}
	c.data = tmp.(map[string]interface{})

	return nil
}

// Version - Get the version
func (c *Config) Version() string {
	return c.version
}

// Name - Get the name
func (c *Config) Name() string {
	return c.name
}

// Data - Get the data
func (c *Config) Data() *Data {
	return &Data{
		data: c.data,
	}
}

// GetParam - Get a param
func (c *Data) GetParam(name string) interface{} {
	return c.data[name]
}

// GetParamAsBool - Get a param as bool
func (c *Data) GetParamAsBool(name string) bool {
	value := c.data[name]
	switch value.(type) {
	case bool:
		return value.(bool)
	}

	panic(fmt.Sprintf("GetParamAsBool Error: value of type %T", value))
}

// GetParamAsInt - Get a param as int
func (c *Data) GetParamAsInt(name string) int {
	value := c.data[name]
	switch value.(type) {
	case bool:
		valueBool := value.(bool)
		if valueBool {
			return 1
		}
		return 0
	case float64:
		return int(value.(float64))
	case string:
		retInt, err := strconv.Atoi(value.(string))
		if err != nil {
			panic(err.Error())
		}
		return retInt
	}

	panic(fmt.Sprintf("GetParamAsInt Error: value of type %T", value))
}

// GetParamAsFloat - Get a param as float
func (c *Data) GetParamAsFloat(name string) float64 {
	value := c.data[name]
	switch value.(type) {
	case bool:
		valueBool := value.(bool)
		if valueBool {
			return 1
		}
		return 0
	case float64:
		return value.(float64)
	case string:
		valueFloat, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			panic(err.Error())
		}
		return valueFloat
	}

	panic(fmt.Sprintf("GetParamAsFloat Error: value of type %T", value))
}

// GetParamAsString - Get a param as string
func (c *Data) GetParamAsString(name string) string {
	value := c.data[name]
	switch value.(type) {
	case bool:
		valueBool := value.(bool)
		if valueBool {
			return "true"
		}
		return "false"
	case float64:
		valueFloat := value.(float64)
		return strconv.FormatFloat(valueFloat, 'f', -1, 64)
	case string:
		return c.data[name].(string)
	}

	panic(fmt.Sprintf("GetParamAsString Error: value of type %T", value))
}

// GetParamAsArray - Get a param as array
func (c *Data) GetParamAsArray(name string) []interface{} {
	value := c.data[name]
	switch value.(type) {
	case []interface{}:
		return value.([]interface{})
	}

	panic(fmt.Sprintf("GetParamAsString Error: value of type %T", value))
}

// GetParamAsData - Get a param as data
func (c *Data) GetParamAsData(name string) *Data {
	value := c.data[name].(map[string]interface{})

	return &Data{
		data: value,
	}
}
