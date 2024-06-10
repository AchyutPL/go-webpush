# build the image
cd go-app

docker build -f build/package/Dockerfile.deploy -t go-app:latest .
docker container rm -f go-app || true
docker image prune -f go-app

docker run -d --network kong-net -e \
  "PORT=8080" \
  "DB_HOST=localhost" \
  "DB_PORT=5433" \
  "DB_DATABASE=postgres" \
  "DB_USERNAME=postgres" \
  "DB_PASSWORD=postgres" \
   --name go-app go-app:latest
