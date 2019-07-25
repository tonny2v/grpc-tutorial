docker build . -t tonny2v/api-service:v1.0
docker push tonny2v/api-service:v1.0
kubectl create -f api-service.yaml


# kubectl get pods -w
# minikube service api-service --url
# curl http://192.168.99.100:30080/add/1/3