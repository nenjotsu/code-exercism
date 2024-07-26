#include <iostream>
template <typename T>
T maximum(T x, T y) {
  return (x > y) ? x : y;
}

template <typename T, typename U>
auto minimum(T x, U y) {
  return (x < y) ? x : y;
}


using namespace std;

int main() {
  cout << maximum(1, 2) << endl;
  cout << minimum(1.2, 2) << endl;
  return 0;
}