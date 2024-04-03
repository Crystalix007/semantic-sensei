import type { ClassificationTaskLabel } from "$lib/types/label";
import type { PaginatedData } from "$lib/types/pagination";
import type { Project } from "$lib/types/project.js";
import type { ClassificationTask } from "$lib/types/task.js";
import { redirect } from "@sveltejs/kit";
import type { PageLoadEvent } from "./$types";

export interface PageProps {
	label_endpoint: string;
	labels: ClassificationTaskLabel[];
	task: ClassificationTask;
}

export async function load({
	params,
	fetch,
}: PageLoadEvent): Promise<PageProps> {
	const taskResponse = await fetch(
		`/api/project/${params.project_id}/pending_classification_tasks`,
	);
	const tasks: PaginatedData<ClassificationTask> = await taskResponse.json();

	if (tasks.total === 0) {
		redirect(302, `/project/${params.project_id}`);
	}

	const projectResponse = await fetch(`/api/project/${params.project_id}`);
	const project: Project = await projectResponse.json();

	return {
		label_endpoint: `/api/project/${params.project_id}/classification_task/${tasks.data[0].id}/label`,
		labels: project.labels,
		task: tasks.data[0],
	};
}
