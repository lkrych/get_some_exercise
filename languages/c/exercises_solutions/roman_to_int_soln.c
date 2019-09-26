#include <stdio.h>
#include <string.h>

int lookup_roman(char numeral) {
    switch (numeral) {
        case 'I':
            return 1;
        case 'V':
            return 5;
        case 'X':
            return 10;
        case 'L':
            return 50;
        case 'C':
            return 100;
        case 'D':
            return 500;
        case 'M':
            return 1000;
        default:
            return 0;
    }
}

int romanToInt(char * s){
    int running_sum = 0;
    int i,x,y;
    for (i = 0; i < strlen(s) - 1; i++) {
        x = lookup_roman(s[i]);
        y = lookup_roman(s[i+1]);
        if (x >= y) {
            running_sum += x;
        } else {
            running_sum += (y - x);
            i++;
        }
    }
    if (i < strlen(s)) {
        running_sum += lookup_roman(s[i]);
    }
    return running_sum;
}




int main() {
    int ans1,ans2,ans3,ans4;

    char str1[] = "IV";
    int n1 = strlen(str1);

    char str2[] = "I";
    int n2 = strlen(str2);

    char str3[] = "MMMLCXVI";
    int n3 = strlen(str3);

    char str4[] = "";
    int n4 = strlen(str4);

    ans1 = roman_to_int()


}