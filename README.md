# FirstPass ![Github tests status badge](https://github.com/agustin-carnevale/firstpass/actions/workflows/ci.yml/badge.svg)

This project is a very minimal/simple tool intended to help developers run very basic checks on a project they received from "untrusted" sources. The motivation for this project was initially based on the fact that many fake recruiters send these projects with malicious code to developers as part of a take home challege during the interviews process. Or also fake clients asking freelancers to work on the project, to modify something, etc.

### Status

This was just started and is a Work-In-Progress. Also any help or contributions, in code or ideas is very much welcome.

## List of basic checks it should run:

- [x] Basic js/ts files check for eval()
- [x] ESlint security check
- [ ] package.json/dependencies check
- [ ] semgrep
- [ ] other third-party tools

## Run locally

You will need `Go` installed in your machine.

See these [INSTALLATION DOCS](./INSTALLATIONS.md)

### Install firstpass

You can now install the cli program with Go:

```bash
go install github.com/agustin-carnevale/firstpass@latest
```

Print the version to check everything worked fine:

```bash
firstpass version
```

### Usage

From a remote repository

```bash
firstpass scan -r <repo_url>
```

For local dir

```bash
firstpass scan -d <path_to_dir>
```
