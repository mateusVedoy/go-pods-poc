# Go Pods Poc
Este projeto é uam poc em Go que busca gerenciar pods de uma aplicação a partir de outra aplicação.

## Comandos úteis

### Criando cluster via Kind
```
kind create cluster --name <cluster-name>
```

### Deletando cluster via Kind
```
kind delete cluster --name <cluster-name>
```

### Gerando imagem local para pod
```
docker build -t <name>:<tag> /path/to/Dockerfile
```

### Importando docker container image para cluster
```
<cluster-name> load docker-image <name>:<tag>
```

### Se precisar limpar docker images e containers
```
docker rmi -f $(docker images -aq)
docker container prune
```

### Criando pod dentro do cluster
```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### Verificando pod
```
kubectl get po
```

### Verificando service
```
kubectl get svc
```

### deletando pod
```
kubectl delete pod <pod-name>

```

### Assistir pods em tempo de exec
```
watch 'kubectl get po'

```

### Acessando cluster (docker container)
```
docker exec -it <container-name> sh
```

### Acessando pod dentro do cluster
Primeiro executa passo acima

```
kubectl exec -it <pod-name> -- sh
```

### Bind de portas pod
```
kubectl port-forward svc/<service-name> <port-local>:<port-exposed-by-docker-image>
```