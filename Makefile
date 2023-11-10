.PHONY: ensure-local-context deploy-traefik deploy-prometheus deploy-grafana deploy-simple-service clean tunnel
ensure-local-context:
	kubectl config use-context docker-desktop

deploy-traefik: ensure-local-context
	kubectl apply \
		-f code-example/traefik/deployment/

deploy-prometheus: ensure-local-context
	kubectl apply \
		-f code-example/prometheus/deployment/

deploy-grafana: ensure-local-context
	kubectl apply \
		-f code-example/grafana/deployment/

deploy-simple-service: ensure-local-context
	docker build \
		-t simple-service \
		-f code-example/simple-service/build/Dockerfile \
		code-example/simple-service/.
	kubectl apply \
    	-f code-example/simple-service/deployment/

all: deploy-prometheus deploy-traefik deploy-simple-service

clean:
	kubectl delete -f code-example/simple-service/deployment/
	kubectl delete -f code-example/prometheus/deployment/
	kubectl delete -f code-example/traefik/deployment/

tunnel:
	kubectl port-forward service/prometheus-service 9090 -n prometheus & \
	kubectl port-forward service/simple-service 8888 8889 -n simple-service & \
	kubectl port-forward service/traefik-dashboard-service 8080 -n traefik & \
	kubectl port-forward service/traefik-metrics-service 8082 -n traefik & \
	kubectl port-forward service/traefik-web-service 8081 -n traefik & \
	kubectl port-forward service/grafana-service 3000 -n grafana &
