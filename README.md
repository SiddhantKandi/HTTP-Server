📰 RSS Aggregator API (Go)

A fully-featured RSS feed aggregator and RESTful API written in Go. This project allows users to subscribe to RSS feeds, view posts from their followed feeds, and manage user accounts — all backed by PostgreSQL and structured with SQLC, Goose, and idiomatic Go practices.

📌 Features

✅ Add, follow, and unfollow RSS feeds
📥 Fetch and store latest posts from followed feeds
👤 Create and authenticate users with API keys
🧩 RESTful API with secure endpoints for all core actions
🔐 Authentication middleware with API key-based access
🛠 SQL migrations with Goose and query generation with SQLC
📂 Project Structure

├── cmd/                # Entry point for the application
├── internal/
│   └── database/       # Generated Go code from SQLC
├── sql/
│   ├── schema/         # Goose migration files
│   └── queries/        # SQL queries for SQLC
├── main.go
└── sqlc.yaml           # SQLC configuration
🧑‍💻 Getting Started

1. Prerequisites
Go 1.20+
PostgreSQL
SQLC
Goose
Install tools:

go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
2. Configure Environment
Create a .env file with your PostgreSQL connection string:

DATABASE_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable
🧱 Database

Migrations
Run database migrations with Goose:

goose -dir sql/schema postgres $DATABASE_URL up
SQLC Code Generation
Generate Go code for your queries:

sqlc generate
🚀 API Endpoints

🔐 Authentication
POST /v1/users — Create a new user
GET /v1/users — Get authenticated user info
Header: Authorization: ApiKey <key>
📡 Feeds
POST /v1/feeds — Create a new feed (Authenticated)
GET /v1/feeds — List all available feeds
When a feed is created, the creator is automatically set to follow it.

🧾 Posts
GET /v1/posts — Get latest posts from feeds a user follows
Header: Authorization: ApiKey <key>
Query param: limit=int (optional)
Posts are saved with the following attributes:

id, title, url (unique), description, published_at, feed_id, created_at, updated_at
📥 Feed Follows
POST /v1/feed_follows — Follow a feed (Authenticated)
GET /v1/feed_follows — List followed feeds (Authenticated)
DELETE /v1/feed_follows/{feedFollowID} — Unfollow a feed (Authenticated)
⚙️ Implementation Highlights

🧵 UUIDs used for all table IDs
🕓 Timestamp fields: created_at, updated_at
🗃 Posts are fetched and stored in the DB using a scraper
⚠️ Duplicate posts (same URL) are skipped
🧠 "Published at" timestamps are properly parsed and normalized
🔑 Secure, unique 256-bit API keys generated with:
encode(sha256(random()::text::bytea), 'hex')
📌 Tech Stack

Go – web server & logic
PostgreSQL – relational database
SQLC – compile-time SQL query generation
Goose – database schema migrations
Chi – lightweight router
uuid – unique identifier generation
dotenv – configuration via environment variables
🛠 Developer Tips

Use sqlc.yaml to manage schema + queries in sql/
Use internal/database for all DB access logic
Use context-based or custom handler-based middleware for auth
Keep error handling robust: log duplicates, but fail on real errors
🧪 Sample Requests

Create User

POST /v1/users
{
  "name": "Lane"
}
Get Posts (Authenticated)

GET /v1/posts?limit=20
Authorization: ApiKey <your_api_key>
Add Feed

POST /v1/feeds
Authorization: ApiKey <your_api_key>
{
  "name": "Boot.dev Blog",
  "url": "https://blog.boot.dev/index.xml"
}
