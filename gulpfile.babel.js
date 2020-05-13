import { series, parallel } from 'gulp'

const { spawn } = require('child_process');

function runGo(cb) {
  return spawn('go run main.go', [], { stdio: 'inherit' })
    .on('close', cb)
    .on('exit', () => cb())
    .on('error', err => {
      console.error(err);
      process.exit(err);
    });
}

const start = series(
    runGo
)
export {
    start
}