# CockroachDB Testing

Goal: Allow user to provide schema and queries, and automatically run load based
on inputted variables.

Flow:
1) Wipe database
2) Take in a given schema and set up a database
3) Run load against database

Configurations:
- Concurrent connections
- Load running on multiple nodes
