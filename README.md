# dedup

Dedup is a simple command line utility to take a list of files from stdin and use
sha1 to generate checksums. It keeps a list of duplicate checksums and displays a
list at the end. I wrote this to go through my photos directory and tell me what
files had matching checksums.

# usage

`find . | dedup` will operate normally but in cases where you want to use grep to
filter the results of `find` the `-grep` option strips the line number from the
output.

It was not built for awesomeness, only utility. On a maxed out MBP it can crunch
91GB of photos (17,214 files) in a little over 3 min.
