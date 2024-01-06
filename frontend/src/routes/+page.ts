import { type Projects } from "$lib/types/projects";

export async function load({ fetch }): Promise<Projects> {
	const projects = await fetch(`/api/projects`);
	return await projects.json();
}
