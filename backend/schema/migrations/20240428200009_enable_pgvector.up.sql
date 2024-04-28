CREATE EXTENSION vector;

ALTER TABLE
	pending_classification_tasks DROP COLUMN embedding,
ADD
	COLUMN embedding vector(4096) NOT NULL;

ALTER TABLE
	classification_tasks DROP COLUMN embedding,
ADD
	COLUMN embedding vector(4096) NOT NULL;
