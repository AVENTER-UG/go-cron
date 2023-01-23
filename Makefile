#vars
IMAGENAME=go-cron
REPO=localhost:5000
TAG=`git describe --tags --abbrev=0`
BRANCH=`git rev-parse --abbrev-ref HEAD`
BUILDDATE=`date -u +%Y-%m-%dT%H:%M:%SZ`
IMAGEFULLNAME=${REPO}/${IMAGENAME}

sboom:
	syft dir:. > sbom.txt
	syft dir:. -o json > sbom.json

seccheck:
	gosec --exclude G104 --exclude-dir ./vendor ./...

go-fmt:
	@gofmt -w .
	@golangci-lint run --fix
	@gocritic check -disable 'whynolint'

build:
	@echo ">>>> Build Docker"
	@docker build --build-arg TAG=${TAG} --build-arg BUILDDATE=${BUILDDATE} -t ${IMAGEFULLNAME}:${TAG} .

build-bin:
	@echo ">>>> Build binary"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.BuildVersion=${BUILDDATE} -X main.GitVersion=${TAG} -extldflags \"-static\"" .

check: sboom seccheck go-fmt
all: check build
