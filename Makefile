build:
	go build -ldflags "-s -w" -o build/server cmd/main.go

clean:
	rm -rf build

install:
	sudo cp build/server /usr/local/bin/lock-server

daemon-rm:
	sudo rm /etc/systemd/system/lock-server.service
	sudo systemctl daemon-reload

daemon: clean build install
	sudo cp lock-server.service /etc/systemd/system/lock-server.service
	sudo systemctl daemon-reload
	sudo systemctl enable lock-server.service
