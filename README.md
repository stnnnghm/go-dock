# go-dock

![source: https://hackernoon.com/hn-images/1*JfSp7LWmVE1nj15IrxWSWQ.png](https://hackernoon.com/hn-images/1*JfSp7LWmVE1nj15IrxWSWQ.png)

## go-dock is a simple application that uses Gradle to build and run a Docker container with an instance of a Golang web browser hosting a simple joke about me!

### How It Works

1. A simple web server is created in `server.go`.
2. The `Dockerfile` configures the Docker image (`golang:latest`) to be used, and which file (`server.go`) will be turned into an executable file and run.
3. `gradle.build` defines the plugins (`com.palantir.docker` and `com.palantir.docker-run`) that will be used to build and run the Docker container, along with establishing the steps (`docker`, `dockerRun`) used to accomplish this.

### Use

To run the application enter the following commands in your terminal:

1. `./gradlew docker`
2. `.gradlew dockerRun`
   The application will build and run a docker container with an instance of the webserver on port :8081.
3. `curl localhost:8081` or use your web browser to navigate to `localhost:8081` to see a joke!
