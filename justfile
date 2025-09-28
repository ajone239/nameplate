run: frontend
    go run .

build: frontend
    go build

frontend:
    (cd ./web && ls && npm run build)
