cd ./signer && ./updateInput.sh && cd -
cat algoTransfer/rawTx.txt >>./signer/id-10001-ed-input.yaml
cat algoTransfer/rawTx.txt >>./signer/id-10002-ed-input.yaml
cat algoTransfer/rawTx.txt >>./signer/id-10003-ed-input.yaml

./sign.sh 1

