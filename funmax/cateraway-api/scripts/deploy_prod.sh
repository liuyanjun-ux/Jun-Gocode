#!/usr/bin/env bash

make
env AWS_SHARED_CREDENTIALS_FILE=${PWD}/.aws/credentials AWS_CONFIG_FILE=${PWD}/.aws/config sls deploy --stage prod --region ap-east-1