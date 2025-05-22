
[← Regresar](../README.md) <br>

---
## 📋 Core library
[🌐 Documentación](https://github.com/miguel-armas-abt/backend-core-library) <br>

---

## ▶️ Despliegue local

1. Generar binario
```sh
go build -o runner ./main.go
```

2. Configurar las [variables de entorno](./variables_local.env) en el IDE.

3. Ejecutar aplicación


---

## ▶️ Despliegue con Docker

⚙️ Crear imagen
```shell
docker build -t miguelarmasabt/repository-finder:v1.0.1 -f ./Dockerfile .
```

⚙️ Ejecutar contenedor
```shell
docker run --rm -p 8080:8080 --env-file ./variables_local.env --name repository-finder-v1  miguelarmasabt/repository-finder:v1.0.1
```

---

## ▶️ Despliegue con Kubernetes

⚙️ Encender Minikube
```shell
docker context use default
minikube start
```

⚙️ Crear imagen
```shell
eval $(minikube docker-env --shell bash)
docker build -t miguelarmasabt/repository-finder:v1.0.1 -f ./Dockerfile .
```

⚙️ Crear namespace y aplicar manifiestos
```shell
kubectl create namespace repos
kubectl apply -f ./k8s.yaml -n repos
```

⚙️ Eliminar orquestación
```shell
kubectl delete -f ./k8s.yaml -n repos
```

⚙️ Port-forward
```shell
kubectl port-forward <pod-id> 8080:8080 -n repos
```
