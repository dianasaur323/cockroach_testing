# CockroachDB Testing

Goal: Make it easier to set up a database and run load against it

Functions:
1) Take in a given schema and set up a database
2) Take a selection of queries and run them against the database

Considerations:
- Allow some sort of connection to different database configurations
- Allow configurable concurrent connections

Flow:
- Pipe arguments and queries in from a config file
- Parse those arguments and queries
