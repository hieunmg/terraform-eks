aws ecr get-login-password --region $AWS_DEFAULT_REGION | docker login --username AWS --password-stdin $ECR_REGISTRY
aws secretsmanager get-secret-value --secret-id $SECRET_MANANGER_ID --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]' > app.env
docker build -t $ECR_REGISTRY/$ECR_REPOSITORY:$CI_COMMIT_SHORT_SHA -t $ECR_REGISTRY/$ECR_REPOSITORY:latest .
docker push $ECR_REGISTRY/$ECR_REPOSITORY