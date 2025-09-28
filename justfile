# Run the application
run: frontend
    go run .

# Build the application
build_pi: frontend
    mkdir -p output/web
    cp -r web/build output/web/build
    GOOS=linux GOARCH=arm go build -o output/nameplate

# Build the application
build: frontend
    go build

# Build the frontend
frontend:
    (cd ./web && npm run build)
