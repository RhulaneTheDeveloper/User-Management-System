User Management Project

This project is a User Management System built with Go (Golang) and the Gin framework, with support for both Percona MySQL and Percona PostgreSQL databases. It exposes RESTful APIs for basic CRUD operations on a users table.
Features

Dual database support: MySQL and PostgreSQL

User registration and authentication

CRUD operations via RESTful API (Gin)

Configurable via environment variables

Built and tested on both WSL (Ubuntu) and cursor
Tech Stack

Backend: Go 1.25.1

Framework: Gin Gonic

Databases:

Percona MySQL (go_dev)

Percona PostgreSQL (go_dev)
What I Did

Set up two database servers (Percona MySQL & PostgreSQL) inside WSL.

Created a go_dev database in both servers.

Installed Go 1.25.1 and set up a project under ~/dev/user-management.

Configured and tested two database connections (MySQL & PostgreSQL).

Designed a users table with authentication fields.

Started building a RESTful API using the Gin framework with CRUD endpoints.

Planned team sessions to walk through the database schema and API design.
