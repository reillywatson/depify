depify is a command-line utility that, given a directory and a Git repo URL, will try to identify an SHA that corresponds to that directory's contents.

Installing: `go install github.com/reillywatson/depify`

Using: `depify <repo> <path>`

Example: `depify github.com/reillywatson/goloose /path/to/goloose`

Example output: `e4c9c47d9c1ebae4151e081add90fc11bcfa498a`
