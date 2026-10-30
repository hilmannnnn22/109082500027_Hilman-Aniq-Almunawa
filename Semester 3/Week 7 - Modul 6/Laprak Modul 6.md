# <h1 align="center">Laporan Praktikum Modul 6 - DOUBLY LINKED LIST (BAGIAN PERTAMA)</h1>
<p align="center">Hilman Aniq Almunawa - 109082500027</p>

## Dasar Teori
Doubly linked list merupakan struktur data linear yang terdiri dari kumpulan node yang saling terhubung menggunakan pointer. Setiap node memiliki tiga bagian, yaitu info untuk menyimpan data, next untuk menunjuk node berikutnya, dan prev untuk menunjuk node sebelumnya. Struktur ini juga memiliki pointer first dan last yang menunjuk node pertama dan terakhir sehingga data dapat ditelusuri dari dua arah (Morin, 2013).

Operasi doubly linked list meliputi penyisipan (insert), penghapusan (delete), pencarian (searching), pembaruan (update), dan penampilan data (view). Operasi tersebut dilakukan dengan mengatur hubungan antarpointer agar struktur list tetap terhubung dengan benar. Implementasinya dapat menggunakan konsep Abstract Data Type (ADT), yaitu pendefinisian data beserta operasi yang dapat dilakukan terhadapnya (Goodrich et al., 2014).

## Guided 

### 1. Doublylist.h

```C++
#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H
#include <string>
using namespace std;

struct barang {
    string kode;
    string nama;
    int tahun;
};

using infotype = barang;
struct elmList;
using address = elmList*;
struct elmList { infotype info; address prev; address next; };
struct list { address first; address last; };

void CreateList(list &L);
address alokasi(const infotype &x);
void dealokasi(address &P);
void insertFirst(list &L, address P);
void insertLast(list &L, address P);
void insertAfter(list &L, address R, address P);
void insertBefore(list &L, address R, address P);
void printInfo(const list &L);
address findElm(const list &L, const string &kode);
void deleteFirst(list &L, address &P);
void deleteLast(list &L, address &P);
void deleteAfter(list &L, address prec, address &P);
void deleteBefore(list &L, address succ, address &P);

#endif
```
#### penjelasan guided 1
Pada file Doublylist.h, terdapat struktur data barang yang berisi kode, nama, dan tahun. File ini juga berisi deklarasi fungsi yang digunakan untuk mengelola data barang, seperti menambahkan, menghapus, mencari, dan menampilkan data menggunakan doubly linked list.

### 2. Doublylist.cpp

```C++
#include <iostream>
#include <iomanip>
#include "Doublylist.h"
using namespace std;

void CreateList(list &L) { L.first = nullptr; L.last = nullptr; }

address alokasi(const infotype &x) {
    address P = new elmList;
    P->info = x;
    P->prev = nullptr; P->next = nullptr;
    return P;
}

void dealokasi(address &P) {
    if (P != nullptr) { delete P; P = nullptr; }
}

void insertFirst(list &L, address P) {
    if (!P) return;
    P->prev = nullptr; P->next = L.first;
    if (L.first) L.first->prev = P;
    else L.last = P;
    L.first = P;
}

void insertLast(list &L, address P) {
    if (!P) return;
    P->next = nullptr; P->prev = L.last;
    if (L.last) L.last->next = P;
    else L.first = P;
    L.last = P;
}

// Menyisipkan P tepat setelah R.
void insertAfter(list &L, address R, address P) {
    if (!R || !P) return;
    P->prev = R;
    P->next = R->next;
    if (R->next) R->next->prev = P;
    else L.last = P;
    R->next = P;
}

// Menyisipkan P tepat sebelum R.
void insertBefore(list &L, address R, address P) {
    if (!R || !P) return;
    P->next = R;
    P->prev = R->prev;
    if (R->prev) R->prev->next = P;
    else L.first = P;
    R->prev = P;
}

void printInfo(const list &L) {
    address P = L.first;
    if (!P) {
        cout << "List kosong.\n";
        return;
    }

    cout << "+----------+------------+-------+\n";
    cout << "| Kode     | Nama       | Tahun |\n";
    cout << "+----------+------------+-------+\n";
    while (P) {
        cout << "| " << left << setw(6) << P->info.kode
             << " | " << left << setw(10) << P->info.nama
             << " | " << setw(5) << P->info.tahun << " |\n";
        P = P->next;
    }
    cout << "+----------+------------+-------+\n";
}

address findElm(const list &L, const string &kode) {
    address P = L.first;
    while (P) {
        if (P->info.kode == kode) return P;
        P = P->next;
    }
    return nullptr;
}

void deleteFirst(list &L, address &P) {
    if (!L.first) { P = nullptr; return; }
    P = L.first; L.first = P->next;
    if (L.first) L.first->prev = nullptr;
    else L.last = nullptr;
    P->prev = P->next = nullptr;
}

void deleteLast(list &L, address &P) {
    if (!L.last) { P = nullptr; return; }
    P = L.last; L.last = P->prev;
    if (L.last) L.last->next = nullptr;
    else L.first = nullptr;
    P->prev = P->next = nullptr;
}

// Menghapus node tepat setelah prec.
void deleteAfter(list &L, address prec, address &P) {
    P = nullptr;
    if (!prec || !prec->next) return;
    P = prec->next; prec->next = P->next;
    if (P->next) P->next->prev = prec;
    else L.last = prec;
    P->prev = P->next = nullptr;
}

// Menghapus node tepat sebelum succ.
void deleteBefore(list &L, address succ, address &P) {
    P = nullptr;
    if (!succ || !succ->prev) return;
    P = succ->prev; succ->prev = P->prev;
    if (P->prev) P->prev->next = succ;
    else L.first = succ;
    P->prev = P->next = nullptr;
}
```
#### penjelasan guided 2
Pada file Doublylist.cpp, terdapat implementasi dari fungsi-fungsi yang sudah dideklarasikan sebelumnya. Fungsi tersebut digunakan untuk mengatur proses penambahan, penghapusan, pencarian, dan penampilan data barang. Setiap node dihubungkan menggunakan pointer prev dan next sehingga data bisa ditelusuri ke depan maupun ke belakang.

### 3. main.cpp

```C++
#include <iostream>
#include "Doublylist.h"
using namespace std;

void tampilkanJudul(const string &judul) {
    cout << "\n====================================\n";
    cout << judul << '\n';
    cout << "====================================\n";
}

int main() {
    list L;
    CreateList(L);

    infotype barang1 = {"B001", "Laptop", 2020};
    infotype barang2 = {"B002", "Monitor", 2022};
    infotype barang3 = {"B003", "Keyboard", 2023};
    infotype barang4 = {"B004", "Mouse", 2024};

    address P = nullptr;
    address posisi = nullptr;

    cout << "DEMONSTRASI OPERASI DOUBLY LINKED LIST\n";
    cout << "Data yang digunakan: B001, B002, B003, dan B004\n";

    tampilkanJudul("[1] INSERT FIRST");
    cout << "Menambahkan B002 di awal list.\n";
    insertFirst(L, alokasi(barang2));
    printInfo(L);

    tampilkanJudul("[2] INSERT LAST");
    cout << "Menambahkan B004 di akhir list.\n";
    insertLast(L, alokasi(barang4));
    printInfo(L);

    tampilkanJudul("[3] INSERT AFTER");
    cout << "Menambahkan B003 setelah B002.\n";
    posisi = findElm(L, "B002");
    if (posisi != nullptr) {
        insertAfter(L, posisi, alokasi(barang3));
    }
    printInfo(L);

    tampilkanJudul("[4] INSERT BEFORE");
    cout << "Menambahkan B001 sebelum B002.\n";
    posisi = findElm(L, "B002");
    if (posisi != nullptr) {
        insertBefore(L, posisi, alokasi(barang1));
    }
    printInfo(L);

    tampilkanJudul("[5] FIND ELEMENT");
    cout << "Mencari barang dengan kode B003.\n";
    posisi = findElm(L, "B003");
    if (posisi != nullptr) {
        cout << "Data ditemukan:\n";
        cout << "  Kode  : " << posisi->info.kode << '\n';
        cout << "  Nama  : " << posisi->info.nama << '\n';
        cout << "  Tahun : " << posisi->info.tahun << '\n';
    } else {
        cout << "Data tidak ditemukan.\n";
    }

    tampilkanJudul("[6] DELETE AFTER");
    cout << "Menghapus node setelah B002 (B003).\n";
    posisi = findElm(L, "B002");
    deleteAfter(L, posisi, P);
    if (P != nullptr) {
        cout << "Node yang dihapus: " << P->info.kode << '\n';
        dealokasi(P);
    } else {
        cout << "Tidak ada node yang dapat dihapus.\n";
    }
    printInfo(L);

    tampilkanJudul("[7] DELETE BEFORE");
    cout << "Menghapus node sebelum B004 (B002).\n";
    posisi = findElm(L, "B004");
    deleteBefore(L, posisi, P);
    if (P != nullptr) {
        cout << "Node yang dihapus: " << P->info.kode << '\n';
        dealokasi(P);
    } else {
        cout << "Tidak ada node yang dapat dihapus.\n";
    }
    printInfo(L);

    tampilkanJudul("[8] DELETE FIRST");
    cout << "Menghapus node pertama.\n";
    deleteFirst(L, P);
    if (P != nullptr) {
        cout << "Node yang dihapus: " << P->info.kode << '\n';
        dealokasi(P);
    } else {
        cout << "List kosong, tidak ada node yang dapat dihapus.\n";
    }
    printInfo(L);

    tampilkanJudul("[9] DELETE LAST");
    cout << "Menghapus node terakhir.\n";
    deleteLast(L, P);
    if (P != nullptr) {
        cout << "Node yang dihapus: " << P->info.kode << '\n';
        dealokasi(P);
    } else {
        cout << "List kosong, tidak ada node yang dapat dihapus.\n";
    }
    printInfo(L);

    tampilkanJudul("SELESAI");
    cout << "Kondisi list: "
         << ((L.first == nullptr && L.last == nullptr)
             ? "kosong.\n" : "masih berisi data.\n");

    return 0;
}
```
#### penjelasan guided 3
Pada file main.cpp, program dijalankan untuk mencoba berbagai operasi pada doubly linked list menggunakan data barang berupa laptop, monitor, keyboard, dan mouse. Operasi yang dilakukan meliputi penambahan data di awal dan akhir list, penyisipan sebelum dan sesudah node tertentu, pencarian data berdasarkan kode, serta penghapusan node. Setiap proses ditampilkan menggunakan printInfo() agar perubahan isi list bisa dilihat secara langsung.

## Unguided 

### 1. Doublylist no 1
```C++
Doublylist.h
#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H

#include <iostream>
#include <string>

using namespace std;

typedef struct kendaraan {
    string nopol;
    string warna;
    int thnBuat;
} infotype;

typedef struct elmList *address;

typedef struct elmList {
    infotype info;
    address next;
    address prev;
} elmList;

typedef struct {
    address First;
    address Last;
} List;

void CreateList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertLast(List &L, address P);

#endif

---------------------------------------------------------------------------------------------------------------------------------------

Doublylist.cpp
#include "Doublylist.h"

void CreateList(List &L) {
    L.First = nullptr;
    L.Last = nullptr;
}

address alokasi(infotype x) {
    address P = new elmList;

    P->info = x;
    P->next = nullptr;
    P->prev = nullptr;

    return P;
}

void dealokasi(address &P) {
    delete P;
    P = nullptr;
}

void insertLast(List &L, address P) {
    if (L.First == nullptr) {
        L.First = P;
        L.Last = P;
    } else {
        P->prev = L.Last;
        L.Last->next = P;
        L.Last = P;
    }
}

void printInfo(List L) {
    address P = L.Last;

    while (P != nullptr) {
        cout << "no polisi : " << P->info.nopol << endl;
        cout << "warna     : " << P->info.warna << endl;
        cout << "tahun     : " << P->info.thnBuat << endl;

        P = P->prev;
    }
}

--------------------------------------------------------------------------------------------------------------------------------------

main.cpp
#include "Doublylist.h"

int main() {
    List L;
    infotype x;
    address P;
    address Q;
    bool ditemukan;

    CreateList(L);

    for (int i = 1; i <= 4; i++) {
        cout << "masukkan nomor polisi: ";
        cin >> x.nopol;

        cout << "masukkan warna kendaraan: ";
        cin >> x.warna;

        cout << "masukkan tahun kendaraan: ";
        cin >> x.thnBuat;

        ditemukan = false;
        Q = L.First;

        while (Q != nullptr) {
            if (Q->info.nopol == x.nopol) {
                ditemukan = true;
                break;
            }

            Q = Q->next;
        }

        if (ditemukan) {
            cout << "nomor polisi sudah terdaftar" << endl;
        } else {
            P = alokasi(x);
            insertLast(L, P);
        }

        cout << endl;
    }

    cout << "DATA LIST 1" << endl;
    printInfo(L);

    return 0;
}
```
#### Output Unguided 1 :
##### Output 1
![Output Unguided 1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/ca67ec001806bd1189700483b9e5128e564ebf10/Semester%203/Week%207%20-%20Modul%206/Unguided%201/Output-1.png)

#### penjelasan unguided 1 
Di kode Doublylist.h dibuat struktur data kendaraan yang berisi nomor polisi, warna, dan tahun pembuatan. Selain itu, dibuat juga pointer next dan prev untuk menghubungkan setiap node. Di sini juga dideklarasikan fungsi-fungsi yang akan digunakan pada program.

Di kode Doublylist.cpp dibuat isi dari fungsi-fungsi yang sudah dideklarasikan sebelumnya. Mulai dari membuat list kosong, membuat node baru, menghapus node, menambahkan data di bagian akhir, dan menampilkan data kendaraan dari belakang ke depan.

Di kode main.cpp dibuat program utama untuk menjalankan Doubly Linked List. Kita diminta memasukkan 4 data kendaraan, lalu setiap nomor polisi dicek terlebih dahulu. Kalau nomornya sudah ada, data tidak akan dimasukkan. Setelah semua data selesai diinput, data yang berhasil disimpan akan ditampilkan.

### 2. Doublylist no 2

```C++
Doublylist.h
#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H

#include <iostream>
#include <string>

using namespace std;

typedef struct kendaraan {
    string nopol;
    string warna;
    int thnBuat;
} infotype;

typedef struct ElmList *address;

typedef struct ElmList {
    infotype info;
    address next;
    address prev;
} ElmList;

typedef struct {
    address First;
    address Last;
} List;

void CreateList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertLast(List &L, address P);

// Kode tambahan buat jawab no. 2
address findElm(List L, infotype x);

#endif

-------------------------------------------------------------------------------------------------------------------------------------

Doublylist.cpp
#include "Doublylist.h"

void CreateList(List &L) {
    L.First = nullptr;
    L.Last = nullptr;
}

address alokasi(infotype x) {
    address P = new ElmList;

    P->info = x;
    P->next = nullptr;
    P->prev = nullptr;

    return P;
}

void dealokasi(address &P) {
    delete P;
    P = nullptr;
}

void insertLast(List &L, address P) {
    if (L.First == nullptr) {
        L.First = P;
        L.Last = P;
    } else {
        P->prev = L.Last;
        L.Last->next = P;
        L.Last = P;
    }
}

void printInfo(List L) {
    address P = L.Last;

    while (P != nullptr) {
        cout << "no polisi : " << P->info.nopol << endl;
        cout << "warna     : " << P->info.warna << endl;
        cout << "tahun     : " << P->info.thnBuat << endl;

        P = P->prev;
    }
}

// Kode tambahan buat jawab no. 2
address findElm(List L, infotype x) {
    address P = L.First;

    while (P != nullptr) {
        if (P->info.nopol == x.nopol) {
            return P;
        }

        P = P->next;
    }

    return nullptr;
}

-------------------------------------------------------------------------------------------------------------------------------------

main.cpp
#include "Doublylist.h"

int main() {
    List L;
    infotype x;
    address P, Q;
    bool ditemukan;

    CreateList(L);

    for (int i = 1; i <= 4; i++) {
        cout << "masukkan nomor polisi: ";
        cin >> x.nopol;

        cout << "masukkan warna kendaraan: ";
        cin >> x.warna;

        cout << "masukkan tahun kendaraan: ";
        cin >> x.thnBuat;

        ditemukan = false;
        Q = L.First;

        while (Q != nullptr) {
            if (Q->info.nopol == x.nopol) {
                ditemukan = true;
                break;
            }

            Q = Q->next;
        }

        if (ditemukan) {
            cout << "nomor polisi sudah terdaftar" << endl;
        } else {
            P = alokasi(x);
            insertLast(L, P);
        }

        cout << endl;
    }

    cout << "DATA LIST 1" << endl;
    printInfo(L);

    // Kode tambahan buat jawab no. 2
    cout << endl;
    cout << "Masukkan Nomor Polisi yang dicari    : ";
    cin >> x.nopol;

    P = findElm(L, x);

    if (P != nullptr) {
        cout << endl;
        cout << "Nomor Polisi : " << P->info.nopol << endl;
        cout << "Warna        : " << P->info.warna << endl;
        cout << "Tahun        : " << P->info.thnBuat << endl;
    } else {
        cout << "Data kendaraan tidak ditemukan" << endl;
    }

    return 0;
}
```
### Output Unguided 2 :

##### Output 1
![Output Unguided 2](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/ca67ec001806bd1189700483b9e5128e564ebf10/Semester%203/Week%207%20-%20Modul%206/Unguided%202/Output-1.png)

#### penjelasan unguided 2
Di kode Doublylist.h ditambahkan deklarasi fungsi findElm yang nantinya digunakan untuk mencari kendaraan berdasarkan nomor polisi.

Di kode Doublylist.cpp dibuat fungsi findElm untuk mencari nomor polisi di dalam list. Pencarian dimulai dari node pertama dan terus berpindah menggunakan next sampai data ditemukan.

Di kode main.cpp ditambahkan input nomor polisi yang ingin dicari. Setelah itu, fungsi findElm dipanggil. Kalau nomor polisi ditemukan, warna dan tahun kendaraan ikut ditampilkan. Kalau tidak ada, program akan menampilkan pesan bahwa data tidak ditemukan.

### 3. Doublylist no 3

```C++
Doublylist.h
#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H

#include <iostream>
#include <string>

using namespace std;

typedef struct kendaraan {
    string nopol;
    string warna;
    int thnBuat;
} infotype;

typedef struct ElmList *address;

typedef struct ElmList {
    infotype info;
    address next;
    address prev;
} ElmList;

typedef struct {
    address First;
    address Last;
} List;

void CreateList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertLast(List &L, address P);
address findElm(List L, infotype x);

// Kode tambahan buat jawab no. 3
void deleteFirst(List &L, address &P);
void deleteLast(List &L, address &P);
void deleteAfter(address Prec, address &P);
void deleteElm(List &L, address P);

#endif

---------------------------------------------------------------------------------------------------------------------------------------

Doublylist.cpp
#include "Doublylist.h"

void CreateList(List &L) {
    L.First = nullptr;
    L.Last = nullptr;
}

address alokasi(infotype x) {
    address P = new ElmList;

    P->info = x;
    P->next = nullptr;
    P->prev = nullptr;

    return P;
}

void dealokasi(address &P) {
    delete P;
    P = nullptr;
}

void insertLast(List &L, address P) {
    if (L.First == nullptr) {
        L.First = P;
        L.Last = P;
    } else {
        P->prev = L.Last;
        L.Last->next = P;
        L.Last = P;
    }
}

void printInfo(List L) {
    address P = L.Last;

    while (P != nullptr) {
        cout << "no polisi : " << P->info.nopol << endl;
        cout << "warna     : " << P->info.warna << endl;
        cout << "tahun     : " << P->info.thnBuat << endl;

        P = P->prev;
    }
}

address findElm(List L, infotype x) {
    address P = L.First;

    while (P != nullptr) {
        if (P->info.nopol == x.nopol) {
            return P;
        }

        P = P->next;
    }

    return nullptr;
}

// Kode tambahan buat jawab no. 3

void deleteFirst(List &L, address &P) {
    P = L.First;

    if (L.First == L.Last) {
        L.First = nullptr;
        L.Last = nullptr;
    } else {
        L.First = L.First->next;
        L.First->prev = nullptr;
    }

    dealokasi(P);
}

void deleteLast(List &L, address &P) {
    P = L.Last;

    if (L.First == L.Last) {
        L.First = nullptr;
        L.Last = nullptr;
    } else {
        L.Last = L.Last->prev;
        L.Last->next = nullptr;
    }

    dealokasi(P);
}

void deleteAfter(address Prec, address &P) {
    P = Prec->next;

    Prec->next = P->next;

    if (P->next != nullptr) {
        P->next->prev = Prec;
    }

    dealokasi(P);
}

void deleteElm(List &L, address P) {
    address Prec;

    if (P == L.First) {
        deleteFirst(L, P);
    } else if (P == L.Last) {
        deleteLast(L, P);
    } else {
        Prec = P->prev;
        deleteAfter(Prec, P);
    }
}

---------------------------------------------------------------------------------------------------------------------------------------

main.cpp
#include "Doublylist.h"

int main() {
    List L;
    infotype x;
    address P, Q;
    bool ditemukan;

    CreateList(L);

    for (int i = 1; i <= 4; i++) {
        cout << "masukkan nomor polisi: ";
        cin >> x.nopol;

        cout << "masukkan warna kendaraan: ";
        cin >> x.warna;

        cout << "masukkan tahun kendaraan: ";
        cin >> x.thnBuat;

        ditemukan = false;
        Q = L.First;

        while (Q != nullptr) {
            if (Q->info.nopol == x.nopol) {
                ditemukan = true;
                break;
            }

            Q = Q->next;
        }

        if (ditemukan) {
            cout << "nomor polisi sudah terdaftar" << endl;
        } else {
            P = alokasi(x);
            insertLast(L, P);
        }

        cout << endl;
    }

    cout << "DATA LIST 1" << endl;
    printInfo(L);

    // Kode tambahan buat jawab no. 2
    cout << endl;
    cout << "Masukkan Nomor Polisi yang dicari    : ";
    cin >> x.nopol;

    P = findElm(L, x);

    if (P != nullptr) {
        cout << endl;
        cout << "Nomor Polisi : " << P->info.nopol << endl;
        cout << "Warna        : " << P->info.warna << endl;
        cout << "Tahun        : " << P->info.thnBuat << endl;
    } else {
        cout << "Data kendaraan tidak ditemukan" << endl;
    }

    // Kode tambahan buat jawab no. 3
    cout << endl;
    cout << "Masukkan Nomor Polisi yang akan dihapus    : ";
    cin >> x.nopol;

    P = findElm(L, x);

    if (P != nullptr) {
        deleteElm(L, P);
        cout << "Data dengan nomor polisi " << x.nopol
             << " berhasil dihapus." << endl;
    } else {
        cout << "Data kendaraan tidak ditemukan." << endl;
    }

    cout << endl;
    cout << "DATA LIST 1" << endl;
    printInfo(L);

    return 0;
}
```
### Output Unguided 3 :
##### Output 1
![Output Unguided 3](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/ca67ec001806bd1189700483b9e5128e564ebf10/Semester%203/Week%207%20-%20Modul%206/Unguided%203/Output-1.png)

#### penjelasan unguided 3
Di kode Doublylist.h ditambahkan deklarasi tiga prosedur untuk menghapus node, yaitu deleteFirst, deleteLast, dan deleteAfter. Ada juga deleteElm untuk menentukan prosedur mana yang digunakan saat menghapus data.

Di kode Doublylist.cpp dibuat isi dari prosedur penghapusan. deleteFirst untuk menghapus data pertama, deleteLast untuk data terakhir, sedangkan deleteAfter untuk menghapus data setelah node tertentu. Ketiganya juga mengatur ulang pointer agar hubungan antar-node tetap benar.

Di kode main.cpp ditambahkan input nomor polisi yang ingin dihapus. Setelah itu, program mencari datanya menggunakan findElm. Kalau ditemukan, data akan dihapus menggunakan deleteElm, lalu list ditampilkan kembali untuk melihat hasilnya.

## Kesimpulan
Dari praktikum ini, saya jadi lebih memahami cara kerja Doubly Linked List dalam menyimpan dan mengelola data. Mulai dari membuat list, menambahkan data, mencari data berdasarkan nomor polisi, sampai menghapus data yang sudah tidak diperlukan. Saya juga belajar bagaimana pointer `next` dan `prev` saling terhubung untuk mengatur setiap node. Dengan latihan ini, saya jadi lebih paham bagaimana penerapan Doubly Linked List menggunakan bahasa C++.

## Referensi
[1] Morin, P. (2013). *Open Data Structures: An Introduction*. Athabasca University Press. https://doi.org/10.15215/aupress/9781927356388.01
[2] Goodrich, M. T., Tamassia, R., & Goldwasser, M. H. (2014). *Data Structures and Algorithms in Java* (6th ed.). John Wiley & Sons.
