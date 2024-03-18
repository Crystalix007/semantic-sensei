CREATE TABLE classification_tasks_new (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id INTEGER NOT NULL,
	llm_input TEXT NOT NULL,
	llm_output TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	embedding BLOB NOT NULL,
	label_id INTEGER NULL,
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

INSERT INTO classification_tasks (
		project_id,
		llm_input,
		llm_output,
		created_at,
		embedding,
		label_id
)
SELECT (
    project_id,
    llm_input,
    llm_output,
    created_at,
    embedding,
    NULL
)
FROM pending_classification_tasks;

DROP TABLE pending_classification_tasks;
