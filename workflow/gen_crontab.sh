#!/usr/bin/env bash

WORKFLOW_DIR=$(realpath "${BASH_SOURCE%/*}")
OUTPUT_FILE=${WORKFLOW_DIR}/update.log
REC_DIR=${WORKFLOW_DIR}/rec

while getopts "w:o:r:g:" opt; do
    case $opt in
        w)
            WORKFLOW_DIR="$OPTARG"
            ;;
        o)
            OUTPUT_FILE="$OPTARG"
            ;;
        r)
            REC_DIR="$OPTARG"
            ;;
        g)
            GENTOO_PREFIX_DIR="$OPTARG"
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            exit 1
            ;;
    esac
done

shift $((OPTIND-1))

if [ -f "$GENTOO_PREFIX_DIR/usr/bin/emerge" ]; then
    GENTOO_ENV="GENTOO_PREFIX=${GENTOO_PREFIX_DIR}"
fi


sed -e "s|@WORKFLOW_DIR@|${WORKFLOW_DIR}|g; 
s|@OUTPUT_FILE@|${OUTPUT_FILE}|g;
s|@ENV@|APP_BIN=${APP_BIN} CFG_FILE=${CFG_FILE} STORAGE_DIR=${STORAGE_DIR} ${GENTOO_ENV}|g;
s|@REC_DIR@|${REC_DIR}|g;" "${WORKFLOW_DIR}/update.crontab"
