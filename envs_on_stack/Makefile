check-profile:
	if [ -z $(profile) ]; then exit 1; fi

diff: check-profile
	CDK_ENV=$(env) cdk diff --profile $(profile) ExampleCdkStack

deploy: check-profile
	CDK_ENV=$(env) cdk deploy --profile $(profile) ExampleCdkStack
