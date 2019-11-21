#!/bin/bash

set -e
. scripts/utils.sh

ENVIRONMENT=$(get_environment)

echo Making Directories...
cd bin
mkdir "example_dist"
mkdir "example_dist/bin"
echo Complete

echo Moving files to own folder...
mv example example_dist/bin
echo Complete

echo Zipping up distribution files...
cd example_dist
zip -r ../example_dist.zip bin
echo Complete

echo Updating Lambda Functions...

aws lambda update-function-code \
      --function-name kms_example_lambda_${ENVIRONMENT} \
      --zip-file fileb://../example_dist.zip
echo Complete