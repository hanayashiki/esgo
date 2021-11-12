# Monoid TypeScript NPM Boilerplate

A TypeScript NPM boilerplate that suits how we usually write TypeScript. 

## Features

1. Outputs modern style typescript package in `dist/{cjs,esm,types}/` using `tsc` and `esbuild`.

1. Respects modern package.json
   1. `type` set to `"module"`
   1. `main` pointed to `./dist/cjs/index.cjs` for legacy NodeJS in commonjs format.
   1. `module` pointed to `./dist/esm/index.mjs` for bundlers in esmodule format.
   1. `types` for typescript declarations.
   1. `exports` following [Conditional Exports](https://nodejs.org/api/packages.html#conditional-exports) for newer NodeJS to find esmodule format when using `import` and commonjs format when using `require`.

1. Test or demonstrate your package under `demo/{node,browser}/`.

1. Only files under `dist/` will be published to npm.

1. `dist/` will NOT be gitignored by default.

1. `.eslintrc.js` based on [airbnb](https://www.npmjs.com/package/eslint-config-airbnb-typescript).

1. [esmo](https://github.com/antfu/esno) installed to load typescript for node.

## Commands

```bash
  yarn build
```

Build the output directory.

```bash
  yarn lint
```

Lint the project.

```bash
  yarn cli
```

Execute the cli with node.

```bash
  yarn ts <file-name>
```

Execute arbitrary typescript file with [esmo](https://github.com/antfu/esno).

## Customization

1. Find `ts-npm-boilerplate`, replace that with your package name.

2. Find `your-cli-name`, replace that with your cli name.

## Usage

1. Degit this repository

```
npx degit https://github.com/MonoidDev/ts-npm-boilerplate
```

2. Use this template

Click the green button `Use this template` on the top-right side of this GitHub repository (https://github.com/MonoidDev/ts-npm-boilerplate).


