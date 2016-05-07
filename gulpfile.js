var gulp  = require('gulp');
var child = require('child_process');
var server = null;

gulp.task('default', ['build', 'watch', 'spawn']);

gulp.task('watch', function() {
  gulp.watch('./**/*.go', ['build','spawn']);
  gulp.watch('./**/*.go', doFmt)
  gulp.watch('./**/*.go', doVet)
});

function doFmt(e) {
  var p = child.spawnSync('go', ['fmt', e.path]);
  dumpStreams("STDOUT:\n", "STDERR:\n", p);
  return p;
}
function doVet(e) {
  var p = child.spawnSync('go', ['vet', e.path]);
  dumpStreams("STDOUT:\n", "STDERR:\n", p);
  return p;
}

gulp.task('build', function() {
  var p = child.spawnSync('go', ['install']);
  dumpStreams("STDOUT:\n", "STDERR:\n", p); 
  return p;
});

gulp.task('spawn', function() {
  if (server)
    server.kill();
  server = child.spawn('swarm-hooks', ['--debug','version']);
  server.stderr.on('data', function(data) {
    process.stdout.write(data.toString());
  });
  server.stdout.on('data', function(data) {
    process.stdout.write(data.toString());
  });
});

function dumpStreams(outPrefix, errPrefix, p) {
  process.stdout.write(outPrefix + p.stdout.toString());
  process.stdout.write(errPrefix + p.stderr.toString());
}
