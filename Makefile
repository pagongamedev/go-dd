ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif
  
ifeq ($(OS),Windows_NT)
build:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";

compile:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";
	go build -o app.exe .\main.go; .\app.exe

run:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";
	go run .\main.go

else

build:
	env GO111MODULE=on go build -o app main.go;

compile:
	env GO111MODULE=on go build -o app main.go;
	env API_ENV=localhost ./app

run:
	env GO111MODULE=on API_ENV=localhost go run main.go

endif

update-tag:
	git tag v${tag} ; git push origin v${tag}