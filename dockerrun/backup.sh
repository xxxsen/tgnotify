#!/bin/bash

docker exec tg_notify_db mariadb-dump --skip-extended-insert tgmessager -usender -psender > bak.sql
