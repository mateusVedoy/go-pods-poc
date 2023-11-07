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
```

### Verificando pod
```
kubectl get po
```

### Verificando service
```
kubectl get svc
```

### Assistir pods em tempo de exec
```
watch 'kubectl get po'