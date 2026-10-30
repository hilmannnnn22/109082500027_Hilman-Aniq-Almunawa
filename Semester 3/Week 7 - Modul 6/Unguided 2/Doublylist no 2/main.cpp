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