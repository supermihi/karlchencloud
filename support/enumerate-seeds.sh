#!/usr/bin/env bash

KC_INVITE_CODE="1234"
KC_TABLE_ID="5678"
KC_ROOT="$(readlink -f "$(git rev-parse --show-toplevel)")"

running_pids=()

function launch_server() {
    CONSTANT_INVITE_CODE="${KC_INVITE_CODE}" \
    CONSTANT_TABLE_ID="${KC_TABLE_ID}" \
    KC_DBG_SEED="${1}"\
    \
    go run cmd/server/main.go 2>&1 &
    running_pids+=( "${!}" )
}

function launch_bots() {
    INVITE_CODE="${KC_INVITE_CODE}" \
    TABLE_ID="${KC_TABLE_ID}" \
    INIT_TABLE=1 \
    NUM_BOTS=4 \
    \
    go run cmd/bot_client/main.go 2>&1 &
    running_pids+=( "${!}" )
}

function kill_running() {
    for pid in "${running_pids[@]}";
    do
        echo "Killing ${pid}"
        pkill -P "${pid}"
    done
    running_pids=()
}

function run_test_seed() {
    launch_server "${1}"
    launch_bots
    sleep 2
    kill_running
}

function test_seed() {
    run_test_seed "${1}" > common.log

    echo "Evaluating seed ${1}"
    [ "$(cat common.log | grep ' plays ' | wc -l)" == 48 ] || exit 1
}

function main() {
    cd "${KC_ROOT}"
    for i in $(seq 1 1000);
    do
        test_seed "${RANDOM}"
    done
}

main
