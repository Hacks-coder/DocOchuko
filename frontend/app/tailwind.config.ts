import containerQueries from '@tailwindcss/container-queries';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				'purple': '#4B30C3',
				'purple-1': '#6B50E0',
				'aqua': '#14B8A6',
				'gold': '#F59E0B',
				'gray': '#141414',
			}
		}
	},

	plugins: [typography, forms, containerQueries]
} satisfies Config;
