list_node:
	kubectl get node
detail_node:
	kubectl describe node

list_pod:
	kubectl get pod
show_log:
	kubectl logs namepod
detail_pod:
	kubectl describe pod nginx
create_pod:
	kubectl create -f pod/nginx.yaml
run_pod:
	kubectl port-forward nginx 8888:80
delete_pod:
	kubectl delete pod nginx-label
delete_pod-all:
	kubectl delete pod --all

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

list_rc:
	kubectl get rc
create_rc:
	kubectl create -f replication_controller/nginx-rc.yaml
delete_rc:
	kubectl delete rc nginx-rc
delete_rc-casecade:
	kubectl delete rc nginx-rc --cascade=false

list_rs:
	kubectl get rs
create_rs:
	kubectl create -f replica_set/nginx-rs.yaml
delete_rs:
	kubectl delete rs nginx-rs
delete_rs-casecade:
	kubectl delete rs nginx-rs --cascade=false

list_ds:
	kubectl get daemonset
create_ds:
	kubectl create -f daemon_set/nginx-ds.yaml
delete_ds:
	kubectl delete ds daemon-nginx
delete_ds-casecade:
	kubectl delete ds daemon-nginx --cascade=false

list_job:
	kubectl get job
create_job:
	kubectl create -f job/nodejs-job.yaml
delete_job:
	kubectl delete job nodejs-job

list_cj:
	kubectl get cj
create_cj:
	kubectl create -f cron_job/nodejs-cj.yaml
delete_cj:
	kubectl delete cj nodejs-cj
