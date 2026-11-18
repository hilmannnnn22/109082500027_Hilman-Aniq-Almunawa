#include <iostream>
#include "stack.h"
using namespace std;

void createStack(Stack &S) {
    S.top = 0;
}

void push(Stack &S, infotype x) {
    if (S.top == 20) {
        cout << "Stack penuh." << endl;
        return;
    }

    S.info[S.top] = x;
    S.top++;
}

infotype pop(Stack &S) {
    if (S.top == 0) {
        cout << "Stack kosong." << endl;
        return 0;
    }

    S.top--;
    return S.info[S.top];
}

void printInfo(Stack S) {
    cout << "[TOP]";

    for (int i = S.top - 1; i >= 0; i--) {
        cout << " " << S.info[i];
    }

    cout << endl;
}

void balikStack(Stack &S) {
    int kiri = 0;
    int kanan = S.top - 1;

    while (kiri < kanan) {
        infotype temp = S.info[kiri];
        S.info[kiri] = S.info[kanan];
        S.info[kanan] = temp;

        kiri++;
        kanan--;
    }
}

//code tambahan yg di pake buat jawab soal no 2
void pushAscending(Stack &S, infotype x) {
    if (S.top == 20) {
        cout << "Stack penuh." << endl;
        return;
    }

    int i = S.top;

    while (i > 0 && S.info[i - 1] > x) {
        S.info[i] = S.info[i - 1];
        i--;
    }

    S.info[i] = x;
    S.top++;
}

//code tambahan yg di pake buat jawab soal no 3
void getInputStream(Stack &S) {
    char input;

    while (cin.get(input) && input != '\n') {
        if (input >= '0' && input <= '9') {
            push(S, input - '0');
        }
    }
}