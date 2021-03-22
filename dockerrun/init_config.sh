#!/bin/bash

. shell_args.sh

#============DONT CHANGE IT================

# create data folder
echo "create DB_PATH folder:$DB_PATH"
echo "create DB_INIT folder:$DB_INIT"
echo "create APP_CONFIG folder:$APP_CONFIG"
mkdir $DB_PATH $DB_INIT $APP_CONFIG -p
echo "create folder succ..."

# create .env 
echo "create .env file..."
echo "CONFIG_PATH=$APP_CONFIG" > .env
echo "DB_INIT=$DB_INIT" >> .env
echo "DB_PATH=$DB_PATH" >> .env
echo "DB_PORT=$DB_PORT" >> .env
echo "DB_USER=$DB_USER" >> .env
echo "DB_PWD=$DB_PWD" >> .env
echo "create .env file succ..."

# create init sql
echo "create db init sql..."
echo 'create database if not exists `tgmessager`;' > tgmessager.sql;
echo 'use tgmessager;' >> tgmessager.sql;
echo 'CREATE TABLE `tbl_tgnotify` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user` varchar(32) NOT NULL,
  `code` varchar(32) NOT NULL,
  `chatid` bigint(20) unsigned NOT NULL,
  `ts` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_chatid` (`chatid`),
  UNIQUE KEY `uniq_uinfo` (`user`,`code`,`chatid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;' >> tgmessager.sql

mv tgmessager.sql $DB_INIT
echo "move db init sql to $DB_INIT succ..."

# create app config
echo "create app config..."
cp config.tplt config.json
sed -i "s/#TOKEN_PLACE_HOLDER#/${BOT_TOKEN}/g" config.json
sed -i "s/#DB_PORT#/${DB_PORT}/g" config.json
sed -i "s/#DB_USER#/${DB_USER}/g" config.json
sed -i "s/#DB_PWD#/${DB_PWD}/g" config.json
mv config.json $APP_CONFIG
echo "move app config to $APP_CONFIG succ..."
