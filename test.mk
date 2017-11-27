.PHONY: test

test: hello
	echo "Makefile!"

hello:
	echo "Hello"

%:
	echo "Wild"