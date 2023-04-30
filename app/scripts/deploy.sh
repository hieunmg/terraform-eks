curl -O https://s3.us-west-2.amazonaws.com/amazon-eks/1.25.6/2023-01-30/bin/linux/amd64/kubectl
chmod +x ./kubectl
mkdir -p $HOME/bin && cp ./kubectl $HOME/bin/kubectl && export PATH=$PATH:$HOME/bin
aws eks update-kubeconfig --name $EKS_CLUSTER --region $AWS_DEFAULT_REGION
aws ecr get-login-password --region $AWS_DEFAULT_REGION | docker login --username AWS --password-stdin $ECR_REGISTRY
kubectl apply -f eks/aws-auth.yaml
kubectl apply -f eks/deployment.yaml
kubectl apply -f eks/service.yaml