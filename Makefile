start-service: 
	cd tools/ && docker compose up -d --build 

clean:
	docker container rm $$(docker ps -aq) -f
