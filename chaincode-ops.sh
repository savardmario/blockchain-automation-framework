#!/bin/bash
set -e

echo "Starting build process..."

echo "Adding env variables..."
export PATH=/root/bin:$PATH

#Path to k8s config file
export KUBECONFIG=/home/blockchain-automation-framework/build/config

echo "Running Chaincode Ops playbook..."
cd platforms/hyperledger-fabric/configuration/
exec ansible-playbook -vv chaincode-ops.yaml -e "@./network.yaml" -e "add_new_org='false'"
