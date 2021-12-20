#!/usr/bin/env bash
#
# Copyright (C) BABEC. All rights reserved.
# Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#set -x
# 将BRANCH_NEW的数据合并到BRANCH
BRANCH=develop
BRANCH_NEW=v2.2.0_alpha_qc
PRE_LOG="【log】 "

projects=("pb-go"
  "common"
  "protocol"
  "logger"
  "utils"
  "localconf"
  "chainconf"
  "txpool-single"
  "txpool-batch"
  "store-leveldb"
  "store-badgerdb"
  "store-sqldb"
  "store"
  "vm-native"
  "vm"
  "vm-wasmer"
  "vm-gasm"
  "vm-wxvm"
  "vm-evm"
  "vm-docker-go"
  "raftwal"
  "libp2p-core"
  "net-common"
  "libp2p-pubsub"
  "net-libp2p"
  "net-liquid"
  "consensus-utils"
  "consensus-dpos"
  "consensus-maxbft"
  "consensus-raft"
  "consensus-solo"
  "consensus-tbft"
  "chainmaker-go"
)

cd ../..
mkdir -p chainmaker
cd chainmaker
for project in ${projects[*]}; do
  #如果文件夹不存在，则克隆
  if [ ! -d ${project} ]; then
    git clone git@git.code.tencent.com:ChainMaker/${project}.git
  fi

  cd $project
  echo "$PRE_LOG $(pwd)"
  echo "$PRE_LOG cd $project && git stash && git checkout $BRANCH && git pull origin $BRANCH_NEW"
  git fetch --all
  if git rev-parse --verify remotes/origin/${BRANCH_NEW}; then
    echo "$PRE_LOG $project find branch remotes/origin/${BRANCH_NEW}"
  else
    echo "$PRE_LOG do nothing for $project not found branch $BRANCH_NEW "
    cd ..
    continue
  fi

  git stash
  git checkout $BRANCH
  git reset --hard
  git pull
  git pull origin $BRANCH_NEW
  sed -i "s%VERSION=.*%VERSION=${BRANCH}%g" Makefile
  make gomod

  git status
  git add Makefile go.mod go.sum
  git commit -m "auto script/gomod_merge.sh from ${BRANCH_NEW} to ${BRANCH}"
  git push origin $BRANCH
  cd ..
  echo "$PRE_LOG merge $project from ${BRANCH_NEW} to ${BRANCH} finish"
  echo ""
done

echo "$PRE_LOG job done"