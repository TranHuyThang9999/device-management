
run:
	go run cmd/main.go

clean:
	go clean

commit:
	@git add .
	@git commit -m "add new feature"