# Полезное для GoLang.

## Производительность

```shell
time go run main.go
```

## Трасировка

```shell
strace go run main.go
```

## Трасировка сборщика мусора

```shell
GODEBUG=gctrace=1 go run examples/gcoll.go
```

## Код в Assebler'е

```shell
GOOS=linux GOARCH=amd64 go tool compile -gcflags -S main.go
```

или

```shell
GOOS=linux GOARCH=amd64 go build -gcflags -S main.go
```

Строки, содержащие директивы FUNCDATA и PCDATA , автоматически генерируются компилятором Go, затем используются
сборщиком мусора Go.

**GOOS** — android , darwin , dragonfly , freebsd , linux , nacl , netbsd , openbsd , plan9 , solaris , windows и zos;  
**GOARCH** — 386 , amd64 , amd64p32 , arm , armbe , arm64 , arm64be , ppc64 , ppc64le , mips , mipsle , mips64 ,
mips64le , mips64p32 , mips64p32le , ppc , s390 , s390x , spar , spar и sparc64~~

## AST - дерево

```shell
go tool compile -W main.go
```

## Генерация Web - Assembly

```shell
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

## Документация

```shell
go doc http Get
```

## Профилирование

```shell
go tool pprof main.prof
```

## Тестирование

1. Выполнение всех тестов

```shell
go test -v
```

2. Запуск отдельного теста

```shell
 go test -run ExampleGetComputerScienceCategoryArticles ./pkgs/arxiv -count=1 -v
```
