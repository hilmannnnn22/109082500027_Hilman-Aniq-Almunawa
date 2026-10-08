#ifndef DATA_MAHASISWA_H
#define DATA_MAHASISWA_H

struct mahasiswa {
    char nama[30];
    char nim[20];
    float uts, uas, tugas, nilaiAkhir;
};

float hitungNilaiAkhir(mahasiswa m);

#endif // DATA_MAHASISWA_H