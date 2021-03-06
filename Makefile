GROUP = $(shell whoami)
NAME = $(shell basename $$(git rev-parse --show-toplevel))
VERSION = $(shell cat VERSION)

.PHONY: all
all: start teardown

.PHONY: help
help:
	@echo "\n$(GROUP)/$(NAME):$(VERSION)\n"
	@cat README.md
	@echo ""

.PHONY: image
image:
	@docker build -t $(GROUP)/$(NAME):$(VERSION) .

.PHONY: start
start:
	@docker run --rm \
		--name $(NAME) \
		-p 8181:8181 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		\
		$(GROUP)/$(NAME):$(VERSION)

.PHONY: bootstrap
bootstrap: 
	$(MAKE) image
	$(MAKE) start

.PHONY: publish
publish:
	$(MAKE) image
	\
	git checkout -B release/$(VERSION)
	git add --all
	git commit -a --allow-empty-message -m '' 2> /dev/null || :
	git rebase --no-ff --autosquash release/$(VERSION)
	git pull origin release/$(VERSION) 2> /dev/null || :
	git push origin release/$(VERSION)
	git checkout master
	\
	docker login -u $(GROUP) https://index.docker.io/v1/
	docker push $(GROUP)/$(NAME)
	\
	$(MAKE) teardown

.PHONY: teardown
teardown:
	@docker stop $(NAME) 2>/dev/null || :
