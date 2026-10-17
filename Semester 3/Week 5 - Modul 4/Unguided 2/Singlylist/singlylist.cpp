#include <iostream>
#include "Singlylist.h"
using namespace std;

void createList(List &L)
{
    L.First = Nil;
}

address alokasi(infotype x)
{
    address P = new ElmList;
    P->info = x;
    P->next = Nil;
    return P;
}

void dealokasi(address &P)
{
    delete P;
}

void printInfo(List L)
{
    address P = L.First;
    while (P != Nil)
    {
        cout << P->info << " ";
        P = P->next;
    }
    cout << endl;
}

void insertFirst(List &L, address P)
{
    P->next = L.First;
    L.First = P;
}

void deleteFirst(List &L, address &P)
{
    P = L.First;
    L.First = (L.First)->next;
    P->next = Nil;
}

void deleteLast(List &L, address &P)
{
    address Q = L.First;
    while (Q->next->next != Nil)
    {
        Q = Q->next;
    }
    P = Q->next;
    Q->next = Nil;
}

void deleteAfter(List &L, address Prec, address &P)
{
    P = Prec->next;
    Prec->next = P->next;
    P->next = Nil;
}

int nbList(List L)
{
    int jumlah = 0;
    address P = L.First;
    while (P != Nil)
    {
        jumlah++;
        P = P->next;
    }
    return jumlah;
}

void deleteList(List &L)
{
    address P;
    while (L.First != Nil)
    {
        deleteFirst(L, P);
        dealokasi(P);
    }
}