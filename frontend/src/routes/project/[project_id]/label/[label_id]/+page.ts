import type { ClassificationTaskLabel } from "$lib/types/label.js";

export function load({ params, fetch }): Promise<ClassificationTaskLabel> {
	return fetch(
		`/api/project/${params.project_id}/classification_task_label/${params.label_id}`,
	).then((resp) => resp.json());
}
