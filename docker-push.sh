version=$(git describe --tags)

docker push mgerb/server-status:latest;
docker push mgerb/server-status:$version;

