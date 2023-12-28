import { type Project } from "$lib/types/project";

export function load({ params, fetch }): Promise<Project> {
	return fetch(`/api/project/${params.project_id}`).then((resp) => resp.json());
}
