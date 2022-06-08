# optistech
**optistech** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite-hq/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/hengmengsroin/optistech@latest! | sudo bash
```
`hengmengsroin/optistech` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)



# Runing a node

### Initialize the Chain
```
./optistechd init validator1 --chain-id optistech
```

### Adding keys to the keyring
```
./optistechd keys add my_validator --keyring-backend test
```
```
MY_VALIDATOR_ADDRESS=$(./optistechd keys show my_validator -a --keyring-backend test)
```
### Adding Genesis Accounts

```
./optistechd add-genesis-account $MY_VALIDATOR_ADDRESS 100000000000stake
```

### Create a gentx.
```
./optistechd gentx my_validator 100000000stake --chain-id optistech --keyring-backend test
```

### Add the gentx to the genesis file.
```
./optistechd collect-gentxs
```

### run localnet
```
./optistechd start
```
### show node id
```
./optistechd tendermint show-node-id
```
# Interacting with the Node
```
./optistechd query bank balances <address>
```

```
./optistechd keys add recipient --keyring-backend test
```
```
RECIPIENT=$(./optistechd keys show recipient -a --keyring-backend test)
```
```
./optistechd tx bank send $MY_VALIDATOR_ADDRESS $RECIPIENT 1000000stake --chain-id optistech --keyring-backend test
```
### Check that the recipient account did receive the tokens.
```
./optistechd query bank balances $RECIPIENT --chain-id optistech
```

```
./optistechd tx staking delegate $(./optistechd keys show my_validator --bech val -a --keyring-backend test) 500stake --from recipient --chain-id optistech --keyring-backend test
```
### Query the total delegations to `validator`.
```
./optistechd query staking delegations-to $(simd keys show my_validator --bech val -a --keyring-backend test) --chain-id optistech
```
