import type { ClassificationTaskLabel } from "./label";
import type { ClassificationTask } from "./task";

export interface Project {
	id: string;
	name: string;
	description: string;
	created_at: string;
	labels: ClassificationTaskLabel[];
	classification_tasks: ClassificationTask[];
}
