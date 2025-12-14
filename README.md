# Advent of Code 2015

## Add new module (day) to the go workspace:

```
mkdir ./day01
cd ./day01
go mod init aoc2015/day01
cd ..
go work use ./day01
go work sync
```

## Run a module from the workspace

From the workspace root:

```
go run aoc2015/day01
```

## Use dependencies from other modules in the workspace

After adding the module to the workspace, public functions are accessible via their package names.

## Read a puzzle input file

Use the `util` module.

For example, when running a module from the workspace root:

```go
s := util.ReadInputString("input/day01-example.txt")
````

Note: Puzzle inputs should not be commited to the repository.
