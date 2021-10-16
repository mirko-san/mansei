#!/usr/bin/env bash

#run the setup script to create the DB and the schema in the DB
mysql -uroot -proot test < "/docker-entrypoint-initdb.d/createTestDatabase.sql"
