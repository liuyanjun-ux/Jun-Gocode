#!/usr/bin/env bash

env AWS_SHARED_CREDENTIALS_FILE=${PWD}/.aws/credentials AWS_CONFIG_FILE=${PWD}/.aws/config sls remove --stage dev --region ap-east-1