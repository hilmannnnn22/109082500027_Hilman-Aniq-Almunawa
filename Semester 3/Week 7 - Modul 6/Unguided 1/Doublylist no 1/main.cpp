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