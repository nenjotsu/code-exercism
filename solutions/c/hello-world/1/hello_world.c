#include "hello_world.h"
#include <stdio.h>

// Define the function itself.
const char *hello(void)
{
   return "Hello, World!";
}

int main(void)
{
   printf("%s\n", hello());
}