/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontSize: {
				'xs': '0.625rem', /* 12px */
				'sm': '0.75rem', /* 12px */
				'base': '0.875rem', /* 14px */
				'lg': '1rem', /* 16px */
				'xl': '1.125rem', /* 18px */
				'2xl': '1.375rem' /* 22px */
			}
		}
	},
	daisyui: {
		themes: [
			{
				themesquare: {
					'primary': '#4B99FF',
					// 'primary-focus': '#355072',
					// 'primary-content': '#6395d7',
					// "primary-content": "#5459d7",

					'secondary': '#21259D',
					'secondary-focus': '#5f61af',
					'secondary-content': '#abadf5',

					'accent': '#B3CDEF',
					'accent-focus': '#91acc5',
					'accent-content': '#f1f5f9',

					'neutral': '#707070',
					'neutral-focus': '#91acc5',
					'neutral-content': '#d3d3d3',

					'base-100': '#ffffff',

					'info': '#3ABFF8',
					'success': '#36D399',
					'warning': '#FBBD23',
					'error': '#F87272',

					// root set: from 0.5 to 2rem
					'--rounded-btn': '0.25rem', // border radius rounded-ui utility class, used in buttons and similar element
					'--rounded-box': '0.25rem', // border radius rounded-box utility class, used in card and other large boxes

					'--padding-btn': '20rem', // border radius rounded-ui utility class, used in buttons and similar element

					'--rounded-badge': '1.9rem', // border radius rounded-badge utility class, used in badges and similar
					'--animation-btn': '0.25s', // duration of animation when you click on button
					'--animation-input': '0.2s', // duration of animation for inputs like checkbox, toggle, radio, etc
					'--btn-text-case': 'capitalize', // set default text transform for buttons
					'--btn-focus-scale': '0.95', // scale transform of button when you focus on it
					'--border-btn': '1px', // border width of buttons
					'--tab-border': '1px', // border width of tabs
					'--tab-radius': '0.5rem' // border radius of tabs
				},
				themeroofix: {
					'primary': '#4B99FF',
					// "primary-content": "#5459d7",

					'secondary': '#21259D',

					'accent': '#B3CDEF',
					'accent-content': '#f1f5f9',
					'accent-focus': '#91acc5',

					'neutral': '#707070',
					'neutral-content': '#d3d3d3',
					'neutral-focus': '#91acc5',

					'base-100': '#ffffff',

					'info': '#76CCE0',
					'success': '#74da88',
					'warning': '#EFA152',
					'error': '#F93434',

					'--rounded-box': '1rem', // border radius rounded-box utility class, used in card and other large boxes

					// root set: from 0.5 to 2rem
					'--rounded-btn': '2rem', // border radius rounded-ui utility class, used in buttons and similar element
					'--padding-btn': '20rem', // border radius rounded-ui utility class, used in buttons and similar element

					'--rounded-badge': '1.9rem', // border radius rounded-badge utility class, used in badges and similar
					'--animation-btn': '0.25s', // duration of animation when you click on button
					'--animation-input': '0.2s', // duration of animation for inputs like checkbox, toggle, radio, etc
					'--btn-text-case': 'uppercase', // set default text transform for buttons
					'--btn-focus-scale': '0.95', // scale transform of button when you focus on it
					'--border-btn': '1px', // border width of buttons
					'--tab-border': '1px', // border width of tabs
					'--tab-radius': '0.5rem' // border radius of tabs
				}
			},
			'dark', 'cupcake'
		]
	},
	plugins: [
		require('@tailwindcss/typography'),
		require('prettier-plugin-tailwindcss'),
		require('@tailwindcss/aspect-ratio'),
		require('daisyui'),
	]

};
