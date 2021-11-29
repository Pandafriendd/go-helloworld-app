# Go Http App

### Run the App locally
If you are not familiar with golang, here is how you could start Hello Service on your local box:
	- Install golang: ``` https://golang.org/dl/ ```
	- Git clone this repo
	- ``` cd api/hello ```
	- ``` go build -o bin/hello cmd/main.go && go run cmd/main.go ```
	- ``` curl -X POST http://localhost:8081/hello -d '{"s":"world"}' ```
	- Note that you can build this service into an executable and run it everywhere
