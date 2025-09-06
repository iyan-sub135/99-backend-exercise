deps:
	cd services/listing-service && go mod tidy
	cd services/user-service && go mod tidy
	cd services/public-api && go mod tidy

run-listing:
	cd services/listing-service && go run .

run-user:
	cd services/user-service && go run . -port=7001

run-public:
	cd services/public-api && go run .
