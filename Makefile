default:
	-rm -rf monitor
	mkdir monitor
	go build -o monitor ./...
run:
	./monitor/pkg
clean:
	find -type f -executable -exec rm {} \;
