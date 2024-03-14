const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,ts,svelte}"],
	theme: {
		screens: {
			...defaultTheme.screens,
		},
		extend: {
			minWidth: {
				screen: "100vw",
			},
			boxShadow: {
				protruding: "inset 0 -2px 4px 2px rgb(0 0 0 / 0.10);",
			},
		},
	},
	plugins: [],
};
