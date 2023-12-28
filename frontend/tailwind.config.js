/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,ts,svelte}"],
	theme: {
		extend: {
			minWidth: {
				screen: "100vw",
			},
		},
	},
	plugins: [],
};
