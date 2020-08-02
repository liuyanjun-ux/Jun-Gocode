#!/usr/bin/env bash

make
env AWS_SHARED_CREDENTIALS_FILE=.aws/credentials AWS_CONFIG_FILE=${PWD}/.aws/config sls deploy --stage dev --region ap-east-1${PWD}/

