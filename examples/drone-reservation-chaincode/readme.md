# Instructions

### Installer les dépendances
- A partir de `/chaincode/`
- `go mod tidy`
- `go mod vendor`

### Rouler les tests unitaires
- `go test drone-reservation/chaincode -v`
### Rouler les tests unitaires avec couverture
- `go test drone-reservation/chaincode -- cover`

### Build le chaincode
- `go build -o drone-reserver`

### Deployer en mode dev
- Suivre les instructions [https://hyperledger-fabric.readthedocs.io/en/latest/peer-chaincode-devmode.html]()
- Assurez vous d'utiliser la release 2.2 de fabric
- Vous aurez peut-être besoin de créer vous même le répertoire
  `/var/hyperldedger` et de lui octroyer toutes les permissions manuellement `chmod 777 /var/hyperledger`
- À l'étape "Build the chaincode", copier plutôt l'exécutable `drone-reserver` à la racine du répertoire `/fabric/`
- À l'étape "Start the chaincode", remplacer `./simpleChaincode` par `./drone-reserver`, vous aurez donc:
  `CORE_CHAINCODE_LOGLEVEL=debug CORE_PEER_TLS_ENABLED=false CORE_CHAINCODE_ID_NAME=mycc:1.0 ./drone-reserver -peer.address 127.0.0.1:7052`
- Il ne l'est pas indiqué dans le tutorial, mais l'étape "Start the chaincode" bloque le terminal, vous devrez donc
ouvrir un nouveau terminal pour l'étape "Approve and commit the chaincode definition", ensuite settez les variables d'
  environnement suivantes: `export PATH=$(pwd)/build/bin:$PATH` et `export FABRIC_CFG_PATH=$(pwd)/sampleconfig`
- Vous pouvez maintenant intéragir avec le chaincode pour tester les fonctions du smartcontract les commandes suivantes:  
`CORE_PEER_ADDRESS=127.0.0.1:7051 peer chaincode invoke -o 127.0.0.1:7050 -C ch1 -n mycc -c '{"Args":["InitLedger"]}' --isInit`
  `CORE_PEER_ADDRESS=127.0.0.1:7051 peer chaincode invoke -o 127.0.0.1:7050 -C ch1 -n mycc -c '{"Args":["CreateAsset", "asset2", "monday", "1, 2, 3", "requested", "lenn", ""]}'`  
  `CORE_PEER_ADDRESS=127.0.0.1:7051 peer chaincode invoke -o 127.0.0.1:7050 -C ch1 -n mycc -c '{"Args":["GetAllAssets"]}'`  
  devraient retourner respectivement  
  `2021-03-10 12:50:46.330 EST [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 001 Chaincode invoke successful. result: status:200 `  
  `2021-03-10 12:53:43.253 EST [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 001 Chaincode invoke successful. result: status:200 `  
  `2021-03-10 12:53:47.808 EST [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 001 Chaincode invoke successful. result: status:200 payload:"[{\"ID\":\"asset1\",\"date\":\"1994-11-05T13:15:30Z\",\"request\":\"125, 134, 898\",\"status\":\"requested\",\"owner\":\"ClientOrg\",\"location\":\"\"},{\"ID\":\"asset2\",\"date\":\"monday\",\"request\":\"1, 2, 3\",\"status\":\"requested\",\"owner\":\"requested\",\"location\":\"\"}]"]}' `
- Notez que le ledger est présentement initialisé avec un asset requested pour des fins de test.