# <h1 align="center">Laporan Praktikum Modul 2 - PENGENALAN BAHASA C++ (BAGIAN KEDUA)</h1>
<p align="center">Hilman Aniq Almunawa - 109082500027</p>

## Dasar Teori

C++ merupakan bahasa pemrograman yang memiliki beberapa konsep dasar yang dapat digunakan untuk mengolah dan menyimpan data. Pada Modul 2 ini dibahas mengenai array, pointer, fungsi, prosedur, serta cara melewatkan parameter pada fungsi [1].

### A. Array<br/>

Array merupakan kumpulan data yang memiliki nama yang sama dan setiap elemennya memiliki tipe data yang sama. Setiap elemen array dapat diakses menggunakan indeks. Pada array satu dimensi, indeks dimulai dari 0 sampai jumlah elemen dikurangi 1 [1].

#### 1. Array Satu Dimensi

Array satu dimensi merupakan array yang hanya memiliki satu larik data. Array dapat dideklarasikan dengan menentukan tipe data, nama variabel, dan ukuran array. Contohnya `int nilai[10];` yang berarti array nilai memiliki 10 elemen bertipe integer [1].

#### 2. Array Dua Dimensi

Array dua dimensi merupakan array yang memiliki dua indeks dan dapat digunakan untuk menyimpan data dalam bentuk tabel. Array ini memiliki dimensi pertama dan dimensi kedua yang digunakan untuk menentukan posisi setiap elemen [1].

### B. Pointer<br/>

Pointer merupakan variabel yang digunakan untuk menyimpan alamat memori dari variabel lain. Setiap data yang digunakan oleh program disimpan di dalam memori dan setiap lokasi memori memiliki alamat yang berbeda. Untuk mengetahui alamat dari suatu variabel dapat digunakan tanda `&` di depan nama variabel [1].

#### 1. Pointer dan Alamat

Pointer dapat digunakan untuk menunjuk alamat dari variabel tertentu. Pointer dideklarasikan menggunakan tanda `*`, misalnya `int *p_int;`. Setelah pointer diberikan alamat suatu variabel menggunakan `p_int = &j`, pointer tersebut dapat digunakan untuk mengakses nilai dari variabel yang ditunjuk [1].

#### 2. Pointer dan Array

Pointer memiliki hubungan dengan array karena pointer dapat digunakan untuk menunjuk elemen pada array. Jika pointer menunjuk ke elemen pertama array, maka pointer dapat digunakan untuk mengakses elemen berikutnya dengan melakukan operasi penambahan pada pointer [1].

### C. Fungsi<br/>

Fungsi merupakan blok kode yang dibuat untuk melakukan tugas tertentu. Penggunaan fungsi dapat membuat program menjadi lebih terstruktur dan mengurangi pengulangan kode. Fungsi dapat menerima parameter dan dapat menghasilkan nilai balik sesuai dengan tipe data yang digunakan [1].

### D. Prosedur<br/>

Prosedur merupakan fungsi yang tidak mengembalikan nilai. Dalam C++, prosedur biasanya menggunakan `void`. Prosedur digunakan untuk menjalankan tugas tertentu tanpa memberikan nilai balik kepada bagian program yang memanggilnya [1].

### E. Parameter Fungsi<br/>

Parameter fungsi digunakan untuk memberikan data yang akan diproses oleh fungsi. Parameter formal merupakan variabel yang terdapat pada saat fungsi didefinisikan, sedangkan parameter aktual merupakan nilai atau variabel yang digunakan ketika fungsi dipanggil [1].

### F. Pemanggilan dengan Nilai (Call by Value)<br/>

Call by value merupakan cara melewatkan parameter dengan menyalin nilai dari parameter aktual ke parameter formal. Perubahan yang dilakukan pada parameter formal tidak akan mengubah nilai dari parameter aktual karena yang digunakan dalam fungsi adalah salinan dari nilai tersebut [1].

### G. Pemanggilan dengan Pointer (Call by Pointer)<br/>

Call by pointer merupakan cara melewatkan alamat suatu variabel ke dalam fungsi. Dengan cara ini, fungsi dapat mengubah nilai dari variabel yang berada di luar fungsi. Pada pemanggilannya, alamat variabel diberikan menggunakan tanda `&`, sedangkan parameter fungsi menggunakan pointer [1].

### H. Pemanggilan dengan Referensi (Call by Reference)<br/>

Call by reference merupakan cara melewatkan variabel dengan menggunakan referensi. Dengan cara ini, perubahan yang dilakukan pada parameter fungsi dapat mengubah nilai variabel yang digunakan saat pemanggilan fungsi. Pada parameter fungsi digunakan tanda `&`, sedangkan saat pemanggilan tidak perlu menggunakan tanda tambahan [1].


## Guided 

### 1. array 1

```C++
#include<iostream>
using namespace std;

int main(){
    int nilai[5];

    nilai[0] = 80;
    nilai[1] = 75;
    nilai[2] = 90;
    nilai[3] = 85;
    nilai[4] = 95;

    for(int i = 0; i < 5; i++){
        cout << "Nilai ke-" << i + 1 << " = "
        << nilai[i] << endl;
    }

    return 0;
}
```
#### penjelasan guided 1
Pada program ini saya membuat array satu dimensi untuk menyimpan 5 nilai. Setiap nilai dimasukkan ke dalam array menggunakan indeks dari 0 sampai 4. Setelah itu, program menggunakan perulangan `for` untuk menampilkan semua nilai yang ada di dalam array.

### 2. array 2

```C++
#include <iostream>
using namespace std;

int main() {
    int nilai[3][3] = {
        {80, 75, 90},
        {85, 90, 88},
        {70, 80, 85}
    };
    //print nilai array 2 dimensi
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            cout << nilai[i][j] << " ";
        }

        cout << endl;
    }
    cout << endl;
    cout << nilai[1][2] << endl; //88
    return 0;
}
```
#### penjelasan guided 2
Pada program ini saya membuat array dua dimensi dengan ukuran 3x3 untuk menyimpan nilai. Program menggunakan perulangan `for` bersarang untuk menampilkan semua nilai yang ada di dalam array. Selain itu, program juga menampilkan nilai pada indeks `[1][2]` yang menghasilkan nilai 88.

### 3. array 3

```C++
#include <iostream>
using namespace std;

int main(){
    int data[2][2][3] = {
        {
            {10, 20, 30},
            {40, 50, 60}
        },
        {
            {70, 80, 90},
            {100, 110, 120}
        }
    };

    cout << data[0][1][2] << endl; //60

    return 0;

}
```
#### penjelasan guided 3
Pada program ini saya membuat array berdimensi tiga dengan ukuran `2x2x3`. Data disimpan pada beberapa indeks yang berbeda. Program kemudian mengambil nilai pada indeks `[0][1][2]` dan hasil yang ditampilkan adalah 60.

### 4. address

```C++
#include <iostream>
using namespace std;

int main(){
    int angka  = 100;

    cout << "Nilai angka: " << angka << endl;
    cout << "Alamat angka: " << &angka << endl;

    return 0;
}
```
#### penjelasan guided 4
Pada program ini saya membuat sebuah variabel `angka` dengan nilai 100. Program menampilkan nilai dari variabel tersebut menggunakan `cout`. Setelah itu, program menampilkan alamat memori dari variabel `angka` menggunakan tanda `&`.

### 5. pointer 1

```C++
#include <iostream>
using namespace std;

int main() {
    char arr[6];

    arr[0] = 'a';
    arr[1] = 'b';
    arr[2] = 'c';
    arr[3] = 'b';
    arr[4] = 'd';
    arr[5] = 'e';

    cout << arr[3] << endl; //value b
    cout << &(arr[4]) << endl; //alamat memory atau address dari arr[4] yaitu d

    return 0;
}
```
#### penjelasan guided 5
Pada program ini saya membuat array karakter dengan 6 elemen. Setiap elemen array diisi dengan karakter yang berbeda. Program kemudian menampilkan nilai dari `arr[3]` yang menghasilkan karakter `b`, lalu menampilkan alamat memori dari `arr[4]` menggunakan tanda `&`.

### 6. pointer 2

```C++
#include <iostream>
using namespace std;

int main(){
    int angka = 100;

    int *pointer;

    pointer = &angka;

    cout << "Nilai angka        : " << angka << endl; //100
    cout << "Alamat angka       : " << &angka << endl; //address
    cout << "Isi pointer        : " << pointer << endl; //addres angka
    cout << "Nilai dari pointer : " << *pointer << endl; //value dari angka

    return 0;
}
```
#### penjelasan guided 6
Pada program ini saya membuat variabel `angka` dengan nilai 100 dan sebuah pointer untuk menyimpan alamat dari variabel tersebut. Pointer diarahkan ke alamat `angka` menggunakan `&angka`. Setelah itu program menampilkan nilai, alamat angka, isi pointer, dan nilai yang ditunjuk oleh pointer menggunakan `*pointer`.

### 7. function

```C++
#include <iostream>
using namespace std;

int maks3(int a, int b, int c) {
    int temp_max = a;

    if (b > temp_max) {
        temp_max = b;
    }

    if (c > temp_max) {
        temp_max = c;
    }

    return temp_max;
}

int main() {
    int x, y, z;

    cout << "Masukkan nilai 1: ";
    cin >> x;

    cout << "Masukkan nilai 2: ";
    cin >> y;

    cout << "Masukkan nilai 3: ";
    cin >> z;
    
    cout << "Nilai maksimum = "
         << maks3(x, y, z) << endl;

    return 0;
}
```
#### penjelasan guided 7
Pada program ini saya membuat fungsi `maks3` untuk mencari nilai terbesar dari tiga bilangan. Program menerima tiga nilai dari pengguna, kemudian ketiga nilai tersebut dikirim ke fungsi `maks3`. Di dalam fungsi, nilai dibandingkan satu per satu dan nilai terbesar disimpan pada `temp_max`, kemudian hasilnya dikembalikan dan ditampilkan.

### 8. procedure

```C++
#include <iostream>
using namespace std;

void sapa(){
    cout << "Selamat datang di praktikum struktur data" << endl;
}

int main(){
    sapa();
    return 0;
}
```
#### penjelasan guided 8
Pada program ini saya membuat prosedur `sapa()` untuk menampilkan pesan selamat datang. Prosedur tersebut menggunakan `void` karena tidak mengembalikan nilai. Pada fungsi `main()`, prosedur `sapa()` dipanggil sehingga pesan dapat ditampilkan.

### 9. callByValue_Pointer_Reference

```C++
#include <iostream>
using namespace std;

// Call by Value
void tukarValue(int x, int y) {
    int temp;

    temp = x;
    x = y;
    y = temp;
}

// Call by Pointer
void tukarPointer(int *x, int *y) {
    int temp;

    temp = *x;
    *x = *y;
    *y = temp;
}

// Call by Reference
void tukarReference(int &x, int &y) {
    int temp;

    temp = x;
    x = y;
    y = temp;
}

int main() {
    int a, b;

    // CALL BY VALUE
    a = 4;
    b = 6;

    cout << "=== CALL BY VALUE ===" << endl;
    cout << "Sebelum ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    tukarValue(a, b);

    cout << "Setelah ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    // CALL BY POINTER
    a = 4;
    b = 6;

    cout << "\n=== CALL BY POINTER ===" << endl;
    cout << "Sebelum ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    tukarPointer(&a, &b);

    cout << "Setelah ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    // CALL BY REFERENCE
    a = 4;
    b = 6;

    cout << "\n=== CALL BY REFERENCE ===" << endl;
    cout << "Sebelum ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    tukarReference(a, b);

    cout << "Setelah ditukar:" << endl;
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;

    return 0;
}
```
#### penjelasan guided 9
Pada program ini saya menggabungkan call by value, call by pointer, dan call by reference untuk melakukan pertukaran nilai `a` dan `b`. Ketiga metode tersebut digunakan untuk melihat perbedaan cara parameter dikirim ke dalam fungsi.


## Unguided 

### 1. Buatlah program yang dapat melakukan operasi penjumlahan, pengurangan, dan perkalian matriks 3x3

```C++
source code unguided 1
#include <iostream>
using namespace std;

int main() {
    int A[3][3] = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };

    int B[3][3] = {
        {9, 8, 7},
        {6, 5, 4},
        {3, 2, 1}
    };

    int tambah[3][3];
    int kurang[3][3];
    int kali[3][3] = {0};

    // yg ini penjumlahan sama pengurangan
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            tambah[i][j] = A[i][j] + B[i][j];
            kurang[i][j] = A[i][j] - B[i][j];
        }
    }

    // kalo yg ini perkalian
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            for (int k = 0; k < 3; k++) {
                kali[i][j] += A[i][k] * B[k][j];
            }
        }
    }

    cout << "Hasil Penjumlahan Matriks:" << endl;
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            cout << tambah[i][j] << " ";
        }
        cout << endl;
    }

    cout << "\nHasil Pengurangan Matriks:" << endl;
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            cout << kurang[i][j] << " ";
        }
        cout << endl;
    }

    cout << "\nHasil Perkalian Matriks:" << endl;
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            cout << kali[i][j] << " ";
        }
        cout << endl;
    }

    return 0;
}
```
#### Output Unguided 1 :

##### Output 1
![Screenshot Output Unguided 1_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%203%20-%20Modul%202/Unguided%201/Output-1.png)

#### penjelasan unguided 1 
Pada program ini saya membuat dua matriks berukuran 3x3 menggunakan array dua dimensi. Nilai dari matriks A dan B langsung dimasukkan ke dalam program. Selanjutnya, program melakukan operasi penjumlahan, pengurangan, dan perkalian matriks menggunakan perulangan `for`. Hasil dari setiap operasi kemudian ditampilkan ke layar.

### 2. Berdasarkan guided pointer dan reference sebelumnya, buatlah keduanya dapat menukar nilai dari 3 variabel.

```C++
source code unguided 2
#include <iostream>
using namespace std;

void tukarPointer(int *a, int *b, int *c) {
    int temp;

    temp = *a;
    *a = *b;
    *b = *c;
    *c = temp;
}

void tukarReference(int &a, int &b, int &c) {
    int temp;

    temp = a;
    a = b;
    b = c;
    c = temp;
}

int main() {
    int a = 10, b = 20, c = 30;

    cout << "Nilai awal: " << a << " " << b << " " << c << endl;

    tukarPointer(&a, &b, &c);
    cout << "Setelah pointer: " << a << " " << b << " " << c << endl;

    tukarReference(a, b, c);
    cout << "Setelah reference: " << a << " " << b << " " << c << endl;

    return 0;
}
```
### Output Unguided 2 :

##### Output 1
![Screenshot Output Unguided 2_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%203%20-%20Modul%202/Unguided%202/Output-1.png)

#### penjelasan unguided 2
Pada program ini saya menggunakan pointer dan reference untuk menukar nilai dari tiga variabel. Pada bagian pointer, alamat dari variabel dikirim ke dalam fungsi, sedangkan pada bagian reference variabel dikirim menggunakan tanda `&`. Nilai dari ketiga variabel kemudian ditukar secara bergantian dengan bantuan variabel `temp`.

### 3.  Diketahui sebuah array 1 dimensi sebagai berikut : arrA = {11, 8, 5, 7, 12, 26, 3, 54, 33, 55}. Buatlah program yang dapat mencari nilai minimum, maksimum, dan rata – rata dari array tersebut! Gunakan function `cariMinimum()` untuk mencari nilai minimum dan function `cariMaksimum()` untuk mencari nilai maksimum, serta gunakan prosedur `hitungRataRata()` untuk menghitung nilai rata – rata! Buat program menggunakan menu switch-case seperti berikut ini :

**--- Menu Program Array ---**

- Tampilkan isi array
- cari nilai maksimum
- cari nilai minimum
- Hitung nilai rata - rata

```C++
source code unguided 3
#include <iostream>
using namespace std;

int cariMinimum(int arrA[]) {
    int minimum = arrA[0];

    for (int i = 1; i < 10; i++) {
        if (arrA[i] < minimum) {
            minimum = arrA[i];
        }
    }

    return minimum;
}

int cariMaksimum(int arrA[]) {
    int maksimum = arrA[0];

    for (int i = 1; i < 10; i++) {
        if (arrA[i] > maksimum) {
            maksimum = arrA[i];
        }
    }

    return maksimum;
}

void hitungRataRata(int arrA[]) {
    int total = 0;

    for (int i = 0; i < 10; i++) {
        total = total + arrA[i];
    }

    cout << "Nilai rata-rata = " << (float)total / 10 << endl;
}

int main() {
    int arrA[10] = {11, 8, 5, 7, 12, 26, 3, 54, 33, 55};
    int pilihan;

    cout << "--- Menu Program Array ---" << endl;
    cout << "1. Tampilkan isi array" << endl;
    cout << "2. Cari nilai maksimum" << endl;
    cout << "3. Cari nilai minimum" << endl;
    cout << "4. Hitung nilai rata-rata" << endl;
    cout << "Pilihan: ";
    cin >> pilihan;

    switch (pilihan) {
        case 1:
            cout << "Isi array: ";
            for (int i = 0; i < 10; i++) {
                cout << arrA[i] << " ";
            }
            cout << endl;
            break;

        case 2:
            cout << "Nilai maksimum = " << cariMaksimum(arrA) << endl;
            break;

        case 3:
            cout << "Nilai minimum = " << cariMinimum(arrA) << endl;
            break;

        case 4:
            hitungRataRata(arrA);
            break;

        default:
            cout << "Pilihan tidak tersedia." << endl;
    }

    return 0;
}
```
### Output Unguided 3 :

##### Output 1
![Screenshot Output Unguided 3_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%203%20-%20Modul%202/Unguided%203/Output-1.png)

##### Output 2
![Screenshot Output Unguided 3_2](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%203%20-%20Modul%202/Unguided%203/Output-2.png)

#### penjelasan unguided 3
Pada program ini saya membuat array satu dimensi yang berisi 10 nilai. Program menggunakan function `cariMinimum()` untuk mencari nilai terkecil dan function `cariMaksimum()` untuk mencari nilai terbesar. Selain itu, saya menggunakan prosedur `hitungRataRata()` untuk menghitung nilai rata-rata dari seluruh isi array. Program juga menggunakan menu `switch-case` untuk memilih operasi yang ingin dijalankan.

## Kesimpulan
Pada praktikum Modul 2 ini saya mempelajari penggunaan array, pointer, fungsi, prosedur, dan parameter pada bahasa C++. Dari praktikum yang dilakukan, saya dapat memahami cara menggunakan array satu dimensi, array dua dimensi, dan array berdimensi banyak. Saya juga mempelajari cara melihat alamat memori dan menggunakan pointer untuk mengakses nilai dari suatu variabel.

Selain itu, saya mempelajari penggunaan fungsi dan prosedur serta cara melewatkan parameter menggunakan call by value, call by pointer, dan call by reference. Dari beberapa program yang dibuat, saya jadi lebih memahami perbedaan cara kerja masing-masing dan bagaimana penggunaannya dalam program.

## Referensi
[1] Triase. (2020). Diktat Edisi Revisi: STRUKTUR DATA. Medan: Universitas Islam Negeri Sumatera Utara Medan.