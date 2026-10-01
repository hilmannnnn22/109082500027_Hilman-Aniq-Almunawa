#include <iostream>
using namespace std;

// Call by Value
void tukarValue(int x, int y) {
    int temp;

    temp = x;
    x = y;
    y = temp;
}

// Call by Pointer
void tukarPointer(int *x, int *y) {
    int temp;

    temp = *x;
    *x = *y;
    *y = temp;
}

// Call by Reference
void tukarReference(int &x, int &y) {
    int temp;

    temp = x;
    x = y;
    y = temp;
}

int main() {
    int a, b;

    // CALL BY VALUE
    a = 4;
    b = 6;

    cout << "=== CALL BY VALUE ===" << endl;
    cout << "Sebelum ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    tukarValue(a, b);

    cout << "Setelah ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;


    // CALL BY POINTER
    a = 4;
    b = 6;

    cout << "\n=== CALL BY POINTER ===" << endl;
    cout << "Sebelum ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    tukarPointer(&a, &b);

    cout << "Setelah ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    // CALL BY REFERENCE
    a = 4;
    b = 6;

    cout << "\n=== CALL BY REFERENCE ===" << endl;
    cout << "Sebelum ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    tukarReference(a, b);

    cout << "Setelah ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    return 0;
}