CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS projects (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pending_classification_tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id INTEGER NOT NULL,
	llm_input TEXT NOT NULL,
	llm_input_sha256 BYTEA GENERATED ALWAYS AS (DIGEST(llm_input, 'sha256')) STORED,
	llm_output TEXT NOT NULL,
	llm_output_sha256 BYTEA GENERATED ALWAYS AS (DIGEST(llm_output, 'sha256')) STORED,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (project_id) REFERENCES projects (id),
	CONSTRAINT pending_classification_tasks_unique UNIQUE (project_id, llm_input_sha256, llm_output_sha256)
);

CREATE TABLE IF NOT EXISTS classification_tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id INTEGER NOT NULL,
	llm_input TEXT NOT NULL,
	llm_input_sha256 BYTEA GENERATED ALWAYS AS (DIGEST(llm_input, 'sha256')) STORED,
	llm_output TEXT NOT NULL,
	llm_output_sha256 BYTEA GENERATED ALWAYS AS (DIGEST(llm_output, 'sha256')) STORED,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	label_id INTEGER NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects (id),
	FOREIGN KEY (label_id) REFERENCES classification_task_labels (id),
	CONSTRAINT classification_tasks_unique UNIQUE (project_id, llm_input_sha256, llm_output_sha256)
);

CREATE TABLE IF NOT EXISTS embeddings (
	sha256 BYTEA NOT NULL,
	model TEXT NOT NULL,
	embedding vector(4096) NOT NULL,
	PRIMARY KEY (sha256, model)
);

CREATE TABLE IF NOT EXISTS classification_task_labels (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id INTEGER NOT NULL,
	label TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (project_id) REFERENCES projects (id),
	UNIQUE (project_id, label)
);
