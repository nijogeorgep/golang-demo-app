app_name := go-booking-app
app_version := $(git rev-parse HEAD)

build_image:
	@echo "Building" $(app_name)
	docker build -t $(app_name):$(app_version) .

run_container:
	@echo "Running" $(app_name)
	docker run $(app_name):$(app_version) -p 8080:8080

up: build_image run_container

run:
	go run .