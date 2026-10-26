#include "Singlylist.h"

void CreateList(List &L) {
    L.First = Nil;
}

address alokasi(infotype x) {
    address P = new ElmList;
    P->info = x;
    P->next = Nil;
    return P;
}

void dealokasi(address &P) {
    if (P != Nil) {
        delete P;
        P = Nil;
    }
}

void insertFirst(List &L, address P) {
    if (P != Nil) {
        P->next = L.First;
        L.First = P;
    }
}

void printInfo(List L) {
    address P = L.First;

    while (P != Nil) {
        cout << P->info << " ";
        P = P->next;
    }

    cout << endl;
}

address findElm(List L, infotype x) {
    address P = L.First;

    while (P != Nil) {
        if (P->info == x) {
            return P;
        }

        P = P->next;
    }

    return Nil;
}

int sumInfo(List L) {
    int total = 0;
    address P = L.First;

    while (P != Nil) {
        total += P->info;
        P = P->next;
    }

    return total;
}

void deleteList(List &L) {
    address P;

    while (L.First != Nil) {
        P = L.First;
        L.First = L.First->next;
        dealokasi(P);
    }
}