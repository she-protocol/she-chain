#!/usr/bin/env sh

# Input parameters
NODE_ID=${ID:-0}
NUM_ACCOUNTS=${NUM_ACCOUNTS:-5}
echo "Configure and initialize environment"

cp build/shed "$GOBIN"/
cp build/price-feeder "$GOBIN"/

# Prepare shared folders
mkdir -p build/generated/gentx/
mkdir -p build/generated/exported_keys/
mkdir -p build/generated/node_"$NODE_ID"

# Testing whether shed works or not
shed version # Uncomment the below line if there are any dependency issues
# ldd build/shed

# Initialize validator node
MONIKER="she-node-$NODE_ID"

shed init "$MONIKER" --chain-id she >/dev/null 2>&1

# Copy configs
ORACLE_CONFIG_FILE="build/generated/node_$NODE_ID/price_feeder_config.toml"
APP_CONFIG_FILE="build/generated/node_$NODE_ID/app.toml"
TENDERMINT_CONFIG_FILE="build/generated/node_$NODE_ID/config.toml"
cp docker/localnode/config/app.toml "$APP_CONFIG_FILE"
cp docker/localnode/config/config.toml "$TENDERMINT_CONFIG_FILE"
cp docker/localnode/config/price_feeder_config.toml "$ORACLE_CONFIG_FILE"


# Set up persistent peers
SHE_NODE_ID=$(shed tendermint show-node-id)
NODE_IP=$(hostname -i | awk '{print $1}')
echo "$SHE_NODE_ID@$NODE_IP:26656" >> build/generated/persistent_peers.txt

# Create a new account
ACCOUNT_NAME="node_admin"
echo "Adding account $ACCOUNT_NAME"
printf "12345678\n12345678\ny\n" | shed keys add "$ACCOUNT_NAME" >/dev/null 2>&1

# Get genesis account info
GENESIS_ACCOUNT_ADDRESS=$(printf "12345678\n" | shed keys show "$ACCOUNT_NAME" -a)
echo "$GENESIS_ACCOUNT_ADDRESS" >> build/generated/genesis_accounts.txt

# Add funds to genesis account
shed add-genesis-account "$GENESIS_ACCOUNT_ADDRESS" 10000000ushe,10000000uusdc,10000000uatom

# Create gentx
printf "12345678\n" | shed gentx "$ACCOUNT_NAME" 10000000ushe --chain-id she
cp ~/.she/config/gentx/* build/generated/gentx/

# Creating some testing accounts
echo "Creating $NUM_ACCOUNTS accounts"
python3 loadtest/scripts/populate_genesis_accounts.py "$NUM_ACCOUNTS" loc >/dev/null 2>&1
echo "Finished $NUM_ACCOUNTS accounts creation"

# Set node shevaloper info
SHEVALOPER_INFO=$(printf "12345678\n" | shed keys show "$ACCOUNT_NAME" --bech=val -a)
PRIV_KEY=$(printf "12345678\n12345678\n" | shed keys export "$ACCOUNT_NAME")
echo "$PRIV_KEY" >> build/generated/exported_keys/"$SHEVALOPER_INFO".txt

# Update price_feeder_config.toml with address info
sed -i.bak -e "s|^address *=.*|address = \"$GENESIS_ACCOUNT_ADDRESS\"|" $ORACLE_CONFIG_FILE
sed -i.bak -e "s|^validator *=.*|validator = \"$SHEVALOPER_INFO\"|" $ORACLE_CONFIG_FILE

echo "DONE" >> build/generated/init.complete
