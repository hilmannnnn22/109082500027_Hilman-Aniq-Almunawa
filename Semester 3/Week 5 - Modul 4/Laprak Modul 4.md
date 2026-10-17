# <h1 align="center">Laporan Praktikum Modul 4 - SINGLY LINKED LIST (BAGIAN PERTAMA)   </h1>
<p align="center">Hilman Aniq Almunawa - 109082500027</p>

## Dasar Teori
Linked list merupakan salah satu bentuk struktur data yang berisi kumpulan elemen data yang saling terhubung satu sama lain. Berbeda dengan array yang bersifat statis, linked list diimplementasikan menggunakan pointer sehingga sifatnya lebih dinamis dan fleksibel karena ukurannya dapat bertambah atau berkurang sesuai kebutuhan program. Data yang disimpan di dalam linked list dapat berupa data tunggal maupun data majemuk seperti record atau struct.

Salah satu model dari linked list adalah Singly Linked List, yaitu linked list yang hanya memiliki satu arah pointer penunjuk. Setiap elemen atau node pada Singly Linked List dibagi menjadi dua bagian, yaitu bagian data (info) untuk menyimpan nilai dan bagian successor (next) yang berupa pointer untuk menunjuk ke alamat elemen berikutnya. Penelusuran pada struktur data ini hanya bisa dilakukan secara maju dari elemen pertama yang ditunjuk oleh pointer first hingga elemen terakhir yang menunjuk ke NULL atau Nil.

Pada bahasa C++, penerapan ADT Singly Linked List dilakukan secara modular dengan memisahkan deklarasi struct serta prototipe fungsi pada file header (.h) dan isi fungsinya pada file implementasi (.cpp). Pengelolaan memori dilakukan secara dinamis menggunakan perintah alokasi untuk memesan ruang memori elemen baru dan dealokasi untuk menghapus memori elemen yang sudah tidak terpakai.

Dalam praktikum ini, konsep Singly Linked List diterapkan melalui beberapa operasi dasar atau primitif untuk memanipulasi isi list. Operasi tersebut meliputi pembuatan list kosong (createList), penambahan elemen (insert first, insert last, insert after), penghapusan elemen (delete first, delete last, delete after), serta penelusuran untuk menampilkan seluruh isi data di dalam list (printInfo).

## Guided 

### 1. list.h

```C++
#ifndef LIST_H
#define LIST_H
#define Nil NULL

#include <iostream>
using namespace std;

// deklarasi isi data struct mahasiswa
struct mahasiswa {
    string nama;
    string nim;
    int umur;
};

typedef mahasiswa dataMahasiswa;

// mendefinisikan alias address sebagai pointer ke struct node
typedef struct node *address;

struct node {
    dataMahasiswa isidata;
    address next;
};

struct linkedList {
    address first;
};

// semua function dan procedure yang akan dipakai
bool isEmpty(linkedList list);
void createList(linkedList &list);
address alokasi(string nama, string nim, int umur);
void dealokasi(address &node);
void printList(linkedList List);

void insertFirst(linkedList &List, address nodeBaru);
void insertAfter(linkedList &List, address nodeBaru, address Prev);
void insertLast(linkedList &List, address nodeBaru);

void delFirst(linkedList &List);
void delLast(linkedList &List);
void delAfter(linkedList &List, address nodeHapus, address nodePrev);

int nbList(linkedList List);
void deleteList(linkedList &List);

#endif
```
#### penjelasan guided 1
Kode list.h ini digunakan untuk menyimpan data mahasiswa berupa nama, NIM, dan umur. Di dalamnya juga terdapat deklarasi fungsi untuk menambah, menghapus, menampilkan, dan menghitung jumlah data mahasiswa dalam linked list.

### 2. list.cpp

```C++
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
```
#### penjelasan guided 2
Kode list.cpp ini berisi fungsi untuk mengelola linked list, mulai dari membuat list kosong, menambahkan data mahasiswa di awal, tengah, atau akhir, sampai menghapus data. Selain itu, ada fungsi untuk menampilkan data mahasiswa, menghitung jumlah node, dan menghapus seluruh isi list.

### 3. main.cpp

```C++
#include "list.h"

#include <iostream>
using namespace std;

int main() {
    linkedList List;
    address nodeA, nodeB, nodeC, nodeD, nodeE = Nil;

    createList(List);

    dataMahasiswa mhs;

    nodeA = alokasi("Dhimas", "2311102151", 20);
    nodeB = alokasi("Arvin", "2211110014", 21);
    nodeC = alokasi("Rizal", "2311110029", 20);
    nodeD = alokasi("Satrio", "2211102173", 21);
    nodeE = alokasi("Joshua", "2311102133", 21);

    insertFirst(List, nodeA);
    insertLast(List, nodeB);
    insertAfter(List, nodeC, nodeA);
    insertAfter(List, nodeD, nodeC);
    insertLast(List, nodeE);

    cout << "--- ISI LIST SETELAH DILAKUKAN INSERT ---" << endl;
    printList(List);
    cout << "jumlah node : " << nbList(List) << endl;
    cout << endl;

    delFirst(List);
    delLast(List);
    delAfter(List, nodeD, nodeC);

    cout << "--- ISI LIST SETELAH DILAKUKAN DELETE ---" << endl;
    printList(List);
    cout << "jumlah node : " << nbList(List) << endl;
    cout << endl;

    deleteList(List);

    cout << "--- ISI LIST SETELAH DILAKUKAN HAPUS LIST ---" << endl;
    printList(List);
    cout << "jumlah node : " << nbList(List) << endl;
    cout << endl;

    return 0;
}
```
#### penjelasan guided 3
Kode main.cpp ini digunakan untuk menjalankan program linked list. Pertama, program membuat list kosong lalu memasukkan lima data mahasiswa menggunakan fungsi insert. Setelah itu, program menampilkan data dan jumlah node, kemudian menghapus beberapa data menggunakan fungsi delete. Terakhir, semua data dihapus menggunakan deleteList(), lalu program menampilkan bahwa list sudah kosong.


## Unguided 

### 1. Implementasi ADT Singly Linked List
```C++
singlylist.h
#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#define Nil NULL

typedef int infotype;
typedef struct ElmList *address;

struct ElmList {
    infotype info;
    address next;
};

struct List {
    address First;
};

void createList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertFirst(List &L, address P);

#endif // SINGLYLIST_H
----------------------------------------------------------------------------------------------------------------------------------------
singlylist.cpp
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
--------------------------------------------------------------------------------------------------------------------------------------
main.cpp
#include <iostream>
#include "Singlylist.h"
using namespace std;

int main()
{
    List L;
    address P1, P2, P3, P4, P5 = Nil;
    createList(L);

    P1 = alokasi(2);
    insertFirst(L, P1);

    P2 = alokasi(0);
    insertFirst(L, P2);

    P3 = alokasi(8);
    insertFirst(L, P3);

    P4 = alokasi(12);
    insertFirst(L, P4);

    P5 = alokasi(9);
    insertFirst(L, P5);

    printInfo(L);

    return 0;
}
```
#### Output Unguided 1 :

##### Output 1
![Output Unguided 1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%204%20-%20Modul%203/Unguided%201/Output-1.png)

#### penjelasan unguided 1 
Di bagian Singlylist.h saya membuat struct ElmList dan List beserta deklarasi fungsi createList(), alokasi(), dealokasi(), printInfo(), dan insertFirst().
Di file Singlylist.cpp berisi isi dari fungsi-fungsi tersebut. createList() digunakan untuk membuat list kosong, alokasi() untuk membuat elemen baru, dan dealokasi() untuk menghapus memori elemen. Lalu insertFirst() digunakan untuk memasukkan elemen di awal list dan printInfo() untuk menampilkan isi list. 
Kemudian di main.cpp saya membuat list L dan memasukkan angka 2, 0, 8, 12, dan 9 menggunakan alokasi() dan insertFirst(). Setelah semuanya dimasukkan, isi list ditampilkan dengan printInfo() sehingga menghasilkan output 9 12 8 0 2.

### 2. Penghapusan Node pada Singly Linked List

```C++
singlylist.h
#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#define Nil NULL

typedef int infotype;
typedef struct ElmList *address;

struct ElmList {
    infotype info;
    address next;
};

struct List {
    address First;
};

void createList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertFirst(List &L, address P);

void deleteFirst(List &L, address &P);
void deleteLast(List &L, address &P);
void deleteAfter(List &L, address Prec, address &P);
int nbList(List L);
void deleteList(List &L);

#endif
--------------------------------------------------------------------------------------------------------------------------------------
singlylist.cpp
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
-------------------------------------------------------------------------------------------------------------------------------------
main.cpp
#include <iostream>
#include "Singlylist.h"
using namespace std;

int main()
{
    List L;
    address P1, P2, P3, P4, P5 = Nil;
    address Pdel = Nil;
    createList(L);

    P1 = alokasi(2);
    insertFirst(L, P1);

    P2 = alokasi(0);
    insertFirst(L, P2);

    P3 = alokasi(8);
    insertFirst(L, P3);

    P4 = alokasi(12);
    insertFirst(L, P4);

    P5 = alokasi(9);
    insertFirst(L, P5);

    deleteFirst(L, Pdel);
    dealokasi(Pdel);

    deleteLast(L, Pdel);
    dealokasi(Pdel);

    deleteAfter(L, P4, Pdel);
    dealokasi(Pdel);

    printInfo(L);
    cout << "Jumlah node : " << nbList(L) << endl;

    deleteList(L);
    cout << "\n- List Berhasil Terhapus -" << endl;
    cout << "Jumlah node : " << nbList(L) << endl;

    return 0;
}
```
### Output Unguided 2 :

##### Output 1
![Output Unguided 2](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%204%20-%20Modul%203/Unguided%202/Output-1.png)

#### penjelasan unguided 2
Di bagian Singlylist.h saya menambahkan deklarasi fungsi deleteFirst(), deleteLast(), deleteAfter(), nbList(), dan deleteList().
Di file Singlylist.cpp berisi isi dari fungsi-fungsi tersebut. deleteFirst() digunakan untuk menghapus elemen pertama, deleteLast() untuk menghapus elemen terakhir, dan deleteAfter() untuk menghapus elemen setelah node tertentu. Lalu nbList() digunakan untuk menghitung jumlah node di dalam list dan deleteList() untuk menghapus semua node sampai list kosong.
Kemudian di main.cpp, setelah data dimasukkan, saya menghapus node 9 dengan deleteFirst(), node 2 dengan deleteLast(), dan node 8 dengan deleteAfter(). Setelah itu sisa isi list dan jumlah nodenya ditampilkan sehingga muncul 12 0 dan jumlah node 2. Terakhir, semua node dihapus menggunakan deleteList() sehingga jumlah nodenya menjadi 0. 


## Kesimpulan
Setelah melakukan praktikum pada Modul 4, saya jadi lebih memahami konsep Singly Linked List dan cara penerapannya menggunakan pointer dalam C++. Penggunaan linked list membuat penyimpanan data menjadi lebih dinamis dan fleksibel dibandingkan array, serta program tetap terstruktur karena deklarasi ADT dan implementasi fungsinya dipisahkan ke dalam file header (.h) dan file .cpp.
Dari praktikum ini saya juga mempraktikkan berbagai operasi dasar pada Singly Linked List melalui program yang dibuat. Operasi tersebut meliputi pembuatan list kosong (createList), alokasi dan dealokasi memori, penambahan elemen di awal list (insertFirst), penghapusan elemen (deleteFirst, deleteLast, deleteAfter, dan deleteList), hingga penelusuran untuk menampilkan isi dan menghitung jumlah node (printInfo dan nbList). Dengan latihan tersebut, saya jadi lebih paham bagaimana cara kerja pointer dalam menghubungkan, menambah, dan menghapus elemen di dalam list.

## Referensi
[1] Mbejo, M. T., & Nopa, L. A. (2025). Analisis Struktur Data Linked List Dalam Pengolahan Data Mahasiswa. Jurnal Sains Informatika Terapan (JSIT), 4(2), 441–444.
[2] Setiyawan, R. D., Hermawan, D., Abdillah, A. F., Mujayanah, A., & Vindua, R. (2024). Penggunaan Struktur Data Stack dalam Pemrograman C++ dengan Pendekatan Array dan Linked List. JUTECH: Journal Education and Technology, 5(2), 484–498.