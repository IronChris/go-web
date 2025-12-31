# Variables - Change these if your Pi IP or username changes
#PI_USER = chris
PI_TARGET = chess-pi
PI_PATH = ~/fide-app
BINARY_NAME = server_pi

.PHONY: build upload restart deploy

# 1. Build the binary for Pi 2B (ARMv7)
build:
	@echo "Building for Raspberry Pi 2B..."
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(BINARY_NAME) .



stop:
	@echo "Stopping service..."
	ssh -t $(PI_TARGET) "sudo systemctl stop fide.service"

upload:
	@echo "Uploading..."
	scp $(BINARY_NAME) $(PI_TARGET):$(PI_PATH)/
	scp -r ./web $(PI_TARGET):$(PI_PATH)/
	scp -r ./web/docs $(PI_TARGET):$(PI_PATH)/
start:
	@echo "Starting service..."
	ssh -t $(PI_TARGET) "sudo systemctl start fide.service"

deploy: build stop upload start
