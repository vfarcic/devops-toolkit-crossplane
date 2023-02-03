export NAME=${1:-a-team-gke}

export CONFIG_PATH=${2:-kubeconfig.yaml}

export PROJECT_ID=$3

export KUBECONFIG=$CONFIG_PATH

gcloud container clusters get-credentials $NAME --region us-east1 \
    --project $PROJECT_ID

chmod 600 $CONFIG_PATH
