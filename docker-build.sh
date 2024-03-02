version=$(git rev-parse --short HEAD)

docker build -t ethorbit/discord-server-status:latest .
docker tag ethorbit/discord-server-status:latest ethorbit/discord-server-status:$version

