default:
	go build ./...
run:
	./test
clean:
	find -type f -executable -exec rm {} \;
