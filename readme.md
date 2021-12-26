api-backend-tester:
    build: ./api/app
    ports:
      - "5000:7000"
    volumes:
      - ./api/app/:/go/src/app
    command: bash -c "go mod download && go build -ldflags '-s -w' ./main.go && air && go run ."
    depends_on:
      - mongo-backend-tester
  