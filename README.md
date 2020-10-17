![Generic badge](https://github.com/konoui/alfred-tldr/workflows/test/badge.svg)

## alfred tldr
tldr alfred workflow written in go.

## Install
- Download and open the workflow with terminal for bypassing GateKeeper on macOS.
```
$ curl -O -L https://github.com/konoui/alfred-tldr/releases/latest/download/tldr.alfredworkflow && open tldr.alfredworkflow
```

- Build the workflow on your computer.
```
$ make package
$ ls
tldr.alfredworkflow (snip)
```

## Usage
`tldr <query>`

Options   
`-u` option updates update local database (tldr repository).  
`-p` option selects platform from `linux`,`osx`,`sunos`,`windows`.  

![alfred-tldr](./alfred-tldr.png)

## License
MIT License.
