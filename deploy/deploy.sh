docker compose down
git pull origin dev
docker network create --driver bridge bi || true
docker compose up --build