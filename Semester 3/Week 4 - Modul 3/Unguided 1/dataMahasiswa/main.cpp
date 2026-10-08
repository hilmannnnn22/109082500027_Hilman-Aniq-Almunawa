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