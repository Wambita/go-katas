#include <stdio.h>

char *ordinal(int v){
//catch exceptions 11, 12, 13
if (v==11 || v==12|| v==13){
    return("th");
}
/*
if num ends in 1 == st
               2 == nd
               3 == rd
               others th
*/
/*just get the last values*/
v = v%10;
if (v==1){
    return("st");
}
if (v==2){
    return("nd");
}
if (v==3){
    return("rd");
}
return("th");
}

//main function
int main(){

    //output a  table of 100 numbers
    for (c:=1; c<=20; c++){
        printf("%d%s\t%d%s\t%d%s\t%d%s\t%d%s\n", 
        c, ordinal(c),
        c+20, ordinal(c+20),
        c+40, ordinal(c+40),
        c+60, ordinal(c+60),
        c+80, ordinal(c+80)
        );
     //ordinal function is called for every number from 1 to 100 in a loop.
    }
    return 0;
}