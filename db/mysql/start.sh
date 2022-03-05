#!/bin/bash
set -e
service mysql start
mysql < /db/bolg-service.sql
service mysql.stop