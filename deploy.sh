
# Obtener Account ID desde AWS CLI
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
AWS_REGION="us-east-1"
REPO_NAME="ftd/s3batch-handler-status"
LAMBDA_FUNCTION_NAME="ftdBatchStatusS3Updater"
IMAGE_TAG="latest"

echo "🔐 Usando AWS Account ID: $AWS_ACCOUNT_ID"

# === BUILD ===
echo "🔧 Construyendo imagen Docker..."
docker build -t $REPO_NAME .

# === TAG ===
echo "🏷️ Etiquetando imagen para ECR..."
docker tag $REPO_NAME:$IMAGE_TAG $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$REPO_NAME:$IMAGE_TAG

# === PUSH ===
echo "🚀 Subiendo imagen a ECR..."
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$REPO_NAME:$IMAGE_TAG

# === UPDATE FUNCTION ===
echo "📦 Actualizando función Lambda..."
aws lambda update-function-code \
  --function-name $LAMBDA_FUNCTION_NAME \
  --image-uri $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$REPO_NAME:$IMAGE_TAG \
  --region $AWS_REGION

echo "✅ ¡Despliegue completo!"