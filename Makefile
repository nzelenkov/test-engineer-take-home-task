.PHONY: build
build:
	@echo 'build docker image'
	@cat Dockerfile | envsubst > .Dockerfile
	@docker build -f .Dockerfile -t $(TESTIMGREPO):$(TESTIMGTAG) .
	@rm -f .Dockerfile
	@docker push $(TESTIMGREPO):$(TESTIMGTAG)

.PHONY: clean
clean:
	@echo 'clean up resources'
	@docker image prune -f --filter label=stage=builder

.PHONY: deploy
deploy:
	@echo 'deploy to k8s'
	@microk8s.helm install --set image.repository=$(IMGREPO) --set image.tag=$(IMGTAG) --set test.name=$(APPNAME) --set test.image.repository=$(TESTIMGREPO) --set test.image.tag=$(TESTIMGTAG) --set service.port=$(HTTPPORT) --name deploy deploy
	@microk8s.helm test deploy
	@microk8s.helm delete --purge deploy
	@microk8s.kubectl delete pod/deploy-test-block-height
