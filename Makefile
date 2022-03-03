build:
	docker build . -t ogr0:latest 


shell:
	docker run --rm -it ogr0:latest sh


run:
	docker run --rm -it -p 2222:2222/tcp ogr0:latest

test:
	docker-compose -f docker-compose.test.yml up -V --force-recreate --exit-code-from sut --abort-on-container-exit --build
	docker-compose -f docker-compose.test.yml down
