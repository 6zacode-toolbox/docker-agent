# docker-agent
An image with drivers to engage with docker host


# How this project was started

```bash 
go mod init github.com/6zacode-toolbox/docker-agent
cobra-cli init
go mod tidy
go get github.com/6zacode-toolbox/docker-operator/operator/api/v1
cobra-cli add test
go run . test 
```