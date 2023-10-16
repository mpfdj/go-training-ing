# Exercise 3

```
go test -gcflags "-N -l" -cpuprofile cpu.out -memprofile mem.out --bench .
```

```
go tool pprof -http :9000 cpu.out
```

```
go tool pprof -http :9000 mem.out
```
