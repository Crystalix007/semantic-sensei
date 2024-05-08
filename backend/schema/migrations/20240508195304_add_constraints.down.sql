ALTER TABLE
	`pending_classification_tasks` DROP CONSTRAINT `unique_project_llm_input_llm_output`;

ALTER TABLE
	`classification_tasks` DROP CONSTRAINT `unique_project_llm_input_llm_output`;
