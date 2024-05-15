import type { Embeddings } from "./embeddings";

export interface ClassificationTask {
	id: number;
	project_id: number;
	llm_input: string;
	llm_output: string;
	embedding: Embeddings;
	created_at: string;
	label_id?: number;
}
