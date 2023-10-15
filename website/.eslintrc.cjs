module.exports = {
	root: true,
	plugins: ["svelte3", "@typescript-eslint"],
	extends: [
		"eslint:recommended",
		"plugin:@typescript-eslint/eslint-recommended",
		"plugin:@typescript-eslint/recommended",
		"prettier",
	],
	overrides: [
		{ files: ["*.svelte"], processor: "svelte3/svelte3" },
		{ files: ["*.ts"], parser: "@typescript-eslint/parser" },
	],
	parserOptions: {
		sourceType: "module",
		ecmaVersion: 2020,
	},
	env: {
		browser: true,
		es2017: true,
		node: true,
	},
};
