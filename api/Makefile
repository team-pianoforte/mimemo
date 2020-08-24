.PHONY: dev
dev:
	dev_appserver.py app.yaml --port=4000

deploy:
	gcloud app deploy  --projectId=${GCLOUD_PROJECT_ID}

watch:
	looper
