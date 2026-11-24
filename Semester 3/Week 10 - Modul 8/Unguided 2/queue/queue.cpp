#include "queue.h"
#include <iostream>
using namespace std;

bool isEmptyQueue(queue Q) {
    return (Q.head == -1 && Q.tail == -1);
}

// BERUBAH: Queue dianggap penuh jika head masih di 0
// dan tail sudah mencapai indeks terakhir.
bool isFullQueue(queue Q) {
    return (Q.head == 0 && Q.tail == MAKSIMAL - 1);
}

void CreateQueue(queue &Q) {
    Q.head = -1;
    Q.tail = -1;
}

void enqueue(queue &Q, infotype x) {
    if (isFullQueue(Q)) {
        cout << "Queue penuh!" << endl;
    } else {
        if (isEmptyQueue(Q)) {
            Q.head = 0;
            Q.tail = 0;
        } else {

            // BERUBAH: Jika tail sudah di ujung array,
            // data yang masih ada digeser ke posisi awal.
            if (Q.tail == MAKSIMAL - 1) {
                int selisih = Q.head;

                for (int i = Q.head; i <= Q.tail; i++) {
                    Q.info[i - selisih] = Q.info[i];
                }

                Q.tail = Q.tail - selisih;
                Q.head = 0;
            }

            Q.tail++;
        }

        Q.info[Q.tail] = x;
    }
}

infotype dequeue(queue &Q) {
    if (isEmptyQueue(Q)) {
        cout << "Queue kosong!" << endl;
        return -1;
    } else {
        infotype x = Q.info[Q.head];

        // BERUBAH: Head cukup maju satu posisi,
        // tidak perlu menggeser semua data seperti Alternatif 1.
        Q.head++;

        if (Q.head > Q.tail) {
            Q.head = -1;
            Q.tail = -1;
        }

        return x;
    }
}

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