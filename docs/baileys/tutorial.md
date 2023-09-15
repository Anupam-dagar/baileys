# Polls App

To give a more indepth overview of how baileys work, we will be building a simple polls app which allows users to create
polls and vote on it and view all the created polls. All the essential components of baileys will be elaborated while
building the app.

## Setup a polls baileys project
1. Create a new directory named `polls`
2. Init a go module
```bash
$ go mod init polls
```
3. Add baileys as a dependency
```bash
$ go get -u github.com/Anupam-dagar/baileys
```

## Database Setup
1. Create a new postgres database named `polls`.
```sql
CREATE DATABASE IF NOT EXISTS polls;
```
2. Create the required tables.
   - `polls` -  Poll created by the user.
```sql
CREATE TABLE IF NOT EXISTS polls (
        id         varchar(255) NOT NULL PRIMARY KEY,
        title      varchar(255) NOT NULL,
        created_at timestamptz  NOT NULL DEFAULT NOW(),
        updated_at timestamptz  NOT NULL DEFAULT NOW(),
        deleted_at timestamptz NULL,
        created_by varchar(255) NOT NULL,
        updated_by varchar(255) NOT NULL,
        deleted_by varchar(255) NULL
);
```
    - `poll_options` -  Options for the poll created by the user.
```sql
CREATE TABLE IF NOT EXISTS poll_options (
        id         varchar(255) NOT NULL PRIMARY KEY,
        poll_id    varchar(255) NOT NULL,
        title      varchar(255) NOT NULL,
        created_at timestamptz  NOT NULL DEFAULT NOW(),
        updated_at timestamptz  NOT NULL DEFAULT NOW(),
        deleted_at timestamptz NULL,
        created_by varchar(255) NOT NULL,
        updated_by varchar(255) NOT NULL,
        deleted_by varchar(255) NULL
);
```
    - `votes` -  Votes received to an option for a poll.
```sql
CREATE TABLE IF NOT EXISTS votes (
        id             varchar(255) NOT NULL PRIMARY KEY,
        poll_id        varchar(255) NOT NULL,
        poll_option_id varchar(255) NOT NULL,
        created_at     timestamptz  NOT NULL DEFAULT NOW(),
        updated_at     timestamptz  NOT NULL DEFAULT NOW(),
        deleted_at     timestamptz NULL,
        created_by     varchar(255) NOT NULL,
        updated_by     varchar(255) NOT NULL,
        deleted_by     varchar(255) NULL
);
```

## Directory Structure
You can use the following commands to setup the directory structure.
```bash
mkdir config
mkdir controller
mkdir entity
mkdir service
mkdir route
mkdir repository
touch config/dev.yaml
touch main.go
```
