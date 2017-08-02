GROUP = $(shell whoami)
NAME = $(shell basename $$(git rev-parse --show-toplevel))
VERSION = $(shell cat VERSION)

.PHONY: all start bootstrap image publish teardown info

all: start teardown

info:
	@echo "\n$(GROUP)/$(NAME):$(VERSION)\n"
	@cat README.md
	@echo ""

image:
	docker build -t $(GROUP)/$(NAME):$(VERSION) .

start:
	@echo "Starting version : $(VERSION)"
	docker run -p 8181:8181 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		\
		$(GROUP)/$(NAME):$(VERSION)

bootstrap: 
	$(MAKE) image
	$(MAKE) start

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

teardown:
	@docker stop $(GROUP)/$(NAME)
	@docker rm -f $(GROUP)/$(NAME)
