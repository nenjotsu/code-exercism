#include <stdio.h>

void doSomethingWithInts(int a = 1, int b = 4)
{
    // Do something with the ints here
    printf("a %d\n", a);
    printf("b %d\n", b);
}

int main()
{
    doSomethingWithInts();      // a = 1,  b = 4
    doSomethingWithInts(20);    // a = 20, b = 4
    doSomethingWithInts(20, 5); // a = 20, b = 5
}