import type { PaginatedData } from "$lib/types/pagination";
import { type Project } from "$lib/types/project";
import type { ClassificationTask } from "$lib/types/task.js";

export interface PageData {
	project: Project;
	pending_tasks: PaginatedData<ClassificationTask>;
	completed_tasks: PaginatedData<ClassificationTask>;
}

export async function load({ params, fetch }): Promise<PageData> {
	const projectResponse = await fetch(`/api/project/${params.project_id}`);
	const project = await projectResponse.json();

	const pendingTasksResponse = await fetch(
		`/api/project/${params.project_id}/pending_classification_tasks`,
	);
	const pendingTasks = await pendingTasksResponse.json();

	const completedTasksResponse = await fetch(
		`/api/project/${params.project_id}/classification_tasks`,
	);
	const completedTasks = await completedTasksResponse.json();

	return {
		project,
		pending_tasks: pendingTasks,
		completed_tasks: completedTasks,
	};
}
