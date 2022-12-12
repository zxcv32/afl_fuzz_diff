#include <iostream>
#include <array>
#include <memory>

using namespace std;


// Run command
std::string exec(const char* cmd) {
    std::array<char, 128> buffer;
    std::string result;
    std::unique_ptr<FILE, decltype(&pclose)> pipe(popen(cmd, "r"), pclose);
    if (!pipe) {
        throw std::runtime_error("popen() failed!");
    }
    while (fgets(buffer.data(), buffer.size(), pipe.get()) != nullptr) {
        result += buffer.data();
    }
    return result;
}

int main() {
    std::string input;
    std::getline(std::cin, input);

    std::string commandA = "~/Desktop/A 2>&1 <<EOF\n" + input+"\nEOF";
    std::string outputA = exec(commandA.c_str());
    if (!outputA.empty() && outputA[outputA.length() - 1] == '\n') {
        outputA.erase(outputA.length() - 1);
    }

    std::string commandB = "~/Desktop/B 2>&1 <<EOF\n" + input+"\nEOF";
    std::string outputB = exec(commandB.c_str());
    if (!outputB.empty() && outputB[outputB.length() - 1] == '\n') {
        outputB.erase(outputB.length() - 1);
    }

    std::string commandC = "~/Desktop/C 2>&1 <<EOF\n" + input+"\nEOF";
    std::string outputC = exec(commandC.c_str());
    if (!outputC.empty() && outputC[outputC.length() - 1] == '\n') {
        outputC.erase(outputC.length() - 1);
    }
    if (!(outputA == outputB && outputA == outputC)){
        ::abort();
    }
    std::cout << "Equal";
    return 0;
}

