makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

build:
	go build

install:
	go install

commit:
	git add . || true
	git commit -m "$(m)"
	git push origin master

test: generate
	go test -v ./...

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

publish:
	make commit m=$(tag)
	make tag tag=$(tag)