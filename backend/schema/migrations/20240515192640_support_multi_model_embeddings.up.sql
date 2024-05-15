-- Create embeddings table
CREATE TABLE embeddings (
	sha256 BYTEA NOT NULL,
	model TEXT NOT NULL,
	embedding vector(4096) NOT NULL,
	PRIMARY KEY (sha256, model)
);

-- Copy embeddings from classification_tasks to the new table
INSERT INTO
	embeddings (sha256, model, embedding)
SELECT
	llm_output_sha256,
	'default',
	embedding
FROM
	classification_tasks;

INSERT INTO
	embeddings (sha256, model, embedding)
SELECT
	llm_output_sha256,
	'default',
	embedding
FROM
	pending_classification_tasks ON CONFLICT DO NOTHING;

-- Delete embeddings from the old tables
ALTER TABLE
	classification_tasks DROP COLUMN embedding;

ALTER TABLE
	pending_classification_tasks DROP COLUMN embedding;
