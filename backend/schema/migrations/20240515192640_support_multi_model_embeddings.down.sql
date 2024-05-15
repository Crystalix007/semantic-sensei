-- Add the embedding column back to the old tables
ALTER TABLE
	classification_tasks
ADD
	COLUMN embedding vector(4096);

ALTER TABLE
	pending_classification_tasks
ADD
	COLUMN embedding vector(4096);

-- Copy the embeddings from the new table back to the old tables
UPDATE
	classification_tasks ct
	JOIN embeddings e ON ct.llm_output_sha256 = e.sha256
	AND e.model = 'default'
SET
	ct.embedding = e.embedding;

UPDATE
	pending_classification_tasks pct
	JOIN embeddings e ON pct.llm_output_sha256 = e.sha256
	AND e.model = 'default'
SET
	pct.embedding = e.embedding;

-- Delete the embeddings from the new table
DROP TABLE embeddings;
