# build the image
cd go-app

# decompress the go-app tar file and remove the tar file
tar -xzf go-app.tar.gz -C . && rm -f go-app.tar.gz

pwd && ls

# build the image using the file Dockerfile.deploy
docker build -f ./build/package/Dockerfile.deploy -t go-app:latest .

# remove the container
docker container rm -f go-app || true

# remove dangling docker images
docker image prune -f


echo "HOST IS" $DB_HOST
echo "SERVER IS" $DEV_HOST

# run the docker container using go-app:latest image
docker run -d --network kong-net -p 8080:8080 \
  -e PORT=8080 \
  -e DB_HOST=$DB_HOST \
  -e DB_PORT=$DB_PORT \
  -e DB_DATABASE=$DB_DATABASE \
  -e DB_USERNAME=$DB_USERNAME \
  -e DB_PASSWORD=$DB_PASSWORD \
  --name go-app go-app:latest
