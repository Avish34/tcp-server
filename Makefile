start-service: 
	cd tools/ && docker compose up -d

clean:
	docker container rm $$(docker ps -aq) -f
