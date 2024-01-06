import type { Project } from "$lib/types/project.js";

export async function load({ params, fetch }): Promise<Project> {
	const response = await fetch(`/api/project/${params.project_id}`);
	const responseJSON: Project = await response.json();

	return responseJSON;
}
