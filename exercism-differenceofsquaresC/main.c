#include <stdio.h>

unsigned int square_of_sum( unsigned int n ){
    unsigned int sum = 0;
    //square of sum
    for (unsigned int i=1; i<=n; i++){
        sum += i;
    }
   return  sum * sum;
}
 unsigned int sum_of_squares(unsigned int n) {
         //sum of squares
       unsigned   int sumOfSquares = 0;
        for (unsigned int j=1; j<=n; j++){
                sumOfSquares += j*j;
        }
        return sumOfSquares;
    }

unsigned int difference_of_squares(unsigned int number){
    return square_of_sum(number) - sum_of_squares(number);
}

int main() {
    // Write C code here
    printf("%d\n", difference_of_squares(5));
    printf("%d\n", difference_of_squares(20));

    return 0;
}