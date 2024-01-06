export interface ClassificationTask {
	id: number;
	project_id: number;
	llm_input: string;
	llm_output: string;
	embedding: string;
	created_at: string;
	label_id?: string;
}
