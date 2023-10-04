docker compose down
git pull origin dev
docker network create --driver bridge my_local_network || true
docker compose up --build