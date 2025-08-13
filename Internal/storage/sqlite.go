package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const ddl = `
CREATE TABLE IF NOT EXISTS request_logs (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  ts TEXT NOT NULL,
  method TEXT NOT NULL,
  path TEXT NOT NULL,
  status INTEGER NOT NULL,
  duration_ms INTEGER NOT NULL,
  username TEXT,
  ip TEXT,
  user_agent TEXT,
  request_id TEXT
);
CREATE INDEX IF NOT EXISTS idx_request_logs_ts ON request_logs(ts);
CREATE INDEX IF NOT EXISTS idx_request_logs_path ON request_logs(path);
`

func OpenLogsDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(ddl); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

func InsertRequestLog(db *sql.DB, ts, method, path string, status, durMs int, username, ip, ua, reqID string) error {
	_, err := db.Exec(
		`INSERT INTO request_logs (ts, method, path, status, duration_ms, username, ip, user_agent, request_id)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		ts, method, path, status, durMs, username, ip, ua, reqID,
	)
	return err
}
