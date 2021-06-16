TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=paulfreaknbaker.com
NAMESPACE=providers
NAME=dockerutils
BINARY_PREFIX=terraform-provider-${NAME}
VERSION=0.1
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

default: install

build:
	@[ -n "${GOOS}" ] || (echo "Set GOOS"; exit 1)
	@[ -n "${GOARCH}" ] || (echo "Set GOARCH"; exit 1)
	@mkdir -p ./bin
	go build -o ./bin/${BINARY_PREFIX}_${VERSION}_${GOOS}_${GOARCH}

install:
	$(MAKE) build
	mkdir -p ${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${GOOS}_${GOARCH}
	mv ./bin/${BINARY_PREFIX}_${VERSION}_${GOOS}_${GOARCH} ${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${GOOS}_${GOARCH}

release:
	$(MAKE) build GOOS=darwin GOARCH=amd64
	$(MAKE) build GOOS=freebsd GOARCH=386
	$(MAKE) build GOOS=freebsd GOARCH=amd64
	$(MAKE) build GOOS=freebsd GOARCH=arm
	$(MAKE) build GOOS=linux GOARCH=386
	$(MAKE) build GOOS=linux GOARCH=amd64
	$(MAKE) build GOOS=linux GOARCH=arm
	$(MAKE) build GOOS=openbsd GOARCH=386
	$(MAKE) build GOOS=openbsd GOARCH=amd64
	$(MAKE) build GOOS=solaris GOARCH=amd64
	$(MAKE) build GOOS=windows GOARCH=386
	$(MAKE) build GOOS=windows GOARCH=amd64

clean:
	[ ! -d bin ] || rm -rfv bin

# test: 
# 	go test -i $(TEST) || exit 1                                                   
# 	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

# testacc: 
# 	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   