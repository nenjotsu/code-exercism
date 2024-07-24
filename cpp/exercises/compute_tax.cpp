#include <iostream>

int main() {
  double sales = 95000;
  double stateTax = 0.04;
  double countyTax = 0.02;
  double totalTax = stateTax + countyTax;
  double total = sales - (sales * totalTax);
  std::cout << total << std::endl;
  
  return 0;
}
