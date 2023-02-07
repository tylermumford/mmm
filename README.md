# mmm

Also known as "mmm build" (because you've gotta be able to search the web for
something)

The `mmm` command is a reimagining of Make. With `mmm`, you can define a
cohesive set of commands for working with a codebase.

- Make: focused on target files, also handles general script running<br>
  mmm: focused on general script running, maybe handles target files

- Make: defines a file format (Makefile)<br>
  mmm: configured with TOML

- Make: many pages of documentation<br>
  mmm: designed to be simple, with less documentation (and prettier, too)

- Make: centralizes development commands<br>
  mmm: yep, same purpose

## Project status

Idea. Pre-alpha. Prototypes, untested, brand new.

## Design notes

This is not a "build system" in my mind. It's a project command organizer.

I hope to make it flexible enough to work with multiple shells. I use
PowerShell, zsh, and Nushell, and other people use many more. Can I make this
shell-agnostic?

Make is fantastic, but it's designed around specifying files as inputs and
outputs. Lots of dev commands don't produce output files: running unit tests,
running a server, building a container, running a linter... I want `mmm` to be
designed around both ideas (target files and just-run-the-commands), but
primarily for just-run-the-commands mode.

## License

TBD, but probably MIT.

## Other projects of note

- Make
- Mage
- Python
- Gulp
- Earthly
- npm scripts
- Scripts to Rule Them All
