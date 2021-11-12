#!/usr/bin/env node
import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

const main = () => {
  const finalArgv = process.argv[0] === 'your-cli-name'
    ? ['node', ...process.argv]
    : process.argv;

  yargs(hideBin(finalArgv))
    .command(
      'command1 [config]',
      'Helper text for command1. ',
      (_argv) => {
        // Handle the command here
      },
    )
    .help()
    .demand(1, 'Must provide a valid command')
    .parse();
};

main();
