import sveltePreprocess from 'svelte-preprocess';

export default {
	preprocess: sveltePreprocess({ postcss: true }),
	plugins: {
		tailwindcss: {},
		autoprefixer: {},
	}
};
