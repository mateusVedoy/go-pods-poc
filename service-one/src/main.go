package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var USER = os.Getenv("USER")

func main() {
	http.HandleFunc("/svc-two/pods/amount/update/", podsHandler)
	http.HandleFunc("/svc-one/health/", healthHandler)

	http.ListenAndServe(":8081", nil)
}

func podsHandler(w http.ResponseWriter, r *http.Request) {
	amountStr := extractAmount(r.URL.Path, "/svc-two/pods/amount/update/")
	amount := normalizeValue(amountStr)
	//criar endpoint p pegar qtd de pods
	err := scaleDeploymentPods("default", "service-two", amount)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "Pods updated: %s", amountStr)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	var enviroment string

	if isInClusterConfig() {
		enviroment = "K8S Cluster"
	} else {
		enviroment = "local machine"
	}

	fmt.Fprintf(w, "Hey, I'm alive and running at "+enviroment)
}

func normalizeValue(value string) int32 {
	intVal, err := strconv.Atoi(value)

	if err != nil {
		panic(err.Error())
	}

	return int32(intVal)
}

func extractAmount(path, prefix string) string {
	return path[len(prefix):]
}

func outClusterConfig() (*kubernetes.Clientset, error) {

	kubeconfigPath := "kubeconfig.yaml"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error building config from flags. Reason: %v", err.Error()))
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error building new config from kubernetes. Reason: %v", err.Error()))
	}

	return clientset, nil

}

func inClusterConfig() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error config in cluster config. Reason: %v", err.Error()))
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error config kubernetes new config. Reason: %v", err.Error()))
	}

	return clientset, nil
}

func isInClusterConfig() bool {
	return USER == "inCluster"
}

func scaleDeploymentPods(namespace, deploymentName string, replicas int32) error {

	var clientset *kubernetes.Clientset
	var err error

	if isInClusterConfig() {
		clientset, err = inClusterConfig()
	} else {
		clientset, err = outClusterConfig()
	}

	if err != nil {
		return err
	}

	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	deployment, err := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})

	if err != nil {
		return errors.New(fmt.Sprintf("Error getting deployment. Reason: %v", err.Error()))
	}

	deployment.Spec.Replicas = &replicas

	_, err = deploymentsClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})

	if err != nil {
		return errors.New(fmt.Sprintf("Error updating deployment. Reason: %v", err.Error()))
	}

	return nil
}
