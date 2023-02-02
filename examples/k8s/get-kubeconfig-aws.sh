export NAME=${1:-a-team-eks}

export CONFIG_PATH=${2:-kubeconfig.yaml}

aws eks update-kubeconfig --region us-east-1 --name $NAME \
    --kubeconfig $CONFIG_PATH

chmod 600 $CONFIG_PATH
