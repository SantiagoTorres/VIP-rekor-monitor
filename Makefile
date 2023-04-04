GO ?= go 
default:
	-rm -rf monitor
	mkdir monitor
	$(GO) build -o monitor ./...
run:
	./monitor/pkg
clean:
	find -type f -executable -exec rm {} \;
