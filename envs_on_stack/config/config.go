package config

import (
	"github.com/aws/constructs-go/constructs/v10"
)

type Config struct {
	StackName string
	AccountId string
	Region    string
	Envs      Envs
	Data      Data
}

type EnvConfig struct {
	Data Data
	Env
}

type Data struct{}

type EnvData struct{}

func NewConfig() *Config {
	return &Config{
		Region:    "us-east-1",
		AccountId: "111111111111",
		StackName: "ExampleCdkStack",
		Envs:      []Env{EnvDevelop{}, EnvProduction{}},
	}
}

func NewEnvConfig(env Env) *EnvConfig {
	return &EnvConfig{
		Env: env,
	}
}

func (c *Config) BindData(scope constructs.Construct) {}

func (e *EnvConfig) BindEnvData(scope constructs.Construct) {}
