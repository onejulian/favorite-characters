#!/bin/bash

tag=$1
if [[ -z "$tag" ]]; then
    echo "Please provide a tag, either 'stage' or 'prod'"
    exit 1
elif [[ "$tag" != "stage" && "$tag" != "prod" ]]; then
    echo "Tag must be either 'stage' or 'prod'"
    exit 1
fi

GOOS=linux GOARCH=amd64 go build -o main &&
zip main.zip main &&
rm main &&

if [[ "$tag" == "prod" ]]; then
    tag="-prod"
fi

if [[ "$tag" == "stage" ]]; then
    tag="-stage"
fi

aws lambda update-function-code --function-name favorite-characters$tag --zip-file fileb://main.zip > /dev/null 2>&1 &&

rm main.zip
