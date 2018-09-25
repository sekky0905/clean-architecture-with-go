# パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gobinary

.PHONY: init
init: clean deps test precommit build

.PHONY: all
all: test precommit build run

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

.PHONY: test
test:
	$(GOCMD) test ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: run
run:
	./$(BINARY_NAME)

.PHONY: deps
deps:
	$(GOGET) github.com/golang/dep/cmd/dep
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) golang.org/x/lint/golint
	$(GOGET) golang.org/x/tools/cmd/goimports
	$(GOGET) github.com/kisielk/errcheck
	$(GOGET) gopkg.in/DATA-DOG/go-sqlmock.v1
	$(GOGET) github.com/go-sql-driver/mysql
	dep ensure

.PHONY: precommit
precommit :
	# 静的解析
	go vet ./...
	# go list で import path を表示する
	# pipe を使用した後、 xargs でコマンドラインを作成し、xargs で作成したコマンドラインは、 go list の表示結果
	go list ./... | xargs golint -set_exit_status
	# エラーハンドリングの確認
	# test と　Close の部分を無視している
	errcheck -ignoretests -ignore 'Close' ./...
