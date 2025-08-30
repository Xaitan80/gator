# gator CLI

gator is a command-line tool to manage RSS feeds and aggregate posts into a PostgreSQL database. It allows you to register feeds, follow them, and continuously fetch and store posts.

---

## Prerequisites

Before using Gator, make sure you have the following installed:

- **Go** (1.20 or higher recommended)
- **PostgreSQL**

### Installing on macOS with Homebrew

If you don’t already have Go and PostgreSQL, you can install them with [Homebrew](https://brew.sh/):

```bash
# Install Go
brew install go

# Verify Go installation
go version

# Install PostgreSQL
brew install postgresql

# Start PostgreSQL service
brew services start postgresql

# Verify PostgreSQL installation
psql --version
