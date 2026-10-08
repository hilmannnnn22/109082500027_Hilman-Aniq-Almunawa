#ifndef MAHASISWA_H_INCLUDED //untuk ngecek apakah header sudah di pakai di folder lain atau belum
#define MAHASISWA_H_INCLUDED //membuat header mahasiswa.h

struct mahasiswa{
    char nim[10];
    int nilai1, nilai2;
};

void inputMhs (mahasiswa &m);
float rata2 (mahasiswa m);
#endif // MAHASISWA_H_INCLUDED