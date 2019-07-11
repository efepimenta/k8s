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
    * ```git clone .........```
* Acessar o diretório criado
    * ```cd ...```

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
not yet
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
### Desligando tudo
```bash
minikube stop
```