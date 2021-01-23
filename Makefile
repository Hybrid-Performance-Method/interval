build:
	docker build . -t hpm/interval --no-cache

docker: build
	docker run -it --rm hpm/interval
push:

interval: build-bin exec
build-bin:
	go build -o ./interval main.go 
exec:
	./interval