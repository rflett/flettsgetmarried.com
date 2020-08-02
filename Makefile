deps:
	npm install

dev:
	npm run dev

deploy: deps
	npm run build
	npm run export
	echo "Upload dist/ to S3"
