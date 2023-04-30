curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | VERIFY_CHECKSUM=false sh
curl -O https://s3.us-west-2.amazonaws.com/amazon-eks/1.25.6/2023-01-30/bin/linux/amd64/kubectl
chmod +x ./kubectl
mkdir -p $HOME/bin && cp ./kubectl $HOME/bin/kubectl && export PATH=$PATH:$HOME/bin
aws eks update-kubeconfig --name $EKS_CLUSTER --region $AWS_DEFAULT_REGION
helm repo add stable https://charts.helm.sh/stable

# Add the Helm repository
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

# Update Helm repositories
helm repo update

# next deploy
if kubectl get namespace prometheus &> /dev/null; then
    echo "Deleting existing prometheus namespace..."
    kubectl delete namespace prometheus
fi 

echo "Creating prometheus namespace..."
kubectl create namespace prometheus


echo "Installing prometheus and grafana..."
helm install prometheus prometheus-grafana/kube-prometheus-stack --namespace prometheus --set alertmanager.enabled=false --set grafana.service.type=LoadBalancer