SHELL := /usr/bin/bash
.ONESHELL:
CLIENT_PLATFORM ?= $(shell go env GOOS)

ifeq ($(CLIENT_PLATFORM),linux)
export CURRENT_HOST_IP=$(shell hostname -I | awk '{print $1; exit}' | cut -d ' ' -f 1)
else
export CURRENT_HOST_IP=$(shell ifconfig en0 | awk '/inet / {print $2; }' | cut -d ' ' -f 2)
endif

ifneq (,$(wildcard ./.env))
    include .env
    export
    ENV_FILE_PARAM = --env-file .env
endif

VAULT_ADDR=http://$(CURRENT_HOST_IP):8200

export

reset-minikube:
	minikube stop > /dev/null 2>&1 || true
	minikube delete > /dev/null 2>&1 || true
	./reset-minikube.sh
	minikube start --vm-driver=virtualbox --kubernetes-version=v1.20.7
	cp /home/$(USER)/.minikube/ca.crt build/
	cp /home/$(USER)/.minikube/profiles/minikube/client.key build/
	cp /home/$(USER)/.minikube/profiles/minikube/client.crt build/

start-vault:
	nohup vault server -config=vault/config.hcl > vault.out  2>&1 &

stop-vault:
	pkil vault
