check-env:
	if [ -z $(env) ]; then exit 1; fi

check-profile:
	if [ -z $(profile) ]; then exit 1; fi

diff: check-env check-profile
	CDK_ENV=$(env) cdk diff --profile $(profile) ExampleCdkStack-$(env)

deploy: check-env check-profile
	CDK_ENV=$(env) cdk deploy --profile $(profile) ExampleCdkStack-$(env)
