aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 963485456147.dkr.ecr.us-east-1.amazonaws.com
docker build -t ftd/s3batch-handler-status .
docker tag ftd/s3batch-handler-status:latest 963485456147.dkr.ecr.us-east-1.amazonaws.com/ftd/s3batch-handler-status:latest
docker push 963485456147.dkr.ecr.us-east-1.amazonaws.com/ftd/s3batch-handler-status:latest