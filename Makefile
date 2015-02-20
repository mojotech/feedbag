release:
	@make test
	cd client && gulp build
	rm -rf public
	mv ./client/dist ./public

test:
	cd feedbag && go test ./...

deps:
	cd client && npm i && bower i
	go get ./...

watch-server:
	@gin

watch-client:
	cd client && gulp serve

.PHONY: release test deps watch-server watch-client
