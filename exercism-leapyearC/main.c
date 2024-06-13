//#include "leap.h" use this only when posting solution on exercism
#include <stdio.h>
#include <stdbool.h>

bool leap_year(int n){
    if (n%100==0){
        return (n%400==0);
    }
    if (n%4==0){
        return true;
    }
    return false;
}

//Use this main function for testing 
int main() {
    printf("%s\n", leap_year(2004) ? "true" : "false");
    printf("%s\n", leap_year(2000) ? "true" : "false");
    printf("%s\n", leap_year(2001) ? "true" : "false");
    printf("%s\n", leap_year(1900) ? "true" : "false");


}