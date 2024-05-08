-- Add pgcrypto extension for the `digest` function.
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Clear out duplicate pending classification tasks.
DELETE FROM
	pending_classification_tasks pct1 using pending_classification_tasks pct2
WHERE
	pct1.project_id = pct2.project_id
	AND pct1.llm_input = pct2.llm_input
	AND pct1.llm_output = pct2.llm_output
	AND pct1.id > pct2.id;

ALTER TABLE
	pending_classification_tasks
ADD
	COLUMN llm_input_sha256 bytea GENERATED ALWAYS AS (DIGEST(llm_input, 'sha256')) STORED,
ADD
	COLUMN llm_output_sha256 bytea GENERATED ALWAYS AS (DIGEST(llm_output, 'sha256')) STORED,
ADD
	CONSTRAINT pending_classification_tasks_unique UNIQUE (
		project_id,
		llm_input_sha256,
		llm_output_sha256
	);

DELETE FROM
	classification_tasks ct1 USING classification_tasks ct2
WHERE
	ct1.project_id = ct2.project_id
	AND ct1.llm_input = ct2.llm_input
	AND ct1.llm_output = ct2.llm_output
	AND ct1.id > ct2.id;

ALTER TABLE
	classification_tasks
ADD
	COLUMN llm_input_sha256 bytea GENERATED ALWAYS AS (DIGEST(llm_input, 'sha256')) STORED,
ADD
	COLUMN llm_output_sha256 bytea GENERATED ALWAYS AS (DIGEST(llm_output, 'sha256')) STORED,
ADD
	CONSTRAINT classification_tasks_unique UNIQUE (
		project_id,
		llm_input_sha256,
		llm_output_sha256
	);
