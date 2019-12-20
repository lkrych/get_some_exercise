//reverse a string in place
#include "string"
using namespace std;
void reverseStr(std::string& str) 
{ 
    int n = str.length(); 
    int i,j;
    // Swap character starting from two 
    // corners 
    for (int i = 0; i < n / 2; i++) 
        int j = n - 1 - i;
        char temp = str[i];
        str[i] = str[j];
        str[j] = temp; 
} 