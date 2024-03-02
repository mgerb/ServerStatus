version=$(git rev-parse --short HEAD)

docker push ethorbit/discord-server-status:latest;
docker push ethorbit/discord-server-status:$version;

