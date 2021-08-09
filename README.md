# Reproduce bug in beta fuzzer

```
gotip test ./sub -fuzz=Fuzz
```

will run but give no output (or will be buffered until a crash or ctrl+c).

```
cd sub
gotip  test -fuzz=Fuzz
```
works fine.
