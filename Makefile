cluster:
	@kind create cluster

kubeconfig:
	@kind get kubeconfig > ./service-one/kubeconfig.yaml

docker-images:
	@cd ./service-one/ ; docker build -t svc-1:v1 . ; cd ../service-two/. ; docker build -t svc-2:v1 .

docker-images-to-cluster:
	@kind load docker-image svc-1:v1
	@kind load docker-image svc-2:v1

set-pods:
	@cd ./service-one/k8s ; kubectl apply -f deployment.yaml ; kubectl apply -f service.yaml
	@cd ./service-two/k8s ; kubectl apply -f deployment.yaml ; kubectl apply -f service.yaml

port-forward:
	@kubectl port-forward svc/service-one 8081:8081 &
	@kubectl port-forward svc/service-two 8080:8080 &

drop-docker-images:
	@docker rmi -f svc-1:v1
	@docker rmi -f svc-2:v1

drop-cluster:
	@kind delete cluster

drop-kubeconfig:
	@rm ./service-one/kubeconfig.yaml

unset-deployments:
	@cd ./service-one/k8s ; kubectl delete -f deployment.yaml ; kubectl delete -f service.yaml
	@cd ./service-two/k8s ; kubectl delete -f deployment.yaml ; kubectl delete -f service.yaml