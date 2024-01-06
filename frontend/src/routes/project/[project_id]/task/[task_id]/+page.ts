import type { ClassificationTaskLabel } from "$lib/types/label";
import type { Project } from "$lib/types/project.js";
import type { ClassificationTask } from "$lib/types/task.js";

export interface PageProps {
	label_endpoint: string;
	labels: ClassificationTaskLabel[];
	task: ClassificationTask;
}

export async function load({ params, fetch }): Promise<PageProps> {
	const taskResponse = await fetch(
		`/api/project/${params.project_id}/classification_task/${params.task_id}`,
	);
	const task: ClassificationTask = await taskResponse.json();

	const projectResponse = await fetch(`/api/project/${params.project_id}`);
	const project: Project = await projectResponse.json();

	return {
		label_endpoint: `/api/project/${params.project_id}/classification_task/${params.task_id}/label`,
		labels: project.labels,
		task: task,
	};
}
