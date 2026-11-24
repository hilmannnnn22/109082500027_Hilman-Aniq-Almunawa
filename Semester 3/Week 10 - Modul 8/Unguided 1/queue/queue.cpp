#include "queue.h"
#include <iostream>
using namespace std;

// Mengecek apakah Queue kosong
bool isEmptyQueue(queue Q) {
    return (Q.head == -1 && Q.tail == -1);
}

// Mengecek apakah Queue penuh
bool isFullQueue(queue Q) {
    return (Q.tail == MAKSIMAL - 1);
}

// Membuat Queue kosong
void CreateQueue(queue &Q) {
    Q.head = -1;
    Q.tail = -1;
}

// Menambahkan elemen ke dalam Queue
void enqueue(queue &Q, infotype x) {
    if (isFullQueue(Q)) {
        cout << "Queue penuh!" << endl;
    } else {
        if (isEmptyQueue(Q)) {
            Q.head = 0;
            Q.tail = 0;
        } else {
            Q.tail++;
        }

        Q.info[Q.tail] = x;
    }
}

// Menghapus elemen paling depan
infotype dequeue(queue &Q) {
    if (isEmptyQueue(Q)) {
        cout << "Queue kosong!" << endl;
        return -1;
    } else {
        infotype x = Q.info[Q.head];

        if (Q.head == Q.tail) {
            Q.head = -1;
            Q.tail = -1;
        } else {
            for (int i = 0; i < Q.tail; i++) {
                Q.info[i] = Q.info[i + 1];
            }
            Q.tail--;
        }

        return x;
    }
}

// Menampilkan isi Queue
void printInfo(queue Q) {
    if (isEmptyQueue(Q)) {
        cout << "-1 - -1 | empty queue" << endl;
    } else {
        cout << Q.head << " - " << Q.tail << " | ";

        for (int i = Q.head; i <= Q.tail; i++) {
            cout << Q.info[i] << " ";
        }

        cout << endl;
    }
}