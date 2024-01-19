const colors = require('tailwindcss/colors');
const plugin = require('tailwindcss/plugin');

module.exports = {
  daisyui: {
		themes: [ "light", "dark", "cupcake", "bumblebee", "emerald", "corporate", "synthwave", "retro", "cyberpunk", "valentine", "halloween", "garden", "forest", "aqua", "lofi", "pastel", "fantasy", "wireframe", "black", "luxury", "dracula", "cmyk", "autumn", "business", "acid", "lemonade", "night", "coffee", "winter", "dim", "nord", "sunset"],
	},
  plugins: [require("daisyui")],
  mode: 'jit',
  content: ['./src/**/*.{html,svelte,js,ts}', './node_modules/svelte-ux/**/*.{svelte,js}'],
};