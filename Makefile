AWS_ECR_REGION := us-east-1
AWS_ECR_REPOSITORY := mansei

.PHONY: docker-build
docker-build:
	docker build -t ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_ECR_REGION}.amazonaws.com/${AWS_ECR_REPOSITORY}:${TAG} app

.PHONY: docker-push
docker-push:
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_ECR_REGION}.amazonaws.com/${AWS_ECR_REPOSITORY}:${TAG}
