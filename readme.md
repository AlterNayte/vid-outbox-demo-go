# Sample Code for the Transactional Outbox Pattern

This repository contains a sample implementation of the Transactional Outbox Pattern. The pattern is used to ensure that messages are sent in a transactionally consistent manner. This is important when you need to ensure that a message is sent if and only if a transaction is committed.

The sample code is written in Go and uses Postgres as the database. 

## Running the Sample Code

To bring up the api, database, redpanda and debezium, run the following command:
```bash
docker-compose up -d
```