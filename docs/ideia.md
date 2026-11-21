# GoDB

GoDB is a hobby project to explore how databases work by building a reliable database in Go. It is intended to run without cgo or compatibility layers for other languages, using the Go runtime as its foundation.

## Fast reads

The idea is to keep frequently accessed data in memory, drawing inspiration from Redis and Memcached. The choice of structures and the amount of data kept in memory remain open. Their effect on read speed and memory use will need to be measured against reading from persistent storage.

## Transactional writes

GoDB aims to support ACID transactions and persistent storage, taking inspiration from databases such as PostgreSQL, SQLite, and MySQL. The initial idea is to persist writes and keep the in-memory state consistent with committed transactions.

Writing to disk and updating memory do not by themselves provide ACID guarantees. The design must define how transactions commit, how concurrent operations are isolated, and how the database recovers after a failure. It must also specify when a write is durable and when its changes become visible to readers.

One option to investigate is a write-ahead log; a rollback journal is another. Their tradeoffs and suitability for GoDB remain to be evaluated. For background, see PostgreSQL's [introduction to write-ahead logging](https://www.postgresql.org/docs/current/wal-intro.html) and SQLite's explanation of [atomic commit and recovery](https://www.sqlite.org/atomiccommit.html).

## Status and validation

These are goals, not implemented guarantees. GoDB does not yet have a database implementation. Once there is one, tests should check transaction behavior under concurrent access and recovery after interruptions at different stages of a write. Read performance should be compared with and without the in-memory structures using representative workloads.
