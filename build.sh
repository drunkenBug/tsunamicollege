echo "building frontend"
docker build -t frontend frontend/
echo "building backend"
docker build -t backend backend/
echo "remove and clean all running containers"
docker kill $(docker ps -q)
docker container prune -f
echo "starting frontend"
docker run -d -p 8080:8080 frontend
echo "starting backend"
docker run -d -p 3080:3080 backend
