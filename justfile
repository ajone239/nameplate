# Run the application
run: frontend
    go run .

# Build the application for the pi
build_pi: frontend
    mkdir -p output/web
    cp -r web/build output/web/build
    GOOS=linux GOARCH=arm GOARM=5 go build -o output/nameplate

# Build the application
build: frontend
    go build

# Build the frontend
frontend:
    (cd ./web && npm run build)
