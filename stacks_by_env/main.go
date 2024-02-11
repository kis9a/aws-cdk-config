package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/kis9a/aws-cdk-config/config"
	"github.com/kis9a/aws-cdk-config/resource"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, props *CdkStackProps, conf *config.Config) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	return awscdk.NewStack(scope, jsii.String(conf.StackName), &sprops)
}

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)
	envName := os.Getenv("CDK_ENV")
	env := config.NewEnv(envName)
	conf := env.Config()
	scope := NewCdkStack(app, &CdkStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String(conf.AccountId),
				Region:  jsii.String(conf.Region),
			},
		},
	}, conf)
	conf.BindData(scope)
	resource.NewResource(scope, conf).Make()
	app.Synth(nil)
}
