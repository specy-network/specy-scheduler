# Specy Scheduler

## Introduction

**Specy Scheduler** is a component that listens to and manages on-chain tasks and triggers the automation engine

The main steps are:
1. Listen on-chain events
2. Register tasks
3. Trigger automation engine
4. Execute task result on-chain

## Project setup
```
make install
```

### Config
set config file 'specy-config.yaml' in the root directory
```
chain_info: #specy chain config
  chain_id: 
  chain_binary_location: 
  home_dir: 
  node_address: 
  validator_wallet_address: 

engine_info: #specy engine config
  engine_node_address: 
  ias_report: 
  enclave_pk: 
```

### Start
start sheduler
```
sch start --home <relayer home dir>
```

add target chain binary location
```
sch chains add <target chain id> <target chain binary location>
```
