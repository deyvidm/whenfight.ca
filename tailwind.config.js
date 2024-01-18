const colors = require('tailwindcss/colors');
const plugin = require('tailwindcss/plugin');

module.exports = {
  daisyui: {
		themes: ["light", "dark", "garden","retro","cupcake","bumblebee","emerald","corporate","synthwave","halloween","aqua","lofi","pastel","fantasy","wireframe","black","luxury","dracula"],
	},
  plugins: [require("daisyui")],
  mode: 'jit',
  content: ['./src/**/*.{html,svelte,js,ts}', './node_modules/svelte-ux/**/*.{svelte,js}'],
};