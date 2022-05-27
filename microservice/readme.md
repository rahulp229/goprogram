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


k8s is orchestration tools
