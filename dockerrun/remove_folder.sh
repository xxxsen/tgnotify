#!/bin/bash

. shell_args.sh

if [ "$DB_PATH" != "" ]; then
    echo "remove folder:$DB_PATH"
    sudo rm "$DB_PATH" -rf
fi	
if [ "$DB_INIT" != "" ]; then
    echo "remove folder:$DB_INIT"	
    sudo rm "$DB_INIT" -rf
fi 	
if [ "$APP_CONFIG" != "" ]; then
    echo "remove folder:$APP_CONFIG"
    sudo rm "$APP_CONFIG" -rf
fi 	

echo "clean .env file..."
rm .env -f

