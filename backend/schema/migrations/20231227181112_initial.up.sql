CREATE TABLE IF NOT EXISTS projects (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS classification_task_labels (
	id SERIAL PRIMARY KEY,
	project_id INTEGER NOT NULL,
	label TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (project_id) REFERENCES projects (id),
	UNIQUE (project_id, label)
);

CREATE TABLE IF NOT EXISTS classification_tasks (
	id SERIAL PRIMARY KEY,
	project_id INTEGER NOT NULL,
	llm_input TEXT NOT NULL,
	llm_output TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	embedding BYTEA NOT NULL,
	label_id INTEGER NULL,
	FOREIGN KEY (project_id) REFERENCES projects (id),
	FOREIGN KEY (label_id) REFERENCES classification_task_labels (id)
);
