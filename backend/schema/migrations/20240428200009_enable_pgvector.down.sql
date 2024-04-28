DROP EXTENSION vector;

ALTER TABLE
	pending_classification_tasks DROP COLUMN embedding,
ADD
	COLUMN embedding BYTEA NOT NULL;

ALTER TABLE
	classification_tasks DROP COLUMN embedding,
ADD
	COLUMN embedding BYTEA NOT NULL;
