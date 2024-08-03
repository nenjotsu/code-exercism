#include <iostream>
#include <vector>
#include <map>

// Helper function to generate combinations
void generateCombinations(const std::vector<int>& args, std::map<int, int>& occurrences, std::vector<int>& combination, int depth) {
    // Base case: when the combination has 4 elements, print it
    if (depth == 4) {
        for (int num : combination) {
            std::cout << num;
        }
        std::cout << '\n';
        return;
    }

    // Recursive case: generate combinations
    for (int i = 0; i < args.size(); ++i) {
        if (occurrences[args[i]] < 2) { // Check if the number can be used
            combination.push_back(args[i]);
            occurrences[args[i]]++;
            generateCombinations(args, occurrences, combination, depth + 1);
            occurrences[args[i]]--; // Backtrack
            combination.pop_back();
        }
    }
}

// Variadic template function to accept 2 to 4 arguments
template<typename... Args>
void generateCodes(Args... args) {
    static_assert(sizeof...(args) >= 2 && sizeof...(args) <= 4, "Function accepts 2 to 4 arguments.");
    std::vector<int> argVec = { args... };
    std::map<int, int> occurrences; // To keep track of occurrences of each number
    std::vector<int> combination;
    generateCombinations(argVec, occurrences, combination, 0);
}

int main() {
    // generateCodes(1, 2);        // Example with 2 arguments
    // generateCodes(1, 2, 3);     // Example with 3 arguments
    generateCodes(4, 6, 8, 9);  // Example with 4 arguments
    return 0;
}
