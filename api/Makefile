cover-html:
	@echo "mode: count" > coverage-all.out

	go test -p 1 ./... -coverprofile=coverage.out -covermode=count
	tail -n +2 coverage.out >> coverage-all.out
	go tool cover -html=coverage-all.out -o coverage.html
	rm coverage.out coverage-all.out
	open ./coverage.html

cover:
	@echo "mode: count" > coverage-all.out
	
	go test -p 1 ./... -coverprofile=coverage.out -covermode=count
	tail -n +2 coverage.out >> coverage-all.out
	go tool cover -func coverage-all.out | grep total:
	rm coverage.out coverage-all.out
