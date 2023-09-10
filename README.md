# tcp-server
An adaptable multi-threaded TCP server equipped with built-in rate limiting, affording you the flexibility to tailor thread configurations and rate limits to your specific needs. It comes with metrics support like total request processed by server at a given time, useful for users to analyse the server performance.

## Installation
### Prerequisite
1. [Docker](https://docs.docker.com/desktop/install/mac-install/)
2. Code editor.

### Steps to run server
1. git clone https://github.com/Avish34/tcp-server
2. cd tcp-server
3. make start-service

>      ⠿ Container tools-prometheus-1  Started                                                                                                                          0.5s
>      ⠿ Container tools-tcp-server-1  Started                                                                                                                     0.5s
After running (3) command, use should see this output.
Server will be listening on 8080. To edit the port and other settings, you can find customise section below.

Send request to the server. You should see "Hello world as response"

> curl http://localhost:8080                                         
> Hello world !

### Metrics
To check the total request accepted by the server. Open http://localhost:9090/. Search for *total_request* in Expression search bar. Check the below image for reference.

## Customisation
This server comes with capability to throttle request using rate limiting, serving multiple request using workers. This is completly customisable, it can be done via changing the .env file present in

>  tools/dockerfiles/tcp-server/.env
