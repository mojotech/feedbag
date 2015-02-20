test-server:
	go get ./...
	cd feedbag && go test ./...

release:
	@make test-server
	cd client && gulp build
	rm -rf public
	mv ./client/dist ./public

deps:
	cd client && npm i && bower i
	go get ./...

watch-server:
	@gin

watch-client:
	cd client && gulp serve

.PHONY: test-serve rrelease deps watch-server watch-client
