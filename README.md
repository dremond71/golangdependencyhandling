# golangdependencyhandling

## Changing Dependencies

With a go.mod

```
require (
	github.com/dremond71/golangdependencyexampleone v1.0.0
	github.com/dremond71/golangdependencyexampletwo v1.0.0
)
```

I tried to ONLY upgrade github.com/dremond71/golangdependencyexampletwo to v1.1.0
with

```sh
go get github.com/dremond71/golangdependencyexampletwo@v1.1.0
```

but it also forced a change to github.com/dremond71/golangdependencyexampleone

```sh
go get: upgraded github.com/dremond71/golangdependencyexampleone v1.0.0 => v1.1.0
go get: upgraded github.com/dremond71/golangdependencyexampletwo v1.0.0 => v1.1.0
```

Resulting go.mod :

```
require (
	github.com/dremond71/golangdependencyexampleone v1.1.0
	github.com/dremond71/golangdependencyexampletwo v1.1.0
)
```


## Output

```sh
obiwan kenobi should look like obiwan kenobi
anakin skywalker should look like anakin skywalker
```