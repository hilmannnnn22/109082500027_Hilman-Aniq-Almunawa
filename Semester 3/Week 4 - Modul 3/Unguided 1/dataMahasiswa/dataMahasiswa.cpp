#include <iostream>
#include "dataMahasiswa.h"
using namespace std;

float hitungNilaiAkhir(mahasiswa m)
{
    float nilai;

    nilai = (0.3 * m.uts) + (0.4 * m.uas) + (0.3 * m.tugas);

    return nilai;
}