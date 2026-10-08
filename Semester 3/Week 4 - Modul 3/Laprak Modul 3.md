# <h1 align="center">Laporan Praktikum Modul 3 - ABSTRACT DATA TYPE (ADT)   </h1>
<p align="center">Hilman Aniq Almunawa - 109082500027</p>

## Dasar Teori
Abstract Data Type (ADT) merupakan suatu tipe data yang memiliki data serta operasi yang dapat dilakukan terhadap data tersebut. ADT lebih menekankan pada apa yang dapat dilakukan terhadap suatu data, sedangkan cara data tersebut diimplementasikan tidak menjadi bagian dari definisinya.

Dalam penggunaannya, ADT memiliki sekumpulan operasi yang digunakan untuk mengolah data. Dengan adanya operasi tersebut, program dapat dibuat lebih terstruktur karena bagian data dan operasi yang digunakan dapat dikelompokkan sesuai kebutuhannya. ADT juga memisahkan bagian spesifikasi dengan implementasinya sehingga cara kerja di dalamnya tidak harus diketahui saat menggunakan ADT.

Pada bahasa C++, konsep ADT dapat diterapkan dengan membuat tipe data menggunakan struct dan menyediakan fungsi atau prosedur untuk mengolah data tersebut. Bagian deklarasi dapat dipisahkan dari implementasinya menggunakan file header dan file .cpp. Dengan cara ini, program menjadi lebih modular dan fungsi yang dibuat dapat digunakan dari program utama.

Dalam praktikum ini, konsep ADT diterapkan dengan membuat beberapa tipe data dan fungsi yang berhubungan dengan data tersebut. Penerapannya meliputi pengolahan data mahasiswa, pembuatan ADT pelajaran, serta penggunaan array 2D dan pointer. Konsep tersebut membantu memahami bagaimana data dapat dikelompokkan dan diolah menggunakan fungsi yang sudah dibuat.

## Guided 

### 1. mahasiswa.h

```C++
#ifndef MAHASISWA_H_INCLUDED //untuk ngecek apakah header sudah di pakai di folder lain atau belum
#define MAHASISWA_H_INCLUDED //membuat header mahasiswa.h

struct mahasiswa{
    char nim[10];
    int nilai1, nilai2;
};

void inputMhs (mahasiswa &m);
float rata2 (mahasiswa m);
#endif // MAHASISWA_H_INCLUDED```C++
```
#### penjelasan guided 1
Pada kode ini dibuat struct mahasiswa yang berisi nim, nilai1, dan nilai2. Kemudian inputMhs() dan rata2() dideklarasikan untuk digunakan pada program. Bagian #ifndef, #define, dan #endif digunakan supaya header tidak didefinisikan lebih dari satu kali.

### 2. mahasiswa.cpp

```C++
#include <iostream>
#include "mahasiswa.h"
using namespace std;

void inputMhs (mahasiswa &m) {
    cout << "input nim = ";
    cin >> (m).nim;
    cout << "input nilai1 = ";
    cin >> (m).nilai1;
    cout << "input nilai2 = ";
    cin >> (m).nilai2;
}

float rata2 (mahasiswa m) {
    return float(m.nilai1 + m.nilai2)/2;
}
```
#### penjelasan guided 2
Pada kode ini merupakan isi dari fungsi yang sebelumnya sudah dibuat di mahasiswa.h. Fungsi inputMhs() digunakan untuk memasukkan NIM, nilai1, dan nilai2. Sedangkan rata2() digunakan untuk menghitung rata-rata dari nilai1 dan nilai2.

### 3. main.cpp

```C++
#include <iostream>
#include "mahasiswa.h"
using namespace std;

int main()
{
    mahasiswa mhs;
    inputMhs (mhs);
    cout << "Rata-rata = " << rata2 (mhs);
    return 0;
}
```
#### penjelasan guided 3
Pada kode ini digunakan main() sebagai program utama. Pertama dibuat variabel mhs dengan tipe mahasiswa, kemudian fungsi inputMhs() dipanggil untuk memasukkan data NIM, nilai1, dan nilai2. Setelah itu fungsi rata2() digunakan untuk menghitung nilai rata-rata dan hasilnya ditampilkan dengan cout.

## Unguided 

### 1. dataMahasiswa
dataMahasiswa.h
```C++
#ifndef DATA_MAHASISWA_H
#define DATA_MAHASISWA_H

struct mahasiswa {
    char nama[30];
    char nim[20];
    float uts, uas, tugas, nilaiAkhir;
};

float hitungNilaiAkhir(mahasiswa m);

#endif // DATA_MAHASISWA_H

--------------------------------------------
dataMahasiswa.cpp
#include <iostream>
#include "dataMahasiswa.h"
using namespace std;

float hitungNilaiAkhir(mahasiswa m)
{
    float nilai;

    nilai = (0.3 * m.uts) + (0.4 * m.uas) + (0.3 * m.tugas);

    return nilai;
}

---------------------------------------------
main.cpp
#include <iostream>
#include "dataMahasiswa.h"
using namespace std;

int main()
{
    mahasiswa mhs[10];

    cout << "Nama = ";
    cin >> mhs[0].nama;
    cout << "NIM = ";
    cin >> mhs[0].nim;
    cout << "UTS = ";
    cin >> mhs[0].uts;
    cout << "UAS = ";
    cin >> mhs[0].uas;
    cout << "Tugas = ";
    cin >> mhs[0].tugas;

    mhs[0].nilaiAkhir = hitungNilaiAkhir(mhs[0]);

    cout << "Nilai Akhir = " << mhs[0].nilaiAkhir;

    return 0;
}
```
#### Output Unguided 1 :

##### Output 1
![Screenshot Output Unguided 1_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%203%20-%20Modul%202/Unguided%201/Output-1.png)

##### Output 2


#### penjelasan unguided 1 
Di bagian dataMahasiswa.h saya membuat struct mahasiswa yang berisi nama, NIM, UTS, UAS, tugas, dan nilai akhir. Lalu saya buat fungsi hitungNilaiAkhir() untuk menghitung nilai akhir.

Di file dataMahasiswa.cpp, fungsi tersebut berisi perhitungan nilai akhir sesuai rumus yang ada di soal. Setelah itu di main.cpp saya membuat array mhs[10] untuk menampung data mahasiswa, kemudian memasukkan nilai UTS, UAS, dan tugas. Setelah semua data dimasukkan, fungsi hitungNilaiAkhir() dipanggil dan hasil nilai akhirnya ditampilkan.


### 2. pelajar

```C++
pelajaran.h
#ifndef PELAJARAN_H
#define PELAJARAN_H

#include <string>
using namespace std;

struct pelajaran {
    string namaMapel;
    string kodeMapel;
};

pelajaran create_pelajaran(string namapel, string kodepel);
void tampil_pelajaran(pelajaran pel);

#endif // PELAJARAN_H

-------------------------------------------------------------
pelajaran.cpp
#include <iostream>
#include "pelajaran.h"
using namespace std;

pelajaran create_pelajaran(string namapel, string kodepel)
{
    pelajaran pel;

    pel.namaMapel = namapel;
    pel.kodeMapel = kodepel;

    return pel;
}

void tampil_pelajaran(pelajaran pel)
{
    cout << "nama pelajaran : " << pel.namaMapel << endl;
    cout << "nilai : " << pel.kodeMapel << endl;
}

------------------------------------------------------------
main.cpp
#include <iostream>
#include "pelajaran.h"
using namespace std;

int main()
{
    string namapel = "Struktur Data";
    string kodepel = "STD";

    pelajaran pel = create_pelajaran(namapel, kodepel);
    tampil_pelajaran(pel);

    return 0;
}
```
### Output Unguided 2 :

##### Output 1
![Screenshot Output Unguided 2_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%203%20-%20Modul%202/Unguided%202/Output-1.png)

#### penjelasan unguided 2
Di bagian pelajaran.h saya membuat struct pelajaran yang berisi nama mata pelajaran dan kode mata pelajaran. Lalu dibuat fungsi create_pelajaran() untuk membuat data pelajaran dan tampil_pelajaran() untuk menampilkan datanya.

Di file pelajaran.cpp, fungsi create_pelajaran() digunakan untuk memasukkan nama dan kode ke dalam struct pelajaran. Setelah itu tampil_pelajaran() digunakan untuk menampilkan data yang sudah dibuat.

Kemudian di main.cpp saya mengisi nama mata pelajaran dengan Struktur Data dan kodenya STD. Setelah itu data tersebut dibuat menggunakan create_pelajaran() dan ditampilkan menggunakan tampil_pelajaran().

### 3.  array2D

```C++
array2D.h
#ifndef ARRAY2D_H
#define ARRAY2D_H

void tampilArray(int arr[3][3]);
void tukarArray(int arr1[3][3], int arr2[3][3], int baris, int kolom);
void tukarPointer(int *p1, int *p2);

#endif // ARRAY2D_H

-------------------------------------------------------------------------
array2D.cpp
#include <iostream>
#include "array2D.h"
using namespace std;

void tampilArray(int arr[3][3])
{
    for (int i = 0; i < 3; i++)
    {
        for (int j = 0; j < 3; j++)
        {
            cout << arr[i][j] << " ";
        }
        cout << endl;
    }
}

void tukarArray(int arr1[3][3], int arr2[3][3], int baris, int kolom)
{
    int temp;

    temp = arr1[baris][kolom];
    arr1[baris][kolom] = arr2[baris][kolom];
    arr2[baris][kolom] = temp;
}

void tukarPointer(int *p1, int *p2)
{
    int temp;

    temp = *p1;
    *p1 = *p2;
    *p2 = temp;
}
-----------------------------------------------------------------------------
main.cpp
#include <iostream>
#include "array2D.h"
using namespace std;

int main()
{
    int array1[3][3] = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };

    int array2[3][3] = {
        {10, 11, 12},
        {13, 14, 15},
        {16, 17, 18}
    };

    int a = 20;
    int b = 30;

    int *p1 = &a;
    int *p2 = &b;

    cout << "Array 1 sebelum ditukar :" << endl;
    tampilArray(array1);

    cout << "\nArray 2 sebelum ditukar :" << endl;
    tampilArray(array2);

    tukarArray(array1, array2, 1, 1);

    cout << "\nArray 1 setelah ditukar :" << endl;
    tampilArray(array1);

    cout << "\nArray 2 setelah ditukar :" << endl;
    tampilArray(array2);

    cout << "\nNilai sebelum pointer ditukar : " << *p1 << " dan " << *p2 << endl;

    tukarPointer(p1, p2);

    cout << "Nilai setelah pointer ditukar : " << *p1 << " dan " << *p2 << endl;

    return 0;
}

```
### Output Unguided 3 :

##### Output 1
![Screenshot Output Unguided 3_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%203%20-%20Modul%202/Unguided%203/Output-1.png)

#### penjelasan unguided 3
Di bagian array2D.h saya membuat beberapa fungsi untuk digunakan pada program, yaitu fungsi untuk menampilkan array, menukar isi dua array pada posisi tertentu, dan menukar nilai dari dua pointer.

Di file array2D.cpp, fungsi-fungsi tersebut dibuat. tampilArray() digunakan untuk menampilkan isi array 3x3, tukarArray() digunakan untuk menukar nilai pada posisi yang ditentukan, sedangkan tukarPointer() digunakan untuk menukar nilai yang ditunjuk oleh dua pointer.

Kemudian di main.cpp saya membuat dua array 3x3 dan dua pointer. Setelah itu kedua array ditampilkan, nilai pada posisi [1][1] ditukar, lalu nilai dari kedua pointer juga ditukar. Hasil dari setiap proses kemudian ditampilkan.

## Kesimpulan
Setelah melakukan praktikum pada Modul 3, saya jadi lebih memahami konsep Abstract Data Type (ADT) dan cara penerapannya dalam C++. ADT membuat program menjadi lebih terstruktur karena bagian deklarasi, fungsi, dan program utama dapat dipisahkan ke dalam file yang berbeda.

Dari praktikum ini saya juga mempraktikkan penggunaan struct, fungsi, array 2D, dan pointer melalui beberapa program yang dibuat. Dengan latihan tersebut, saya jadi lebih memahami bagaimana data dapat disimpan, diolah menggunakan fungsi, serta bagaimana array dan pointer dapat digunakan dalam program.

## Referensi
[1] Guttag, J. V. (1977). Abstract Data Types and the Development of Data Structures. Communications of the ACM, 20(6), 396–404.