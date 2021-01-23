build:
	docker build . -t hpm/looseleaf --no-cache

run:
	docker run -it --rm hpm/looseleaf

push: