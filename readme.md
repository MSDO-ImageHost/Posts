# Posts / p3s

![Docker Image CI](https://github.com/MSDO-ImageHost/Posts/workflows/Docker%20Image%20CI/badge.svg)

![Publish Docker image](https://github.com/MSDO-ImageHost/Posts/workflows/Publish%20Docker%20image/badge.svg)

---
## Documentation
[API](docs/api-spec.md) \
[DB models](docs/db-models.md)

---
## Resources
Go RabbitMQ: 
- https://github.com/kedacore/sample-go-rabbitmq
- http://www.inanzzz.com/index.php/post/iamo/creating-a-rabbitmq-producer-example-with-golang


## How to use

**Local development with docker compose and .devcontainer**
Requires:
* Docker
* Docker Compose
* VS Code 

## Kubernetes deployment

```shell
$ kubectl apply -f https://raw.githubusercontent.com/MSDO-ImageHost/Posts/main/deploy/posts-deployment.yaml
$ kubectl apply -f https://raw.githubusercontent.com/MSDO-ImageHost/Posts/main/deploy/posts-service.yaml

```