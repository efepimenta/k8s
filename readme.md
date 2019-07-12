### Imagem Go no docker hub
```bash
https://hub.docker.com/r/efepimenta/go-wserver
```

### Passos para que o "cluster" funcione

* [Instalar o docker](https://docs.docker.com/install/)
* [Instalar o kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* [Instalar o minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
* [Instalar o VirtualBox](https://www.virtualbox.org/wiki/Downloads)
* Criar o Node local
    * ```minikube start```
* Clonar este repositório
    * ```git clone https://github.com/efepimenta/k8s.git```
* Acessar o diretório criado
    * ```cd k8s```

---

### Criando o pod nginx
```bash
kubectl apply -f nginx/configmap.yaml
kubectl apply -f nginx/deployment.yaml
kubectl apply -f nginx/service.yaml

minikube service nginx-service
```

### Criando o pod mysql
```bash
kubectl apply -f mysql/volume.yaml
kubectl create secret generic mysql-pass --from-literal=password='sua-senha-aqui'
kubectl apply -f mysql/deployment.yaml
kubectl apply -f mysql/service.yaml
```

### Criando o pod go
```bash
kubectl apply -f go/deployment.yaml
kubectl apply -f go/service.yaml

minikube service go-service

```


### Conferindo se tudo está ok
```bash
kubectl get configmaps
kubectl get deployments
kubectl get services
kubectl get persistentvolumes
kubectl get secrets
kubectl get pods
```

### Apagando tudo
```bash
kubectl delete configmaps --all
kubectl delete deployments --all
kubectl delete services --all
kubectl delete persistentvolumes --all
kubectl delete secrets --all
kubectl delete pods --all
```

### Desligando tudo
```bash
minikube stop
```
