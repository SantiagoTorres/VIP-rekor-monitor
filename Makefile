default:
	-rm -rf monitor
	mkdir monitor
	go build -o monitor ./...
run:
	./monitor/pkg
clean:
	find -type f -executable -exec rm {} \;

server: 
	docker run -p 3306:3306 --env "MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=yes" mariadb

db: 
	mysql -u root -P 3306 
