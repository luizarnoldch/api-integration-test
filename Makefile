TEMPLATE_FILE := templates/main.yml
STACK_NAME := api-integration-test

init:
	go mod init main
update:
	go mod tidy
mock:
	mockery --all --output ./tests/mocks
build:
	./scripts/build.sh
unit:
	go test ./tests/unit/...
integration:
	go test ./tests/integration/...
f_test:
	./scripts/func_test.sh
deploy:
	sam deploy --template-file $(TEMPLATE_FILE) --stack-name $(STACK_NAME) --capabilities CAPABILITY_NAMED_IAM --resolve-s3
destroy:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
tt:
	git pull
	make build
	make deploy
