UNAME=$(shell uname)
DT = $(shell date +%FT%T%Z)
WEB=./website
AWS_PROFILE_PROD = roofix-stage
AWS_PROFILE_STAGE = roofix-stage
FN_DIR = ./bin
FN_BUILD_DIR = ./cdk.out/build
ENT = entgo.io/ent/cmd/ent
##
##
clean-node-packages:
	cd $(WEB) && rm -rf node_modules
install-node-packages:
	cd $(WEB) && pnpm i
# remove and reinstall node packages
install-node-packages-fresh: clean-node-packages install-node-packages

#
# dev, serve website
start-website:
	cd $(WEB) && pnpm dev
# dev, server website backend
start-api:
	go run ./bin/server

test-ui:
	cd $(WEB) && npx playwright test --ui

pnpm-outdated:
	cd $(WEB) && pnpm outdated
pnpm-update:
	cd $(WEB) && pnpm update

# Graph commands
gql-gen:
	go generate ./server
# generate client side typings needed in website UI
# before invoke, make sure that graph server is running
gql-codegen:
	cd $(WEB) && pnpm run codegen

# ent init an entity e.g. > make ent-init name=EntityName
ent-init:
	go run $(ENT) new $(name)
# ent generate files from schema
ent-gen:
	go generate ./ent
# local ent schema migration
ent-migrate:
	go run ./bin/migrate
ent-migrate-with-admins:
	go run ./bin/migrate -with-admin
ent-migrate-with-all:
	go run ./bin/migrate -with-admin -with-postal -with-pricing
# ent schema describe
ent-describe:
	go run $(ENT) describe ./ent_schema

cdk-bootstrap-stage:
	cdk --profile $(AWS_PROFILE_STAGE) bootstrap
cdk-deploy-stage:
	cdk --profile $(AWS_PROFILE_STAGE) deploy

cdk-bootstrap-prod:
	STACK_ENV=prod cdk --profile $(AWS_PROFILE_PROD) bootstrap
cdk-deploy-prod:
	STACK_ENV=prod cdk --profile $(AWS_PROFILE_PROD) deploy

#
# update lambda STAGE ==>
#
update-clearExpiredJobInvitesFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,clearExpiredJobInvites-stage,clearExpiredJobInvitesFn,stage)
update-clearExpiredSessionsFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,clearExpiredSessions-stage,clearExpiredSessionsFn,stage)
update-docUploadFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,docUpload-stage,docUploadFn,stage)
update-mailerFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,mailer-stage,mailerFn,stage)
update-mailFeedbackFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,mailFeedback-stage,mailFeedbackFn,stage)
update-migrateFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,migrate-stage,migrateFn,stage)
invoke-migrateFn-stage:
	@printf '\n=> Invoking Lambda function: migrate-stage ==\n'
	aws --profile $(AWS_PROFILE_STAGE) lambda invoke --function-name migrate-stage --payload '{}' /dev/stdout
update-notificationFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,notification-stage,notificationFn,stage)
update-taskFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,task-stage,taskFn,stage)
# update both frontend & backend changes
update-websiteFn-stage:
	$(call deploy-assets,$(AWS_PROFILE_STAGE),assets.roofixstaging.com,stage)
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,website-stage,websiteFn,stage)
# update only the backend-server
update-server-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,website-stage,websiteFn,stage)
# update only the frontend-ui changes
update-ui-stage:
	$(call deploy-assets,$(AWS_PROFILE_STAGE),assets.roofixstaging.com,stage)
update-htmltopdfSM-step01-stage:
	$(call update-lambda-02,$(AWS_PROFILE_STAGE),arm64,html-to-pdf-step01-stage,htmltopdfSM/step01,html-to-pdf-step01-stage,stage)
update-htmltopdfSM-step03-stage:
	$(call update-lambda-02,$(AWS_PROFILE_STAGE),arm64,html-to-pdf-step03-stage,htmltopdfSM/step03,html-to-pdf-step03-stage,stage)
# websocket API
update-wsApiFn-stage:
	$(call update-lambda,$(AWS_PROFILE_STAGE),arm64,wsApi-stage,wsApiFn,stage)
#
# update lambda PROD ==>
#
update-clearExpiredJobInvitesFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,clearExpiredJobInvites-PROD,clearExpiredJobInvitesFn,production)
update-clearExpiredSessionsFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,clearExpiredSessions-PROD,clearExpiredSessionsFn,production)
update-docUploadFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,docUpload-PROD,docUploadFn,production)
update-mailerFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,mailer-PROD,mailerFn,production)
update-mailFeedbackFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,mailFeedback-PROD,mailFeedbackFn,production)
update-migrateFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,migrate-PROD,migrateFn,production)
invoke-migrateFn-prod:
	@printf '\n=> Invoking Lambda function: migrate-PROD ==\n'
	aws --profile $(AWS_PROFILE_PROD) lambda invoke --function-name migrate-PROD --payload '{}' /dev/stdout
update-notificationFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,notification-PROD,notificationFn,stage)
update-taskFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,task-PROD,taskFn,production)
# update both frontend & backend changes
update-websiteFn-prod:
	$(call deploy-assets,$(AWS_PROFILE_STAGE),app-assets.rfx.services,production)
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,website-PROD,websiteFn,production)
# update only the backend-server
update-server-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,website-PROD,websiteFn,production)
# update only the frontend-ui changes
update-ui-prod:
	$(call deploy-assets,$(AWS_PROFILE_STAGE),app-assets.rfx.services,production)
update-htmltopdfSM-step01-prod:
	$(call update-lambda-02,$(AWS_PROFILE_STAGE),arm64,html-to-pdf-step01-PROD,htmltopdfSM/step01,html-to-pdf-step01-PROD,production)
update-htmltopdfSM-step03-prod:
	$(call update-lambda-02,$(AWS_PROFILE_STAGE),arm64,html-to-pdf-step03-PROD,htmltopdfSM/step03,html-to-pdf-step03-PROD,production)

# websocket API
update-wsApiFn-prod:
	$(call update-lambda,$(AWS_PROFILE_PROD),arm64,wsApi-PROD,wsApiFn,production)

# ~~~


sync-email-tpls-stage:
	@printf '\n=> Sync Email Templates ==\n'
	aws --profile $(AWS_PROFILE_STAGE) s3 sync --delete ./template/email/ s3://data.roofixstaging.com/mail-templates
sync-email-tpls-prod:
	@printf '\n=> Sync Email Templates(prod) ==\n'
	aws --profile $(AWS_PROFILE_PROD) s3 sync --delete ./template/email/ s3://data.rfx.services/mail-templates
sync-pdf-templates-stage:
	@printf '\n=> Sync Email Templates ==\n'
	aws --profile $(AWS_PROFILE_STAGE) s3 sync --delete ./template/pdf/ s3://data.roofixstaging.com/pdf-templates
sync-pdf-templates-prod:
	@printf '\n=> Sync Email Templates(prod) ==\n'
	aws --profile $(AWS_PROFILE_PROD) s3 sync --delete ./template/pdf/ s3://data.rfx.services/pdf-templates


build-assets:
	cd $(WEB) && ASSETS_HOST="https://$(domain)" pnpm run build --mode $(mode)

build-lambda:
	mkdir -p $(FN_BUILD_DIR)/$(name)
	GOOS=linux GOARCH=$(arch) CGO_ENABLED=0 go build -o $(FN_BUILD_DIR)/$(name)/bootstrap \
		-ldflags="-s -w -X 'roofix/config.AppEnv=$(env)' -X 'roofix/config.UseSecretMgr=true' \
		-X 'roofix/config.BuildUser=$(USER)' -X 'roofix/config.BuildTime=$(DT)'" \
		$(FN_DIR)/$(path)/main.go
	upx --best --lzma $(FN_BUILD_DIR)/$(name)/bootstrap

# update-lambda fn code
# ex. $(call update-lambda,$(AWS_PROFILE_STAGE):1, arm64:2, fnName:3, fnCodeDir:4 ,env:5)
define update-lambda
	@printf '\n=> Update Lambda: $(3) Env: stage ==\n'
	mkdir -p $(FN_BUILD_DIR)/$(4)
	GOOS=linux GOARCH=$(2) CGO_ENABLED=0 go build -o $(FN_BUILD_DIR)/$(3)/bootstrap \
		-ldflags="-s -w -X 'roofix/config.AppEnv=$(5)' -X 'roofix/config.UseSecretMgr=true' \
		-X 'roofix/config.BuildUser=$(USER)' -X 'roofix/config.BuildTime=$(DT)'" \
		$(FN_DIR)/$(4)/main.go
	upx --best --lzma $(FN_BUILD_DIR)/$(3)/bootstrap
	cp $(FN_BUILD_DIR)/$(3)/bootstrap ./bootstrap
	zip ./$(3).zip ./bootstrap
	rm  ./bootstrap
	aws --profile $(1) lambda update-function-code --no-paginate --function-name $(3) --zip-file fileb://./$(3).zip
	rm  ./$(3).zip
	@printf '\n=> DONE with updating lambda $(3) ==\n'
endef

# update-lambda fn code
# ex. $(call update-lambda,$(AWS_PROFILE_STAGE):1, arm64:2, fnName:3, fnCodeDir:4, lambdaName:5 ,env:6)
define update-lambda-02
	@printf '\n=> Update Lambda: $(3) Env: stage ==\n'
	mkdir -p $(FN_BUILD_DIR)/$(4)
	GOOS=linux GOARCH=$(2) CGO_ENABLED=0 go build -o $(FN_BUILD_DIR)/$(3)/bootstrap \
		-ldflags="-s -w -X 'roofix/config.AppEnv=$(6)' -X 'roofix/config.UseSecretMgr=true' \
		-X 'roofix/config.BuildUser=$(USER)' -X 'roofix/config.BuildTime=$(DT)'" \
		$(FN_DIR)/$(4)/main.go
	upx --best --lzma $(FN_BUILD_DIR)/$(3)/bootstrap
	cp $(FN_BUILD_DIR)/$(3)/bootstrap ./bootstrap
	zip ./$(3).zip ./bootstrap
	rm  ./bootstrap
	aws --profile $(1) lambda update-function-code --no-paginate --function-name $(5) --zip-file fileb://./$(3).zip
	rm  ./$(3).zip
	@printf '\n=> DONE with updating lambda $(3) ==\n'
endef

define deploy-assets
	@printf '\n=> Building Website UI ==\n'
	cd $(WEB) && ASSETS_HOST="https://$(2)" pnpm run build --mode $(3)
	@printf '\n=> syncing $(WEB)/build to s3://$(2) ==\n'
	aws --profile $(1) s3 sync --delete $(WEB)/build/ s3://$(2)
endef


bintree-build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./server-bin ./bin/websiteFn/main.go
	go tool nm -size ./server-bin | c++filt > ./bintree/out/symtab.txt
	python3 ./bintree/tab2pydic.py ./bintree/out/symtab.txt > ./bintree/out/out.py
	python3 ./bintree/simplify.py ./bintree/out/out.py > ./bintree/data.js
	rm ./server-bin

bintree-serve:
	cd ./bintree && python3 -m http.server

build-compress:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./server-bin ./bin/server/main.go
	upx --best --lzma -o ./server ./server-bin
