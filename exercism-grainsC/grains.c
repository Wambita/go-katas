//#include "grains.h" (use this in exercism)

#include <stdio.h>
#include <stdin.h>
uint64_t square(uint8_t index){   

    if (index<1 || index>64){
        return 0;
    }
    uint64_t out=1;
return out << (index-1);
}
uint64_t total(void){
    uint64_t sumOfGrains=0;
    uint64_t grainInEachSquare=1;
    for (uint8_t i=0; i<64;i++ ){

        sumOfGrains+=grainInEachSquare<<i;

    }

    return sumOfGrains;

} 

        

          


int main() {
    uint64_t grains = square(3);
    printf("%ld\n", grains);  // Correct format specifier for uint64_t
    return 0;  // Return 0 to indicate successful execution
}