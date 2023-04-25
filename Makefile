GO ?= go 
default:
	-rm -rf monitor
	mkdir monitor
	$(GO) build -o monitor ./...
run:
	./monitor/pkg
clean:
	find -type f -executable -exec rm {} \;

server: 
	docker run -d -p 3306:3306 --env "MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=yes" \
		--name=main-db --rm mariadb

server-database-init:
	@echo $(shell docker exec -i main-db mariadb -uroot -P 3306 < pkg/database/schema.sql)


db: 
	mysql -u root -P 3306 
