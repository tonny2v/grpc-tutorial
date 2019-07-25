docker build . -t shuzasa/summation-service:v1.0
docker push shuzasa/summation-service:v1.0
kubectl create -f summation-service.yaml