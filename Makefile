# Docker
# ======

.PHONY: docker-build
docker-build:
	docker build -t ssouthcity/visitorz:latest .

.PHONY: docker-push
docker-push:
	docker push ssouthcity/visitorz:latest

.PHONY: docker-build-and-push
docker-build-and-push: docker-build docker-push


# Kubernetes
# ==========

.PHONY: k8s-view
k8s-view:
	kubectl get all

.PHONY: k8s-apply
k8s-apply:
	kubectl apply --recursive -f k8s/

.PHONY: k8s-purge
k8s-purge:
	kubectl delete all --all

.PHONY: k8s-update-images
k8s-update-images:
	kubectl rollout restart deploy


# Workflows
# =========

.PHONY: full-deployment
full-deployment: docker-build-and-push k8s-apply k8s-update-images

