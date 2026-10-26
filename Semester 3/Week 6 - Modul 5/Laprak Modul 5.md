# <h1 align="center">Laporan Praktikum Modul 5 - SINGLY LINKED LIST (BAGIAN KEDUA)   </h1>
<p align="center">Hilman Aniq Almunawa - 109082500027</p>

## Dasar Teori
Singly linked list merupakan struktur data linear yang terdiri dari kumpulan node yang saling terhubung menggunakan pointer. Setiap node memiliki dua bagian, yaitu info untuk menyimpan data dan next untuk menunjuk ke node berikutnya. Node terakhir memiliki pointer next bernilai Nil, sedangkan list kosong ditandai dengan First bernilai Nil (Morin, 2013).

Operasi dasar singly linked list meliputi pembuatan list, alokasi memori, penyisipan elemen, dan penelusuran data. Pada Modul 5, operasi tersebut diimplementasikan menggunakan bahasa C++ melalui fungsi CreateList(), alokasi(), insertFirst(), dan printInfo(). Fungsi insertFirst() menempatkan node baru di awal list dengan mengatur pointer agar node baru terhubung dengan elemen pertama sebelumnya.

Selain itu, searching digunakan untuk mencari elemen berdasarkan nilai yang tersimpan pada node. Fungsi findElm() menelusuri list dari elemen pertama hingga nilai yang dicari ditemukan atau mencapai akhir list. Penelusuran juga dapat digunakan untuk menghitung total nilai seluruh elemen.

## Guided 

### 1. listBuah.h

```C++
  //Header guard digunakan untuk mencegah file header yang sama 
  //di-include lebih dari sekali dalam satu program.
  #ifndef LISTBUAH_H
  #define LISTBUAH_H
  #define Nil NULL

  #include<iostream>
  using namespace std;

  struct buah{
      string nama;
      int jumlah; 
      float harga;
  };

  typedef buah dataBuah; //Memberikan nama alias databuah untuk struct buah.

  typedef struct node *address; //Mendefinisikan alias address sebagai pointer ke struct node

  struct node{ // node untuk isi dari linked listnya, isi setiap node adalah data & pointer next
      dataBuah isidata;
      address next;
  };

  struct linkedlist{ //ini linked list nya
      address first;
  };

  //semua function & prosedur yang akan dipakai
  //Materi modul 4
  bool isEmpty(linkedlist List);
  void createList(linkedlist &List);
  address alokasi(string nama, int jumlah, float harga);
  void dealokasi(address &node);
  void printList(linkedlist List);
  void insertFirst(linkedlist &List, address nodeBaru);
  void insertAfter(linkedlist &List, address nodeBaru, address Prev);
  void insertLast(linkedlist &List, address nodeBaru);
  void delFirst(linkedlist &List);
  void delLast(linkedlist &List);
  void delAfter(linkedlist &List, address nodeHapus, address nodePrev);
  int nbList(linkedlist List);
  void deleteList(linkedlist &List);

  //materi modul 5 (update)
  void updateFirst(linkedlist List);
  void updateLast(linkedlist List);
  void updateAfter(linkedlist List, address prev);

  //materi modul 5 (searching)
  void FindNodeByData(linkedlist list, string data);
  void FindNodeByAddress(linkedlist list, address node);
  void FindNodeByRange(linkedlist list, float hargaAwal, float HargaAkhir);

  #endif
```
#### penjelasan guided 1
Kode `list.h` ini digunakan untuk menyimpan data buah berupa nama, jumlah, dan harga. Di dalamnya juga terdapat deklarasi fungsi untuk menambah, menghapus, menampilkan, menghitung jumlah data, memperbarui, dan mencari data buah berdasarkan nama, alamat node, serta rentang harga dalam linked list.

### 2. listBuah.cpp

```C++
#include "listBuah.h"
#include <iostream>
using namespace std;

//fungsi untuk cek apakah list kosong atau tidak
bool isEmpty(linkedlist List) {
    if(List.first == Nil){
        return true; 
    } else {
        return false;
    }
}

//pembuatan linked list kosong
void createList(linkedlist &List) {
    List.first = Nil;
}

//pembuatan node baru dengan menerapkan manajemen memori
address alokasi(string nama, int jumlah, float harga) { 
    address nodeBaru = new node; 
    nodeBaru->isidata.nama = nama;
    nodeBaru->isidata.jumlah = jumlah; 
    nodeBaru->isidata.harga = harga;
    nodeBaru->next = Nil;
    return nodeBaru;
}

//penghapusan node dengan menerapkan manajemen memori
void dealokasi(address &node) {
    node->next = Nil;
    delete node;
}

//prosedur-prosedur untuk insert / menambahkan node baru kedalam list
void insertFirst(linkedlist &List, address nodeBaru) {
    nodeBaru->next = List.first; 
    List.first = nodeBaru;
}

void insertAfter(linkedlist &List, address nodeBaru, address Prev) {
    if (Prev != Nil) {
        nodeBaru->next = Prev->next;
        Prev->next = nodeBaru;
    } else {
        cout << "Node sebelumnya tidak valid!" << endl;
    }
}

void insertLast(linkedlist &List, address nodeBaru) {
    if (isEmpty(List)) {
        List.first = nodeBaru;
    } else {
        address nodeBantu = List.first;
        while (nodeBantu->next != Nil) {
            nodeBantu = nodeBantu->next;
        }
        nodeBantu->next = nodeBaru;
    }
}

//prosedur-prosedur untuk delete / menghapus node yang ada didalam list
void delFirst(linkedlist &List){
    address nodeHapus;
    if (isEmpty(List) == false) {
        nodeHapus = List.first;
        List.first = List.first->next;
        nodeHapus->next = Nil;
        dealokasi(nodeHapus);
        cout << "Node pertama berhasil terhapus!" << endl;
    } else {
        cout << "List kosong!" << endl;
    }
}

void delLast(linkedlist &List){
    address nodeHapus, nodePrev;
    if(isEmpty(List) == false){
        nodeHapus = List.first;
        if(nodeHapus->next == Nil){
            List.first->next = Nil;
            dealokasi(nodeHapus);
        } else { 
            while(nodeHapus->next != Nil){
                nodePrev = nodeHapus; 
                nodeHapus = nodeHapus->next;
            }
            nodePrev->next = Nil; 
            dealokasi(nodeHapus);
        }
        cout << "Node terakhir berhasil terhapus!" << endl;
    } else {
        cout << "list kosong" << endl;
    }
}

void delAfter(linkedlist &List, address nodeHapus, address nodePrev){
    if(isEmpty(List) == true){
        cout << "List kosong!" << endl;
    } else { //jika list tidak kosong
        if (nodePrev != Nil && nodePrev->next != Nil) { 
            nodeHapus = nodePrev->next;       
            nodePrev->next = nodeHapus->next;  
            nodeHapus->next = Nil;         
            dealokasi(nodeHapus);
            cout << "Node setelah node " << nodePrev->isidata.nama << " berhasil terhapus!" << endl;
        } else {
            cout << "Node sebelumnya (prev) tidak valid!" << endl;
        }
    }
}

//prosedur untuk menampilkan isi list
void printList(linkedlist List) {
    if (isEmpty(List)) {
        cout << "List kosong." << endl;
    } else {
        address nodeBantu = List.first;
        while (nodeBantu != Nil) { 
            cout << "Nama Buah : " << nodeBantu->isidata.nama << ", Jumlah : " << nodeBantu->isidata.jumlah << ", Harga : " << nodeBantu->isidata.harga << endl;
            nodeBantu = nodeBantu->next;
        }
    }
}

//function untuk menampilkan jumlah node didalam list
int nbList(linkedlist List) {
    int count = 0;
    address nodeBantu = List.first;
    while (nodeBantu != Nil) {
        count++;
        nodeBantu = nodeBantu->next; 
    }
    return count;
}

//prosedur untuk menghapus list (menghapus semua node didalam list)
void deleteList(linkedlist &List){
    address nodeBantu, nodeHapus;
    nodeBantu = List.first;
    while(nodeBantu != Nil){
        nodeHapus = nodeBantu;
        nodeBantu = nodeBantu->next;
        dealokasi(nodeHapus); 
    }
    List.first = Nil; 
    cout << "List berhasil terhapus!" << endl;
}

/*----- MATERI PERTEMUAN 5 - SINGLY LINKED LIST (BAGIAN KEDUA) - UPDATE -----*/
//prosedur-prosedur untuk melakukan update data node
void updateFirst(linkedlist List){
    if(isEmpty(List) == true){
        cout << "List kosong!" << endl;
    } else {
        cout << "Masukkan update data node pertama : " << endl;
        cout << "Nama buah : ";
        cin >> List.first->isidata.nama;
        cout << "Jumlah : ";
        cin >> List.first->isidata.jumlah;
        cout << "Harga : ";
        cin >> List.first->isidata.harga;
        cout << "Data Berhasil Diupdate!" << endl;
        cout << endl;
    }
}

void updateLast(linkedlist List){
    if (isEmpty(List) == true) {
        cout << "List Kosong!" << endl;
    } else {
        address nodeBantu = List.first;
        while (nodeBantu->next != Nil) {
            nodeBantu = nodeBantu->next;
        }
        cout << "masukkan update data node terakhir : " << endl;
        cout << "Nama buah : ";
        cin >> nodeBantu->isidata.nama;
        cout << "Jumlah : ";
        cin >> nodeBantu->isidata.jumlah;
        cout << "Harga : ";
        cin >> nodeBantu->isidata.harga;
        cout << "Data Berhasil Diupdate!" << endl;
        cout << endl;
    }
}

void updateAfter(linkedlist List, address nodePrev){
    if(isEmpty(List) == true){
        cout << "List kosong!" << endl;
    } else {
        if (nodePrev != Nil && nodePrev->next != Nil){
            address nodeBantu = nodePrev->next;
            cout << "masukkan update data node setelah node " << nodePrev->isidata.nama << " : " << endl;
            cout << "Nama buah : ";
            cin >> nodeBantu->isidata.nama;
            cout << "Jumlah : ";
            cin >> nodeBantu->isidata.jumlah;
            cout << "Harga : ";
            cin >> nodeBantu->isidata.harga;
            cout << "Data Berhasil Diupdate!" << endl;
            cout << endl;
        } else {
            cout << "Node sebelumnya (prev) tidak valid!" << endl;
        }
    }
}

/*----- MATERI PERTEMUAN 5 - SINGLY LINKED LIST (BAGIAN KEDUA) - SEARCHING -----*/
//prosedur-prosedur untuk searching data
//prosedur untuk mencari node berdasarkan data
void FindNodeByData(linkedlist list, string data){
    if(isEmpty(list) == true){
        cout << "List kosong!" << endl;
    } else {
        address nodeBantu = list.first;
        int posisi = 0;
        bool found = false;
        while(nodeBantu != Nil){
            posisi++;
            if(nodeBantu->isidata.nama == data){
                cout << "Data " << data << " ditemukan pada posisi ke-" << posisi << "!" << endl;
                cout << "Nama Buah : " << nodeBantu->isidata.nama << ", Jumlah : " << nodeBantu->isidata.jumlah << ", Harga : " << nodeBantu->isidata.harga << endl;
                found = true;
                break;
            }
            nodeBantu = nodeBantu->next;
        }
        if(found == false){
            cout << "Node dengan data " << data << " tidak ditemukan!" << endl;
        }
    }
    cout << endl;
}

//prosedur untuk mencari node berdasarkan alamat node
void FindNodeByAddress(linkedlist list, address node) {
    if(isEmpty(list) == true) {
        cout << "List kosong!" << endl;
    } else {
        address nodeBantu = list.first;
        int posisi = 0;
        bool found = false;
        while (nodeBantu != Nil) {
            posisi++;
            if(nodeBantu == node) {
                cout << "Node ditemukan pada posisi ke-" << posisi << "!" << endl;
                cout << "Alamat node : " << nodeBantu << endl;
                cout << "Nama Buah : " << nodeBantu->isidata.nama << ", Jumlah : " << nodeBantu->isidata.jumlah << ", Harga : " << nodeBantu->isidata.harga << endl;
                found = true;
                break;
            }
            nodeBantu = nodeBantu->next;
        }
        if(found == false) {
            cout << "Node dengan alamat " << node << " tidak ditemukan dalam list!" << endl;
        }
    }
    cout << endl;
}

//prosedur untuk mencari node berdasarkan range data (range harga)
void FindNodeByRange(linkedlist list, float hargaAwal, float hargaAkhir) {
    if(isEmpty(list) == true) {
        cout << "List kosong!" << endl;
    } else {
        address nodeBantu = list.first;
        int posisi = 0;
        bool found = false;
        cout << "--- Buah dalam range harga " << hargaAwal << " - " << hargaAkhir << " ---" << endl;
        cout << "-------------------------------------------" << endl;
        while (nodeBantu != Nil) {
            posisi++;
            float harga = nodeBantu->isidata.harga;
            if(harga >= hargaAwal && harga <= hargaAkhir) {
                cout << "Data ditemukan pada posisi ke-" << posisi << " :" << endl;
                cout << "Nama Buah : " << nodeBantu->isidata.nama << ", Jumlah : " << nodeBantu->isidata.jumlah << ", Harga : " << nodeBantu->isidata.harga << endl;
                cout << "-------------------------------------------" << endl;
                found = true;
            }
            nodeBantu = nodeBantu->next;
        }
        if(found == false) {
            cout << "Tidak ada data buah dalam range harga tersebut!" << endl;
            cout << "-------------------------------------------" << endl;
        }
    }
    cout << endl;
}
```
#### penjelasan guided 2
Kode `listBuah.cpp` ini berisi implementasi fungsi untuk mengelola data buah dalam linked list, mulai dari membuat list, menambah, menghapus, dan menampilkan data. Selain itu, terdapat fungsi untuk menghitung jumlah node, memperbarui data buah, serta mencari data berdasarkan nama, alamat node, dan rentang harga.

### 3. main.cpp

```C++
#include "listBuah.h"

#include<iostream>
using namespace std;

int main(){
    linkedlist List;
    address nodeA, nodeB, nodeC, nodeD, nodeE = Nil;
    createList(List);

    dataBuah dtBuah;

    nodeA = alokasi("Jeruk", 100, 3000);
    nodeB = alokasi("Apel", 75, 4000);
    nodeC = alokasi("Pir", 87, 5000);
    nodeD = alokasi("Semangka", 43, 11500);
    nodeE = alokasi("Durian", 15, 31450);

    insertFirst(List, nodeA);
    insertLast(List, nodeB);
    insertAfter(List, nodeC, nodeA);
    insertAfter(List, nodeD, nodeC);
    insertLast(List, nodeE);

    cout << "--- ISI LIST SETELAH DILAKUKAN INSERT ---" << endl;
    printList(List);
    cout << "jumlah node : " << nbList(List) << endl;
    cout << endl;

    updateFirst(List);
    printList(List);
    updateLast(List);
    printList(List);
    updateAfter(List, nodeD);
    printList(List);
    
    cout << "--- ISI LIST SETELAH DILAKUKAN UPDATE ---" << endl;
    printList(List);
    cout << "jumlah node : " << nbList(List) << endl;
    cout << endl;
    
    FindNodeByData(List, "kelapa");
    FindNodeByAddress(List, nodeC);
    FindNodeByRange(List, 5000, 10000);
    
    delFirst(List);
    printList(List);
    delLast(List);
    printList(List);
    delAfter(List, nodeD, nodeC);
    printList(List);

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
Kode `main.cpp` ini digunakan untuk menjalankan program linked list dengan memasukkan lima data buah, yaitu Jeruk, Apel, Pir, Semangka, dan Durian. Program melakukan proses penambahan, pembaruan, dan pencarian data berdasarkan nama, alamat node, serta rentang harga. Setelah itu, program melakukan penghapusan data dan menampilkan isi list beserta jumlah node pada setiap proses.


## Unguided 

### 1. Singlylist
```C++
Singlylist.h
#ifndef SINGLYLIST_H_INCLUDED
#define SINGLYLIST_H_INCLUDED

#include <iostream>
#define Nil NULL

using namespace std;

typedef int infotype;
typedef struct ElmList *address;

struct ElmList {
    infotype info;
    address next;
};

struct List {
    address First;
};

void CreateList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void insertFirst(List &L, address P);
void printInfo(List L);
address findElm(List L, infotype x);
int sumInfo(List L);
void deleteList(List &L);

#endif
--------------------------------------------------------------------------------------------------------------------------------------
Singlylist.cpp
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
-----------------------------------------------------------------------------------------------------------------------------------
main.cpp
#include <iostream>
#include "Singlylist.h"

using namespace std;

int main() {
    List L;
    address P1, P2, P3, P4, P5;

    CreateList(L);

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

    deleteList(L);

    return 0;
}
```
#### Output Unguided 1 :

##### Output 1
![Output Unguided 1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%206%20-%20Modul%205/Unguided%201/Output-1.png)

#### penjelasan unguided 1 
Di bagian Singlylist.h, saya membuat struktur ElmList dan List, serta deklarasi fungsi yang digunakan dalam linked list. Pada Singlylist.cpp, saya mengisi fungsi-fungsi tersebut, mulai dari membuat list kosong, mengalokasikan memori, menambahkan elemen di awal list, sampai menampilkan isi list.

Kemudian, di main.cpp, saya memasukkan angka 2, 0, 8, 12, dan 9 menggunakan alokasi() dan insertFirst(). Karena setiap elemen dimasukkan di awal list, urutan datanya menjadi 9, 12, 8, 0, 2. Setelah ditampilkan menggunakan printInfo(), seluruh elemen dihapus dengan deleteList() agar memori yang digunakan dikembalikan.

### 2. Pencarian Elemen dengan Info 8

```C++
#include <iostream>
#include "Singlylist.h"

using namespace std;

int main() {
    List L;
    address P1, P2, P3, P4, P5;
    address cari;

    CreateList(L);

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

    cari = findElm(L, 8);

    if (cari != Nil) {
        cout << cari->info << " ditemukan dalam list" << endl;
    } else {
        cout << "8 tidak ditemukan dalam list" << endl;
    }

    deleteList(L);

    return 0;
}
```
### Output Unguided 2 :

##### Output 1
![Output Unguided 2](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%206%20-%20Modul%205/Unguided%202/Output-1.png)

#### penjelasan unguided 2
Pada soal ini, saya menggunakan kembali file Singlylist.h dan Singlylist.cpp dari soal sebelumnya. Kedua file tersebut tidak diubah karena fungsi findElm() sudah dibuat di dalamnya. Saya hanya mengubah isi main.cpp untuk menjalankan proses pencarian elemen dengan nilai 8. Hasilnya, program menampilkan pesan bahwa elemen 8 ditemukan dalam list.

### 3. Menghitung Total Info Seluruh Elemen

```C++
#include <iostream>
#include "Singlylist.h"

using namespace std;

int main() {
    List L;
    address P1, P2, P3, P4, P5;

    CreateList(L);

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

    cout << "Total info dari kelima elemen adalah "
         << sumInfo(L) << endl;

    deleteList(L);

    return 0;
}
```
### Output Unguided 3 :

##### Output 1
![Output Unguided 3](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%206%20-%20Modul%205/Unguided%203/Output-1.png)

#### penjelasan unguided 3
Pada soal ini, saya tetap menggunakan file Singlylist.h dan Singlylist.cpp yang sama seperti soal sebelumnya. Saya hanya mengubah isi main.cpp untuk menjalankan fungsi sumInfo() dan menampilkan hasil penjumlahan seluruh elemen dalam list. Dari nilai 9, 12, 8, 0, dan 2, diperoleh total sebesar 31.

## Kesimpulan
Dari praktikum Modul 5 ini saya jadi lebih paham cara kerja singly linked list dengan pointer di C++. Saya belajar membuat list kosong, membuat elemen baru dengan `alokasi()`, dan memasukkannya ke awal list menggunakan `insertFirst()`. Karena elemen selalu dimasukkan di awal, urutan data yang ditampilkan menjadi 9, 12, 8, 0, 2.

Selain itu, saya mencoba pencarian menggunakan `findElm()` dan penjumlahan menggunakan `sumInfo()`. Hasilnya, elemen 8 berhasil ditemukan dan total seluruh info adalah 31. Saya juga menggunakan `deleteList()` untuk menghapus semua elemen dan mengembalikan memori yang digunakan setelah program selesai.

## Referensi
[1] Morin, P. (2013). Open Data Structures: An Introduction. Bab 3: Linked Lists. https://open.umn.edu/opentextbooks/textbooks/171