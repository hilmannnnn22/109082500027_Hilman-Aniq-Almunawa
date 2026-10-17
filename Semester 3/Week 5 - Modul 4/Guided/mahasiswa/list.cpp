#include "list.h"
#include <iostream>
using namespace std;

// I.S = Initial State / kondisi awal
// F.S = Final State / kondisi akhir

// Fungsi untuk cek apakah list kosong atau tidak
bool isEmpty(linkedList list) {
    if (list.first == Nil) {
        return true;
    } else {
        return false;
    }
}

// Pembuatan linked list kosong
void createList(linkedList &list) {
    /*
       I.S. sembarang
       F.S. terbentuk list kosong
    */
    list.first = Nil;
}

// Pembuatan node baru dengan menerapkan manajemen memori
address alokasi(string nama, string nim, int umur) {
    /*
       I.S. sembarang
       F.S. mengembalikan alamat node baru dengan isidata
       sesuai parameter dan next = Nil
    */
    address nodeBaru = new node;

    nodeBaru->isidata.nama = nama;
    nodeBaru->isidata.nim = nim;
    nodeBaru->isidata.umur = umur;
    nodeBaru->next = Nil;

    return nodeBaru;
}

// Penghapusan node dengan menerapkan manajemen memori
void dealokasi(address &node) {
    /*
       I.S. P terdefinisi
       F.S. memori yang digunakan node dikembalikan ke sistem
    */
    if (node != Nil) {
        node->next = Nil;
        delete node;
        node = Nil;
    }
}

// Prosedur untuk insert / menambahkan node baru ke dalam list
void insertFirst(linkedList &List, address nodeBaru) {
    /*
       I.S. sembarang, P sudah dialokasikan
       F.S. menempatkan elemen list (node) pada awal list
    */
    if (nodeBaru != Nil) {
        nodeBaru->next = List.first;
        List.first = nodeBaru;
    }
}

// Menambahkan node setelah node Prev
void insertAfter(linkedList &List, address nodeBaru, address Prev) {
    /*
       I.S. sembarang, nodeBaru dan Prev alamat salah satu elemen list
       F.S. menempatkan elemen (node) sesudah elemen node Prev
    */
    if (Prev != Nil && nodeBaru != Nil) {
        nodeBaru->next = Prev->next;
        Prev->next = nodeBaru;
    } else {
        cout << "Node sebelumnya tidak valid!" << endl;
    }
}

// Menambahkan node di akhir list
void insertLast(linkedList &List, address nodeBaru) {
    /*
       I.S. sembarang, nodeBaru sudah dialokasikan
       F.S. menempatkan elemen nodeBaru pada akhir list
    */
    if (nodeBaru == Nil) {
        return;
    }

    if (isEmpty(List)) {
        List.first = nodeBaru;
    } else {
        address nodeBantu = List.first;

        while (nodeBantu->next != Nil) {
            nodeBantu = nodeBantu->next;
        }

        nodeBantu->next = nodeBaru;
    }

    nodeBaru->next = Nil;
}

// Prosedur untuk delete / menghapus node dalam list
void delFirst(linkedList &List) {
    /*
       I.S. list tidak kosong
       F.S. node pertama di list terhapus
    */
    address nodeHapus;

    if (!isEmpty(List)) {
        nodeHapus = List.first;
        List.first = List.first->next;

        nodeHapus->next = Nil;
        dealokasi(nodeHapus);
    } else {
        cout << "List kosong!" << endl;
    }
}

// Menghapus node terakhir
void delLast(linkedList &List) {
    /*
       I.S. list tidak kosong
       F.S. node terakhir di list terhapus
    */
    address nodeHapus, nodePrev;

    if (!isEmpty(List)) {
        nodeHapus = List.first;

        if (nodeHapus->next == Nil) {
            List.first = Nil;
            dealokasi(nodeHapus);
        } else {
            nodePrev = Nil;

            while (nodeHapus->next != Nil) {
                nodePrev = nodeHapus;
                nodeHapus = nodeHapus->next;
            }

            nodePrev->next = Nil;
            dealokasi(nodeHapus);
        }
    } else {
        cout << "List kosong" << endl;
    }
}

// Menghapus node setelah nodePrev
void delAfter(linkedList &List, address nodeHapus, address nodePrev) {
    /*
       I.S. list tidak kosong, Prev alamat salah satu elemen list
       F.S. node setelah Prev dihapus dari list
    */
    if (isEmpty(List)) {
        cout << "List kosong!" << endl;
    } else {
        if (nodePrev != Nil && nodePrev->next == nodeHapus &&
            nodeHapus != Nil) {

            nodePrev->next = nodeHapus->next;
            nodeHapus->next = Nil;

            dealokasi(nodeHapus);
        } else {
            cout << "Node sebelumnya (Prev) tidak valid!" << endl;
        }
    }
}

// Prosedur untuk menampilkan isi list
void printList(linkedList List) {
    /*
       I.S. mungkin kosong
       F.S. jika list tidak kosong, menampilkan semua info
       yang ada pada list
    */
    if (isEmpty(List)) {
        cout << "List kosong." << endl;
    } else {
        address nodeBantu = List.first;

        while (nodeBantu != Nil) {
            cout << "Nama : " << nodeBantu->isidata.nama
                 << ", NIM : " << nodeBantu->isidata.nim
                 << ", Usia : " << nodeBantu->isidata.umur
                 << endl;

            nodeBantu = nodeBantu->next;
        }
    }
}

// Function untuk menghitung jumlah node di dalam list
int nbList(linkedList List) {
    /*
       I.S. list sudah ada
       F.S. mengembalikan jumlah node di dalam list
    */
    int count = 0;
    address nodeBantu = List.first;

    while (nodeBantu != Nil) {
        count++;
        nodeBantu = nodeBantu->next;
    }

    return count;
}

// Prosedur untuk menghapus seluruh list
void deleteList(linkedList &List) {
    /*
       I.S. list sudah ada
       F.S. menghapus semua node di dalam list
    */
    address nodeBantu, nodeHapus;

    nodeBantu = List.first;

    while (nodeBantu != Nil) {
        nodeHapus = nodeBantu;
        nodeBantu = nodeBantu->next;

        dealokasi(nodeHapus);
    }

    List.first = Nil;

    cout << "List sudah terhapus!" << endl;
}