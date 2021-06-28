# Vault

## Ref
<https://blockchain-automation-framework.readthedocs.io/en/develop/developer/dev_prereq.html>

## Setup

1. Follow the instruction to install Vault on your machine -> https://www.vaultproject.io/downloads/

1. Create a config.hcl file in the project directory with the following contents (use a file path in the path attribute which exists on your local machine) For example, /home/users/Desktop/project/vault.

```hcl
ui = true
storage "file" {
   path    = "$(pwd)/data"
}

listener "tcp" {
   address     = "127.0.0.1:8200"
   tls_disable = 1
}

```

1. Start the Vault server by executing (this will occupy one terminal). Do not close this terminal.
```shell
vault server -config=config.hcl
```

1. Open browser at http://localhost:8200/. And initialize the Vault by providing your choice of key shares and threshold.

1. Click Download Keys or copy the keys, you will need them. Then click Continue to Unseal. Provide the unseal key first and then the root token to login.

1. In a new terminal, execute the following (assuming vault is in your PATH):
```shell
export VAULT_ADDR='http://<Your Vault local IP address>:8200' #e.g. http://192.168.2.121:8200
export VAULT_TOKEN="<Your Vault root token>"
vault secrets enable -version=1 -path=secret kv
```


> On Linux, to give the Vault executable the ability to use the mlock syscall without running the process as root, run:
```shell
sudo setcap cap_ipc_lock=+ep $(readlink -f $(which vault))
```