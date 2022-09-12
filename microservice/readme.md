# first way
docker build  -t microservice .
docker run --publish 8001:8005 microservice

# docker compose
docker-compose up -d 
or 
docker-compose -f docker-compose.yml up -d 

docker-compose down 
or 
docker-compose -f docker-compose.yml down 

## Container shell access and viewing MongoDB logs
 docker exec -it  microservice_mongo_db_1 bash
## open shell for mongo 
 mongo database

# The MongoDB Server log is available through Docker's container log:

$ docker logs microservice_mongo_db_1


k8s is orchestration tools
