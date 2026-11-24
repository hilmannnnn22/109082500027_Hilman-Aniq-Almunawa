#include "queue.h"
#include <iostream>
using namespace std;

// Mengecek apakah Queue kosong
bool isEmptyQueue(queue Q) {
    return (Q.head == -1 && Q.tail == -1);
}

// BAGIAN ALTERNATIF 3: Mengecek apakah Queue penuh
bool isFullQueue(queue Q) {
    return ((Q.tail + 1) % MAKSIMAL == Q.head);
}

// Membuat Queue kosong
void CreateQueue(queue &Q) {
    Q.head = -1;
    Q.tail = -1;
}

// Menambahkan data ke dalam Queue
void enqueue(queue &Q, infotype x) {
    if (isFullQueue(Q)) {
        cout << "Queue penuh!" << endl;
    } else {
        if (isEmptyQueue(Q)) {
            Q.head = 0;
            Q.tail = 0;
        } else {
            // BAGIAN ALTERNATIF 3: Tail bergerak melingkar
            Q.tail = (Q.tail + 1) % MAKSIMAL;
        }

        Q.info[Q.tail] = x;
    }
}

// Menghapus data paling depan dari Queue
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
            // BAGIAN ALTERNATIF 3: Head bergerak melingkar
            Q.head = (Q.head + 1) % MAKSIMAL;
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

        int i = Q.head;

        while (true) {
            cout << Q.info[i] << " ";

            if (i == Q.tail) {
                break;
            }

            // BAGIAN ALTERNATIF 3: Indeks kembali ke awal setelah mencapai akhir
            i = (i + 1) % MAKSIMAL;
        }

        cout << endl;
    }
}