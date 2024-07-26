#include <iostream>

void add_arrays(int* __restrict__ a, int* __restrict__ b, int* __restrict__ result, std::size_t size) {
    for (std::size_t i = 0; i < size; ++i) {
        result[i] = a[i] + b[i];
    }
}

int main() {
    const std::size_t size = 5;
    int a[size] = {1, 2, 3, 4, 5};
    int b[size] = {6, 7, 8, 9, 10};
    int result[size] = {0};

    add_arrays(a, b, result, size);

    for (std::size_t i = 0; i < size; ++i) {
        std::cout << result[i] << " ";
    }
    std::cout << std::endl;

    return 0;
}
