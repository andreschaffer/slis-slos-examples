.PHONY: ensure-local-context deploy-prometheus deploy-grafana deploy-traefik deploy-simple-service all clean

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
