#!/bin/bash

graphnode_data_path="./graphnode-data"
deploy_home_dir=$(pwd)
#if debug is true ,the gov voting period time will be set to "180s"
debug=true
LOG_ERROR="ERROR"
LOG_WARNING="WARNING"
LOG_SUCCESS="SUCCESS"
CHAIN_ID="ibc-2"
REGULATORYD=`which iris`
help() {
    cat << EOF
Usage:
  
    -clean <Clean chain data>
    -check <Env check >
    -deploy <Deploy chain>          
    -graphnode <Start graph node>   
    -proposal <Send rule, binding,register and relation proposal> 
    -manifest <Deploy system manifest to graph node >
    -transfer <transfer test (after system proposal passed)>

    -help <Help>
EOF
exit 0
}

issue_nft_denom()
{
    iris tx nft issue class222  --name=class222 --from=$($REGULATORYD  --home ./data/ibc-1 keys show validator --keyring-backend test -a) \
     --mint-restricted=false --update-restricted=false --chain-id=ibc-1 \
     --keyring-backend=test --home=./data/ibc-1 \
     --node=tcp://localhost:26557
}

mint_nft_token()
{
    iris tx nft mint class222 token222 --recipient=$($REGULATORYD  --home ./data/ibc-1 keys show validator --keyring-backend test -a)  \
    --from=$($REGULATORYD  --home ./data/ibc-1 keys show validator --keyring-backend test -a)  --chain-id=ibc-1 --keyring-backend=test \
    --home=./data/ibc-1 --node=tcp://0.0.0.0:26557
}

ibc_nft_transfer()
{
    iris tx nft-transfer transfer nft-transfer channel-0 $($REGULATORYD  --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
    class222 token222 --from=$($REGULATORYD  --home ./data/ibc-1 keys show validator --keyring-backend test -a)  --chain-id=ibc-1 \
    --keyring-backend=test --home=./data/ibc-1 --node=tcp://0.0.0.0:26557
}

parse_params()
{
    case $1 in 
        clean) clean
        ;;
        deploy) deploy_chain 
        ;;
        check) check_env
        ;;
        graphnode) deploy_graph_node
        ;;
        proposal) system_proposal
        ;;
        manifest) deploy_system_manifest
        ;;
        transfer) transfer_test $2 $3
        ;;
        help) help
        ;;
        issue) issue_nft_denom
        ;;
        mint) mint_nft_token
        ;;
        ibc-nft-transfer)ibc_nft_transfer
        ;;
    esac
}



clean(){
    read -p "Do you want to remove chain network and delete data? (y/n) " answer
    if [[ "$answer" == "y" ]]; then
        log "start remove graphnode and regchain node ."
            # select graph-node tmux session 
        if tmux has-session -t graphnode 2>/dev/null; then
            # if exsist ,kill it
            log $LOG_SUCCESS "kill tmux graphnoe session"
            tmux kill-session -t graphnode
        else
            # if is not exisis 
            log $LOG_SUCCESS "Tmux session graphnode does not exist"
        fi
        killall firehose-cosmos


        log "start remove graphnode data."
        cd $graphnode_data_path
        if [ ! -d "docker-compose.yml" ]; then
            docker-compose down
        fi

        graphnode_data_save_path='./data'
        if [ -d "$graphnode_data_save_path" ]; then
            rm -rf "$graphnode_data_save_path"
        else
            log $LOG_WARNING "$graphnode_data_save_path does not exist."
        fi
        cd $deploy_home_dir
        log "start remove logs ."
        rm -rf './logs/firehose.log'
        rm -rf './logs/regchain.log'
        rm -rf './logs/graphnode.log'

        log $LOG_SUCCESS "clean end"

    elif [[ "$answer" == "n" ]]; then
        log "Exiting..."
    else
        log $LOG_ERROR "Invalid input. Please enter 'y' or 'n'."
    
    fi

    
}

check_env(){
    log "===========================CHECK ENV==================================="
    log "check regulatoryd env......"
    iris version > /dev/null
        if [ $? -eq  0 ]; then
            log $LOG_SUCCESS "regulatoryd already installed!"
        else
            log $LOG_ERROR "regulatoryd is not install"
            exit
        fi
    log "check firehose-cosmos env......"
    firehose-cosmos --version > /dev/null
        if [ $? -eq  0 ]; then
            log $LOG_SUCCESS  "firehose-cosmos already installed!"
        else
            log $LOG_ERROR "firehose-cosmos is not install"
            exit
        fi
    log "check graph-node env......"
    graph-node --version > /dev/null
        if [ $? -eq  0 ]; then
            log $LOG_SUCCESS  "graph-node already installed!"
        else
            log $LOG_ERROR "graph-node is not install"
            exit
        fi
    
    log "check compliance-engine env......"
    if [ -e /usr/local/bin/compliance-engine ]; then
        log $LOG_SUCCESS  "compliance-engine already installed!"
    else
        log $LOG_ERROR "compliance-engine is not install"
            exit
    fi

    log "check docker env......"
    docker -v > /dev/null
        if [ $? -eq  0 ]; then
            log $LOG_SUCCESS  "docker already installed!"
        else
            log $LOG_ERROR "docker is not install"
            exit
        fi

    log "check docker-compose env......"
    docker-compose -v > /dev/null
        if [ $? -eq  0 ]; then
            log $LOG_SUCCESS  "docker-compose already installed!"
        else
            log $LOG_ERROR "docker-compose is not install"
            exit
        fi 

    log "check nodejs env......"
    node --version > /dev/null
        if [ $? -eq  0 ]; then
            log $LOG_SUCCESS  "node is alredy installed!"
        else
            log $LOG_ERROR "nodejs is not install"
            exit
        fi
    log "checkout node version and install yarn"
    npm --version > /dev/null
        if [ $? -eq  0 ]; then
            log $LOG_SUCCESS  "npm is alredy installed!"
        else
            log $LOG_ERROR "npm is not install"
            exit
        fi
    node_version=$(node -v)
    required_version="v16.0.0"
    if version_compare $node_version  $required_version; then
        log $LOG_SUCCESS  "node version is up to date"
    else
        log $LOG_ERROR "node verison need GE 16.0.0"
        exit
    fi
    if npm list -g --depth=0 yarn >/dev/null 2>&1; then
        log $LOG_SUCCESS "yarn is already installed globally"
    else
        log $LOG_ERROR "yarn not installed"
        exit
    fi
    log "=============================================================="
}


system_proposal(){
    
    log "register proposal"
    $REGULATORYD tx regulatory \
        submit-register-proposal ./source/proposal/register.json \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Submit register proposal failed ,stop !"
        exit 1
    else
        log $LOG_SUCCESS "Submit register proposal succeeded"
    fi

    sleep 10
    $REGULATORYD tx gov vote 1 yes \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Vote register proposal failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Vote register proposal succeeded"
    fi
    sleep 10
    log "rule proposal"
    $REGULATORYD tx regulatory \
        submit-rule-proposal ./source/proposal/rule.json \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Submit rule proposal  failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Submit rule proposal succeeded"
    fi
    sleep 10
    $REGULATORYD tx gov vote 2 yes \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Vote rule proposal failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Vote rule proposal succeeded"
    fi
    sleep 10
    log "binding proposal"

    $REGULATORYD tx regulatory \
        submit-binding-proposal ./source/proposal/binding.json \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Submit binding proposal  failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Submit binding proposal succeeded"
    fi
    sleep 10
    $REGULATORYD tx gov vote 3 yes \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Vote binding proposal failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Vote binding proposal succeeded"
    fi
    sleep 10
    log "relation proposal"
    $REGULATORYD tx regulatory \
        submit-relation-proposal ./source/proposal/relation.json \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Submit relation proposal  failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Submit relation proposal succeeded"
    fi
    sleep 10
    $REGULATORYD tx gov vote 4 yes \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Vote relation proposal failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Vote relation proposal succeeded"
    fi
    sleep 10
    log "reward list proposal"
    $REGULATORYD tx regulatory \
        submit-reward-list-proposal ./source/proposal/rewardlist.json \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Submit reward list proposal  failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Submit reward list proposal succeeded"
    fi
    sleep 10
    $REGULATORYD tx gov vote 5 yes \
        --from $($REGULATORYD --home ./data/ibc-2 keys show validator --keyring-backend test -a) \
        --keyring-backend test \
        --gas auto \
        --chain-id $CHAIN_ID \
        --home ./data/ibc-2 \
        --yes > /dev/null
    if [ $? -ne 0 ]; then
        log $LOG_ERROR "Vote reward list proposal failed ,stop !"
         exit 1
    else
        log $LOG_SUCCESS "Vote reward list proposal succeeded"
    fi
    log "proposal results"
    $REGULATORYD query gov proposals
}
deploy_ipfs_postgresql(){
    cd $graphnode_data_path
    log "deploy ipfs and postgresql docker  container..."
    docker-compose up -d 
    sleep 5
    ipfs_container_status=$(docker inspect --format '{{.State.Status}}' graphnode-data-ipfs-1)

    if [ "$ipfs_container_status" = "running" ]; then
        log $LOG_SUCCESS  "Ipfs container is running"
    else
        log $LOG_WARNING "Ipfs container is not running"
        exit  
    fi
    postgres_container_status=$(docker inspect --format '{{.State.Status}}' graphnode-data-postgres-1)
    if [ "$postgres_container_status" = "running" ]; then
        log $LOG_SUCCESS  "postgres container is running"
    else
        log $LOG_WARNING "postgres container is not running"
        exit  
    fi
    cd $deploy_home_dir
}

deploy_graph_node(){
    
    log "Start ipfs and postgresql."
    deploy_ipfs_postgresql
    sleep 10

    # Start a new tmux session and detach it
    tmux new-session -d -s graphnode

    # Create a new window and run a command in it
    tmux new-window -t graphnode:1 -n "graphnode window"
    tmux send-keys -t graphnode:1 "graph-node --config ./source/graphnode/config.toml --ipfs 127.0.0.1:5001 --node-id index_node_cosmos_1 &> ./logs/graphnode.log " C-m

    log $LOG_SUCCESS "View the background graphnode node program through the command ' tmux attach -t graphnode-node '"
    log $LOG_SUCCESS "Graphnode deploy success! "

    sleep 10

    log "Start deploy manifest"

    deploy_system_manifest

    
}
deploy_system_manifest(){
    log "deploy system graph node manifest"
    cd "../manifests/system-manifest"
    yarn 
    yarn codegen && yarn build
    yarn create-local && yarn deploy-local
    cd $deploy_home_dir

    cd "../manifests/ics721-manifest"
    yarn 
    yarn codegen && yarn build
    yarn create-local && yarn deploy-local
    cd $deploy_home_dir
    log "End of deployment manifest. Please review the information to determine whether the deployment succeeded or failed "


}



dir_must_exists() {
    if [ ! -d "$1" ]; then
        exit_with_clean "$1 DIR does not exist, please check!"
    fi
}
exit_with_clean()
{
    local content=${1}
    echo -e "\033[31m[ERROR] ${content}\033[0m"
    if [ -d "${chain_deploy_path}" ];then
       rm -rf ${chain_deploy_path}
    fi
    exit 1
}
log() {
  # Define colors
  local RESET='\033[0m'
  local RED='\033[0;31m'
  local GREEN='\033[0;32m'
  local YELLOW='\033[0;33m'
  
  # Get the current date and time
  local timestamp=$(date +"%Y-%m-%d %H:%M:%S")
  # Get the current line number
  local line_num=${BASH_LINENO[0]}

  # Print the log message with timestamp and line number, and color based on the log level
  if [ "$1" == "ERROR" ]; then
    echo -e "${RED}$timestamp [Line $line_num] [ERROR]: $2${RESET}"
  elif [ "$1" == "WARNING" ]; then
    echo -e "${YELLOW}$timestamp [Line $line_num] [WARNING]: $2${RESET}"
  elif [ "$1" == "SUCCESS" ]; then
    echo -e "${GREEN}$timestamp [Line $line_num] [SUCCESS]: $2${RESET}"
  else
    echo -e "${RESET}$timestamp [Line $line_num]: $1${RESET}"
  fi
}
version_compare() {
    # Remove 'v' character from version strings
    local ver1="${1#v}"
    local ver2="${2#v}"
    
    if dpkg --compare-versions "$ver1" ge "$ver2"; then
        return 0
    else
        return 1
    fi
}
parse_params $@
