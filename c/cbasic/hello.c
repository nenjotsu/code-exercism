#include <stdio.h>

int main(void) {
  // print "Hello, World!"
  printf("Hello, World!\n");

  // conditionals
  int x = 5;
  int y = 10;
  if (x < y) {
    printf("x is less than y\n");
  } else if (x > y) {
    printf("x is greater than y\n");
  } else {
    printf("x is equal to y\n");
  }
  return 0;
}