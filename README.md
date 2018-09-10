Basic virtual machine written with go.

Fibonacci demonstration:

```
go build main.go
./main
```

Main function can pass 20-bytes of RAM (represented with an array) and the CPU will parse against the following ISA:

```
0x01 LOAD
0x02 STORE
0x03 ADD
0x04 SUB
0x05 PRINT
0xff HALT
```

The first 14 bytes of memory are reserved for instructions. Then 2 bytes are reserved for output, and 4 bytes are split between two 2-byte inputs. 14 + 2 + (2 \* 2) = 20

Processor has 2 general purpose internal registers:

```
0x01 R1
0x02 R2
```

`0x00` register is reserved for the program counter. The registers are persistent.
