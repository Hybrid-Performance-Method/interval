build:
	docker build . -t hpm/interval --no-cache

run: build
	docker run -it --rm hpm/interval

push:

interval:
	go build -o ./interval main.go
	./interval