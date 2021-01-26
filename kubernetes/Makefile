list_node:
	kubectl get node
detail_node:
	kubectl describe node

list_pod:
	kubectl get pod
detail_pod:
	kubectl describe pod nginx
create_pod:
	kubectl create -f pod/nginx.yaml
run_pod:
	kubectl port-forward nginx 8888:80
delete_pod:
	kubectl delete pod nginx-label

list_pod_label:
	kubectl get pod --show-labels
list_pod-label:
	kubectl get pod -l team
list_pod-label-value:
	kubectl get pod -l team=IT
delete_pod-label:
	kubectl delete pod -l team

list_namespace:
	kubectl get namespace
list_pod-namespace:
	kubectl get pod -n default
create_namespace:
	kubectl create -f namespace/staging.yaml
create_pod-namespace:
	kubectl create -f pod/nginx.yaml -n staging
# will delete pod to
delete_namespace:
	kubectl delete namespace staging
delete_pod-namespace:
	kubectl delete pod nginx -n staging
delete_allpod-namespace:
	kubectl delete pod --all -n staging
