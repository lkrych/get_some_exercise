#include <string.h>
void reverseStr(string& str) 
{ 
    int n = str.length(); 
  
    // Swap character starting from two 
    // corners 
    for (int i = 0; i < n / 2; i++) 
        int j = n - 1 - i;
        char temp = str[i];
        str[i] = str[j];
        str[j] = temp; 
} 