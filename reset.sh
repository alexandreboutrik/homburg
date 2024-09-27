#!/bin/bash

function Try() {
  color=${1}
  shift
  char=${1}
  shift

  psql -U postgres -d postgres -p 9900 -c "${*}" 1>>/dev/null 2>&1
  code=$?
  if [ ${code} -ne 0 ]; then
    echo -e "\e[${color};1m[${char}] (code: ${code}) ${*}\e[0m"
  else
    echo -e "[+] ${*}"
  fi
}

curl -s -L "http://localhost:9900"
if [ $? -ne 52 ]; then
  echo -e "\e[31;1m[-] postgresql not running\e[0m"
  exit
fi

pwd=$(echo $RANDOM | md5sum | head -c 32)
echo -e "[\e[33;1m!\e[0m] new HAT_DB_PWD: ${pwd}"

Try 33 "!" "drop database hat_db;"
Try 33 "!" "drop user hat_user;"

Try 31 "-" "create database hat_db;"
Try 31 "-" "create user hat_user with encrypted password '${pwd}';"
Try 31 "-" "grant all privileges on database hat_db to hat_user;"
Try 31 "-" "alter database hat_db owner to hat_user;"

rm -f sqlite.db

go run migrate/migrate.go
if [ $? -eq 0 ]; then
  echo -e "[+] go run migrate/migrate.go"
fi
