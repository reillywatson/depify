depify is a command-line utility that, given a directory and a Git repo URL, will try to identify an SHA that corresponds to that directory's contents.

Installing: `go get github.com/reillywatson/depify`

Using: `depify <repo> <path>`

Example: `depify github.com/reillywatson/goloose /path/to/goloose`

Example output: `e4c9c47d9c1ebae4151e081add90fc11bcfa498a`

If depify can't find a match, it outputs the commit that has the smallest diff from the directory's contents.