.PHONY: ensure-local-context deploy-prometheus deploy-grafana deploy-traefik deploy-simple-service all clean tunnel

ensure-local-context:
	kubectl config use-context docker-desktop

deploy-prometheus: ensure-local-context
	kubectl apply \
		-f prometheus/deployment/

deploy-grafana: ensure-local-context
	kubectl apply \
		-f grafana/deployment/

deploy-traefik: ensure-local-context
	kubectl apply \
		-f traefik/deployment/

deploy-simple-service: ensure-local-context
	docker build \
		-t simple-service \
		-f simple-service/build/Dockerfile \
		simple-service/.
	kubectl apply \
    	-f simple-service/deployment/

all: deploy-prometheus deploy-grafana deploy-traefik deploy-simple-service

clean:
	kubectl delete -f simple-service/deployment/
	kubectl delete -f traefik/deployment/
	kubectl delete -f grafana/deployment/
	kubectl delete -f prometheus/deployment/

tunnel:
	kubectl port-forward service/prometheus-service 9090 -n prometheus & \
	kubectl port-forward service/grafana-service 3000 -n grafana & \
	kubectl port-forward service/traefik-dashboard-service 8080 -n traefik & \
	kubectl port-forward service/traefik-metrics-service 8082 -n traefik & \
	kubectl port-forward service/traefik-web-service 8081 -n traefik & \
	kubectl port-forward service/simple-service 8888 8889 -n simple-service &
