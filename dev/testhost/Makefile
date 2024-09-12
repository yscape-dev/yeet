.PHONY: build run stop connect

build:
	docker build -t yeet-testhost .

run:
	docker run -p 23502:22 -d --name yeet-testhost-container yeet-testhost

stop:
	docker rm -f yeet-testhost-container

connect:
	ssh -o StrictHostKeyChecking=no -p 23502 root@localhost