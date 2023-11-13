# Go Pods Poc
Este projeto é uam poc em Go que busca gerenciar pods de uma aplicação a partir de outra aplicação.

## Dependências
* Go versão 20
* Kind
* Kubectl

## Formas de rodar
* Localmente
* Cluster

### Rodando locamente
[Doc Aqui](./docs/Local.README.md)

### Rodando no cluster
[Doc Aqui](./docs/cluster.README.md)

## Respostas

#### Sucesso

#### /pods/amount/update/{amount} 
```
Pods updated: {amount}
```

#### /health
```
Hey, I'm alive and running at local machine
```

```
Hey, I'm alive and running at Cluster
```

#### Erro
Dependerá de onde estourar erro, mas via de regraq virá uma string contendo uma descrição do erro acompanhado de sua causa.
Ex.:
```
Error building config from flags. Reason: stat ./kubeconfig.yaml: no such file or directory
```

### Lista de comandos úteis
[Doc Aqui](./docs/commands.README.md)