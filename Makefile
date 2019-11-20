.PHONY: build clean deploy init create.envfile s3.download.webhooksettings s3.upload.webhooksettings s3.download.cmswebhooksettings s3.upload.cmswebhooksettings update.makepayment update.webhook vendor

init: create.envfile s3.download.webhooksettings

create.envfile:
	cp -n env.example.yml env.yml

build.example:
	env GOOS=linux go build -ldflags="-s -w" -o bin/example lambda/example/main.go

get.yml:
	envi yml --id {{App Name}}

build: clean build.example

deploy.example: build
	./scripts/deploy_example.sh

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

vendor:
	go mod tidy
	go mod vendor
