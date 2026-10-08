#include <iostream>
#include "array2D.h"
using namespace std;

int main()
{
    int array1[3][3] = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };

    int array2[3][3] = {
        {10, 11, 12},
        {13, 14, 15},
        {16, 17, 18}
    };

    int a = 20;
    int b = 30;

    int *p1 = &a;
    int *p2 = &b;

    cout << "Array 1 sebelum ditukar :" << endl;
    tampilArray(array1);

    cout << "\nArray 2 sebelum ditukar :" << endl;
    tampilArray(array2);

    tukarArray(array1, array2, 1, 1);

    cout << "\nArray 1 setelah ditukar :" << endl;
    tampilArray(array1);

    cout << "\nArray 2 setelah ditukar :" << endl;
    tampilArray(array2);

    cout << "\nNilai sebelum pointer ditukar : " << *p1 << " dan " << *p2 << endl;

    tukarPointer(p1, p2);

    cout << "Nilai setelah pointer ditukar : " << *p1 << " dan " << *p2 << endl;

    return 0;
}