package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	/* I.S. Data tersedia dalam piranti masukan
	   F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
	   Proses input selama karakter bukanlah TITIK dan n <= NMAX */
	var char rune
	*n = 0
	for *n < NMAX {
		// Membaca karakter satu per satu
		fmt.Scanf("%c", &char)

		// Berhenti jika menemukan karakter titik '.'
		if char == '.' {
			break
		}

		// Mengabaikan karakter newline/enter atau carriage return agar array bersih
		if char != '\n' && char != '\r' {
			(*t)[*n] = char
			(*n)++
		}
	}
}

func cetakArray(t tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. n karakter dalam array muncul di layar */
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. Urutan isi array t terbalik */
	for i := 0; i < n/2; i++ {
		temp := (*t)[i]
		(*t)[i] = (*t)[n-1-i]
		(*t)[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	/* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
	   dan false apabila sebaliknya. Petunjuk: Manfaatkan prosedur balikanArray */

	// Simpan salinan isi array asli sebelum dibalik
	var tAsli tabel = t

	// Balikkan array (t di sini adalah copy, jadi aman untuk diubah)
	balikanArray(&t, n)

	// Bandingkan elemen per elemen
	for i := 0; i < n; i++ {
		if tAsli[i] != t[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	// Isi array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &m)

	// Cek apakah array tersebut palindrome
	isPalindrom := palindrom(tab, m)

	fmt.Print("Reverse teks: ")
	// Balikkan isi array tab aslinya untuk dicetak
	balikanArray(&tab, m)

	// Cetak isi array tab yang sudah dibalik
	cetakArray(tab, m)

	// Tampilkan hasil pengecekan palindrom
	fmt.Printf("Palindrom: %v\n", isPalindrom)
}
