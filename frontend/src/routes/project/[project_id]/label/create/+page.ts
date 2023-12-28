export interface PageProps {
	form_url: string;
}

export function load({ params }): PageProps {
	return {
		form_url: `/api/project/${params.project_id}/classification_task_label`,
	};
}
