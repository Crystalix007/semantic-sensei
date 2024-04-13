export async function handleFetch({ request, fetch }): Promise<Response> {
	const url = new URL(request.url);

	if (url.hostname.endsWith(".localdomain")) {
		url.protocol = "http:";
		url.hostname = "backend";
	}

	console.log("fetching", url.toString());

	return fetch(new Request(url.toString(), request));
}
