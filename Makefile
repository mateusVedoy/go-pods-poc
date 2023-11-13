## in / out cluster

cluster:
	@kind create cluster


## out cluster

kubeconfig:
	@kind get kubeconfig > ./service-one/kubeconfig.yaml

out-docker-img:
	@cd ./service-one/ ; docker build -t svc-1:v1 -f outCluster.Dockerfile . ; cd ../service-two/. ; docker build -t svc-2:v1 .

out-set-pods:
	@cd ./service-two/k8s ; kubectl apply -f deployment.yaml ; kubectl apply -f service.yaml

out-docker-img-to-cluster:
	@kind load docker-image svc-2:v1

out-port-forward:
	@kubectl port-forward svc/service-one 8081:8081 &
	@kubectl port-forward svc/service-two 8080:8080 &

## in cluster

in-docker-img:
	@cd ./service-one/ ; docker build -t svc-1:v1 -f inCluster.Dockerfile . ; cd ../service-two/. ; docker build -t svc-2:v1 .

in-docker-img-to-cluster:
	@kind load docker-image svc-1:v1
	@kind load docker-image svc-2:v1

in-set-pods:
	@cd ./service-one/k8s ; kubectl apply -f deployment.yaml ; kubectl apply -f service.yaml
	@cd ./service-two/k8s ; kubectl apply -f deployment.yaml ; kubectl apply -f service.yaml

in-port-forward:
	@kubectl port-forward svc/service-one 8081:8081 &
	@kubectl port-forward svc/service-two 8080:8080 &

## Drop configs

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