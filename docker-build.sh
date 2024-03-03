version=$(git describe --tags)

docker build -t mgerb/server-status:latest .
docker tag mgerb/server-status:latest mgerb/server-status:$version

