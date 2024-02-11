package resource

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kis9a/aws-cdk-config/config"
)

type Resource struct {
	conf  *config.Config
	scope constructs.Construct
}

func NewResource(scope constructs.Construct, conf *config.Config) *Resource {
	return &Resource{
		scope: scope,
		conf:  conf,
	}
}

func (r *Resource) Make() {
	for _, env := range r.conf.Envs {
		envConfig := config.NewEnvConfig(env)
		envConfig.BindEnvData(r.scope)
		fmt.Printf("%s\n", envConfig.Identifier().ToString())
	}
}
