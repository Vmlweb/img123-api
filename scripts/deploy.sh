#!/usr/bin/env bash

docker build -t img123:latest .
docker tag img123:latest us.gcr.io/vmlweb/img123:latest
docker push us.gcr.io/vmlweb/img123:latest
gcloud run deploy img123-api --project vmlweb --image us.gcr.io/vmlweb/img123:latest --platform managed --region us-central1 --allow-unauthenticated --memory 1024