// gcc -Og -S mstore.c
// gcc -Og -S -masm=intel mstore.c
// gcc -Og -c mstore.c
// objdump -d mstore.o

long mult2(long, long);

void multstore(long x, long y, long *dest)
{
    long t = mult2(x, y);
    *dest = t;
}