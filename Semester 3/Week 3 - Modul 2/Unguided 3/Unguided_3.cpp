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