sonic:
	go build cmd/sonic.go

db:
	go build tools/build-db.go
	./build-db
	@rm build-db
	@mv sonic.db ./tests

tool:
	go build ./tools/totcoverage.go
	go build ./tools/build-db.go
	@mv totcoverage build-db ./tools
test:
	go test -cover ./lib...

total:
	@go build tools/totcoverage.go
	go test -cover ./lib... | ./totcoverage
	@rm totcoverage

clean:
	@if [ -a sonic ] ;\
		then \
		rm sonic; \
	fi;
	@if [ -a tools/build-db ] ;\
		then \
		rm tools/build-db; \
	fi;
	@if [ -a tools/totcoverage ] ;\
		then \
		rm tools/totcoverage; \
	fi;

