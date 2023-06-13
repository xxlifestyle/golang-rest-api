build:
	go build -C cmd/main/ -o ../../bin.exe
run: 
	bin.exe
br:
	go build -C cmd/main/ -o ../../bin.exe && bin.exe