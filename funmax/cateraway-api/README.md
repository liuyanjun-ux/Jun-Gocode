# Cateraway Backend APIs
Cateraway Apis.

## Development Enviroment
 * serverless [1.54.0]
 * go [1.x] 
 * gin [1.3.0]
 * aws-sam-cli [0.6.0]

## Local Development Prerequisite
 * Docker

## Project setup
```
- setup go environment in your local machine
  Ref: https://golang.org/doc/install

- install aws-sam-cli globally in your local machine
  Ref: https://docs.aws.amazon.com/en_us/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html

- cd to your project directory, rename 
    template_sample.yml to  template.yml
    env_sample.yml to env.yml

  update the local environment variables in template.yml
  update the deployment environment variables in env.yml

Local development, we use aws-sam-cli
- run 'make'
- run 'sam local start-api --skip-pull-image'

Deploy to AWS (Development, ask admin for aws credentials)
- run './scripts/deploy-dev.sh'

Deploy to AWS (Production, ask admin for aws credentials)
- run './scripts/deploy-prod.sh'

```
