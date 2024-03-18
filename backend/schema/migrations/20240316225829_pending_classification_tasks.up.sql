CREATE TABLE pending_classification_tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id INTEGER NOT NULL,
	llm_input TEXT NOT NULL,
	llm_output TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	embedding BLOB NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects (id)
);

INSERT INTO pending_classification_tasks (
	project_id,
	llm_input,
	llm_output,
	created_at,
	embedding
)
SELECT
	project_id,
	llm_input,
	llm_output,
	created_at,
	embedding
FROM classification_tasks
WHERE label_id IS NULL;

DELETE FROM classification_tasks
WHERE label_id IS NULL;

CREATE TABLE classification_tasks_new (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id INTEGER NOT NULL,
	llm_input TEXT NOT NULL,
	llm_output TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	embedding BLOB NOT NULL,
	label_id INTEGER NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects (id),
	FOREIGN KEY (label_id) REFERENCES classification_task_labels (id)
);

INSERT INTO classification_tasks_new (
	project_id,
	llm_input,
	llm_output,
	created_at,
	embedding,
	label_id
) SELECT
	project_id,
	llm_input,
	llm_output,
	created_at,
	embedding,
	label_id
FROM classification_tasks;

DROP TABLE classification_tasks;

ALTER TABLE classification_tasks_new RENAME TO classification_tasks;
