#include <iostream>
using namespace std;

void tukarPointer(int *a, int *b, int *c) {
    int temp;

    temp = *a;
    *a = *b;
    *b = *c;
    *c = temp;
}

void tukarReference(int &a, int &b, int &c) {
    int temp;

    temp = a;
    a = b;
    b = c;
    c = temp;
}

int main() {
    int a = 10, b = 20, c = 30;

    cout << "Nilai awal: " << a << " " << b << " " << c << endl;

    tukarPointer(&a, &b, &c);
    cout << "Setelah pointer: " << a << " " << b << " " << c << endl;

    tukarReference(a, b, c);
    cout << "Setelah reference: " << a << " " << b << " " << c << endl;

    return 0;
}