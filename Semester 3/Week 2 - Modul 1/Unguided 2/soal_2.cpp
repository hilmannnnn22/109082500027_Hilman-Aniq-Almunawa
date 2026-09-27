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