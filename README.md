Compile C++ code
```shell
./afl-clang++ ./afl-fuzz-diff.cpp -o afl-fuzz-diff.cpp.binary
```

Run AFL until crash is detected
```shell
./afl-fuzz -i input -o output -- ./afl-fuzz-diff.cpp.binary
```