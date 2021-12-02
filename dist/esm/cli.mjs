#!/usr/bin/env node

// src/cli.ts
import yargs from "yargs";
import {hideBin} from "yargs/helpers";
var main = () => {
  const finalArgv = process.argv[0] === "your-cli-name" ? ["node", ...process.argv] : process.argv;
  yargs(hideBin(finalArgv)).command("command1 [config]", "Helper text for command1. ", (_argv) => {
  }).help().demand(1, "Must provide a valid command").parse();
};
main();
