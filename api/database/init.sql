CREATE TABLE tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  description VARCHAR(250) NOT NULL,
  completed BOOLEAN NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL
);
