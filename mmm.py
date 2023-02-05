#!/usr/bin/env python3.11

import tomllib
import argparse

# TOML-related constants
FILE = "mmm.toml"
COMMAND = "command"
SCRIPT = "script"

with open(FILE, "rb") as f:
    config = tomllib.load(f)

if not COMMAND in config:
    print(f"loaded {FILE}, but found no entries in '{COMMAND}' table.")
    exit(1)

parser = argparse.ArgumentParser(
        prog="mmm",
        description=f"Runs commands for a project. Commands are defined in {FILE} files."
)
parser.add_argument("command", help="The command to run from mmm.toml.")
args = parser.parse_args()

if not args.command in config[COMMAND]:
    print(f"loaded {FILE}, but found no command named {COMMAND}.{args.command}.")
    exit(2)

toRun = []

selectedCommand = config[COMMAND][args.command]

if isinstance(selectedCommand , str):
    toRun.append(selectedCommand)
elif isinstance(selectedCommand[SCRIPT], str):
    toRun.append(selectedCommand[SCRIPT])
# TODO: Handle "script is an array of strings"

print(f"I would run {toRun}")
