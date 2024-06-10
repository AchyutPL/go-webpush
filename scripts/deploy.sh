# build the image
cd go-app

echo "ARGS $1"

# decompress the go-app tar file and remove the tar file
tar -xzf go-app.tar.gz -C . && rm -f go-app.tar.gz

pwd && ls

# build the image using the file Dockerfile.deploy
docker build -f ./build/package/Dockerfile.deploy -t go-app:latest .

# remove the container
docker container rm -f go-app || true

# remove dangling docker images
docker image prune -f

# run the docker container using go-app:latest image
docker run -d --network kong-net -p 8080:8080 \
  -e PORT=8080 \
  -e DB_HOST=$1 \
  -e DB_PORT=5432 \
  -e DB_DATABASE=postgres \
  -e DB_USERNAME=postgres \
  -e DB_PASSWORD=postgres \
  --name go-app go-app:latest
