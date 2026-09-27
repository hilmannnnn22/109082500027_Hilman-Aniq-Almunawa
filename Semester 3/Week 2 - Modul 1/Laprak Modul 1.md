    # <h1 align="center">Laporan Praktikum Modul 1 - Codeblocks IDE & Pengenalan Bahasa C++ (Bagian Pertama)</h1>
    <p align="center">Hilman Aniq Almunawa - 109082500027</p>

    ## Dasar Teori

    C++ merupakan bahasa pemrograman yang dikembangkan dari bahasa C. Dalam pemrograman C++ terdapat beberapa hal dasar yang perlu dipahami seperti tipe data, variabel, input, output, operator, dan kondisi. Hal-hal tersebut digunakan untuk membuat program agar dapat menerima, mengolah, dan menampilkan data [1].

    ### A. Pengenalan Bahasa C++<br/>

    C++ memiliki struktur dasar program yang terdiri dari header, deklarasi, fungsi, dan fungsi utama `main()`. Setiap variabel yang digunakan harus dideklarasikan terlebih dahulu sesuai dengan tipe datanya.

    #### 1. Tipe Data

    Tipe data digunakan untuk menentukan jenis data yang akan digunakan dalam program. Beberapa tipe data yang terdapat pada C++ yaitu `int`, `float`, `double`, `char`, dan `long` [1].

    #### 2. Variabel dan Konstanta

    Variabel digunakan untuk menyimpan data yang nilainya dapat berubah selama program dijalankan. Sedangkan konstanta merupakan nilai yang tetap dan tidak berubah selama program berjalan [1].

    #### 3. Input dan Output

    Input digunakan untuk memasukkan data ke dalam program, sedangkan output digunakan untuk menampilkan hasil. Pada C++, input dapat dilakukan menggunakan `cin` dan output menggunakan `cout` [1].

    ### B. Operator dan Kondisional<br/>

    Operator digunakan untuk melakukan operasi terhadap data yang ada di dalam program. Sedangkan kondisional digunakan untuk menjalankan perintah berdasarkan kondisi tertentu.

    #### 1. Operator Aritmatika

    Operator aritmatika digunakan untuk melakukan perhitungan seperti penjumlahan (`+`), pengurangan (`-`), perkalian (`*`), pembagian (`/`), dan sisa pembagian (`%`) [1].

    #### 2. Operator Perbandingan

    Operator perbandingan digunakan untuk membandingkan dua buah nilai. Operator yang digunakan antara lain `==`, `!=`, `<`, `>`, `<=`, dan `>=`.

    #### 3. Kondisional

    Kondisional digunakan untuk menentukan perintah yang dijalankan berdasarkan suatu kondisi. Pada C++ kondisional dapat menggunakan `if`, `if-else`, dan `switch` [1].

    ## Unguided 

    ### 1. Buatlah program yang menerima input-an dua buah bilangan bertipe float, kemudian memberikan output-an hasil penjumlahan, pengurangan, perkalian, dan pembagian dari dua bilangan tersebut.


    ```C++
    source code unguided 1
    #include <iostream>
    using namespace std;

    int main(){
        float bil1, bil2;

        cout << "Masukkan bilangan pertama = ";
        cin >> bil1;

        cout << "Masukkan bilangan kedua = ";
        cin >> bil2;

        cout << "Hasil penjumlahan = " << bil1 + bil2 << endl;
        cout << "Hasil pengurangan = " << bil1 - bil2 << endl;
        cout << "Hasil perkalian = " << bil1 * bil2 << endl;
        cout << "Hasil pembagian = " << bil1 / bil2 << endl;

        return 0;
    }
    ```
    #### Output Unguided 1 :

    ##### Output 1

    ![Screenshot Output Unguided 1_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%202%20-%20Modul%201/Unguided%201/Output-1.png)

    ##### Output 2

    ![Screenshot Output Unguided 1_2](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%202%20-%20Modul%201/Unguided%201/Output-2.png)

    #### penjelasan unguided 1 
    Pada program ini saya membuat program sederhana untuk melakukan operasi pada dua bilangan bertipe float. Pertama, saya membuat variabel bil1 dan bil2 untuk menyimpan dua bilangan yang dimasukkan oleh pengguna. Setelah itu, saya menggunakan cin untuk memasukkan nilai bilangan pertama dan kedua.
    Setelah kedua bilangan dimasukkan, program melakukan empat operasi yaitu penjumlahan, pengurangan, perkalian, dan pembagian. Hasil dari setiap operasi kemudian ditampilkan menggunakan cout.

    ### 2. Buatlah sebuah program yang menerima masukan angka dan mengeluarkan output nilai angka tersebut dalam bentuk tulisan. Angka yang akan di-input-kan user adalah bilangan bulat positif mulai dari 0 s.d 100. Contoh: 79 : tujuh puluh sembilan

    ```C++
    source code unguided 2
    #include <iostream>
    using namespace std;

    int main(){
        int angka;

        cout << "Masukkan angka = ";
        cin >> angka;

        if(angka == 0)
            cout << "nol";
        else if(angka == 100)
            cout << "seratus";
        else if(angka < 20){
            switch(angka){
                case 1: cout << "satu"; break;
                case 2: cout << "dua"; break;
                case 3: cout << "tiga"; break;
                case 4: cout << "empat"; break;
                case 5: cout << "lima"; break;
                case 6: cout << "enam"; break;
                case 7: cout << "tujuh"; break;
                case 8: cout << "delapan"; break;
                case 9: cout << "sembilan"; break;
                case 10: cout << "sepuluh"; break;
                case 11: cout << "sebelas"; break;
                case 12: cout << "dua belas"; break;
                case 13: cout << "tiga belas"; break;
                case 14: cout << "empat belas"; break;
                case 15: cout << "lima belas"; break;
                case 16: cout << "enam belas"; break;
                case 17: cout << "tujuh belas"; break;
                case 18: cout << "delapan belas"; break;
                case 19: cout << "sembilan belas"; break;
            }
        }
        else{
            int puluhan = angka / 10;
            int satuan = angka % 10;

            switch(puluhan){
                case 2: cout << "dua puluh"; break;
                case 3: cout << "tiga puluh"; break;
                case 4: cout << "empat puluh"; break;
                case 5: cout << "lima puluh"; break;
                case 6: cout << "enam puluh"; break;
                case 7: cout << "tujuh puluh"; break;
                case 8: cout << "delapan puluh"; break;
                case 9: cout << "sembilan puluh"; break;
            }

            if(satuan == 1) cout << " satu";
            else if(satuan == 2) cout << " dua";
            else if(satuan == 3) cout << " tiga";
            else if(satuan == 4) cout << " empat";
            else if(satuan == 5) cout << " lima";
            else if(satuan == 6) cout << " enam";
            else if(satuan == 7) cout << " tujuh";
            else if(satuan == 8) cout << " delapan";
            else if(satuan == 9) cout << " sembilan";
        }

        return 0;
    }
    ```
    ### Output Unguided 2 :

    ##### Output 1

    ![Screenshot Output Unguided 2_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%202%20-%20Modul%201/Unguided%202/Output-1.png)

    ##### Output 2

    ![Screenshot Output Unguided 2_2](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%202%20-%20Modul%201/Unguided%202/Output-2.png)

    #### penjelasan unguided 2
    Pada program ini saya membuat program untuk mengubah angka menjadi bentuk tulisan dari angka 0 sampai 100. Pertama, program menerima angka yang dimasukkan menggunakan cin dan disimpan dalam variabel angka. Setelah itu, program mengecek angka yang dimasukkan menggunakan if, else if, dan switch.

    Untuk angka 0, 100, dan angka 1 sampai 19, program langsung menampilkan tulisan yang sesuai. Sedangkan untuk angka 20 sampai 99, angka dipisahkan menjadi puluhan dan satuan menggunakan pembagian dan sisa bagi. Setelah itu hasilnya digabungkan menjadi bentuk tulisan, misalnya 45 menjadi "empat puluh lima".

    ### 3. Buatlah program yang dapat memberikan input dan output sbb.

    **Input: 3**

    **Output:**
    ```text
    3 2 1 * 1 2 3
    2 1 * 1 2
        1 * 1
        *

    ```C++
    source code unguided 3
    #include <iostream>
using namespace std;

int main(){
    int n;

    cout << "Input: ";
    cin >> n;

    for(int i = n; i >= 1; i--){
        for(int j = n; j > i; j--)
            cout << "  ";

        for(int j = i; j >= 1; j--)
            cout << j << " ";

        cout << "* ";

        for(int j = 1; j <= i; j++)
            cout << j << " ";

        cout << endl;
    }

    for(int j = 0; j < n; j++)
        cout << "  ";

    cout << "*";

    return 0;
}
    ```
    ### Output Unguided 3 :

    ##### Output 1

    ![Screenshot Output Unguided 3_1](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%202%20-%20Modul%201/Unguided%203/Output-1.png)

    ##### Output 2

    ![Screenshot Output Unguided 3_2](https://github.com/hilmannnnn22/109082500027_Hilman-Aniq-Almunawa/blob/main/Semester%203/Week%202%20-%20Modul%201/Unguided%203/Output-2.png)

    #### penjelasan unguided 3
    Pada program ini saya membuat pola angka menggunakan perulangan for. Program menerima input berupa nilai n yang digunakan untuk menentukan jumlah baris. Pada setiap baris, program menampilkan angka dari nilai terbesar menuju 1, kemudian tanda * dan angka dari 1 sampai nilai tersebut. Spasi digunakan untuk membuat posisi pola semakin menjorok ke kanan pada setiap baris.

    ## Kesimpulan
    Pada praktikum Modul 1 ini saya mempelajari dasar-dasar bahasa C++ seperti tipe data, variabel, input dan output, operator, kondisi, dan perulangan. Dari program yang dibuat, saya dapat memahami penggunaan `cin`, `cout`, `if`, `switch`, dan `for` untuk membuat program sederhana sesuai dengan permasalahan yang diberikan.

    ## Referensi
    [1] Triase. (2020). Diktat Edisi Revisi: STRUKTUR DATA. Medan: Universitas Islam Negeri Sumatera Utara Medan.
    <br>...
