package config

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
)

const (
	stackNamePrefix = "ExampleCdkStack"
)

type Config struct {
	StackName string
	AccountId string
	Region    string
	Data      Data
}

type Data struct{}

func (e EnvDevelop) Config() *Config {
	return &Config{
		StackName: fmt.Sprintf("%s-%s", stackNamePrefix, e.Identifier().ToString()),
		AccountId: "111111111111",
		Region:    "us-east-1",
	}
}

func (e EnvProduction) Config() *Config {
	return &Config{
		StackName: fmt.Sprintf("%s-%s", stackNamePrefix, e.Identifier().ToString()),
		AccountId: "222222222222",
		Region:    "ap-northeast-1",
	}
}

func (c *Config) BindData(scope constructs.Construct) {}
