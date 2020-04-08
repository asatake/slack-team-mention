GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=main
SRC_PATH="src/main.go"

all: build
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ${SRC_PATH}
	./$(BINARY_NAME)
deps:
	${GOGET} "github.com/aws/aws-lambda-go/events"
	${GOGET} "github.com/aws/aws-lambda-go/lambda"
	${GOGET} "github.com/slack-go/slack"
	${GOGET} "gopkg.in/yaml.v2"
build-linux:
	echo "Linux"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v ${SRC_PATH}
build-macos:
	echo "Mac OS"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v ${SRC_PATH}

ifeq ($(shell uname),Linux)
build: build-linux
else ifeq ($(shell uname),Darwin)
build: build-macos
else
echo "OS not supported."
endif
