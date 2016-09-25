all: build

clean:
	if [ -a user_auto.go ] ;\
	then \
		rm user_auto.go ; \
	fi \

build:
	go generate .

test:
	go test .

.PHONY: all
