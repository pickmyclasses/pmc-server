## The server side of PickMyClasses

### Prerequisites for running the server

- Go 1.8 is required to be installed on the machine if desired to run locally
- The connection to our **backup** database info is included in config.yaml
- There is no further requirements for setting up the databases, we have remote databases deployed
- If running locally, please change app:
  name: "pmc"
  mode: "release" to mode: "dev", if remotely, change it to "release"

### Running the server

- If on Docker, run Docker build . then run the image 
- If on local machine, go to root directory of the project, run `go mod tidy` first to install the libraries, then run `go run main.go` or `go build .`
- You should be able to run the server and see the API portals running.