package main

import (
	"fmt"
)

const N int = 25

type kantin struct {
	namaMakanan  string
	hargaMakanan int
}

type kantinA [N]kantin
type kantinB [N]kantin

var dataA [N]kantin
var dataB [N]kantin

func makananKantinA() {
	dataA[0].namaMakanan = "AyamGoreng"
	dataA[0].hargaMakanan = 10000
	dataA[1].namaMakanan = "AyamBakar"
	dataA[1].hargaMakanan = 10000
	dataA[2].namaMakanan = "MieAceh"
	dataA[2].hargaMakanan = 15000
	dataA[3].namaMakanan = "GulaiKepalaIkan"
	dataA[3].hargaMakanan = 16000
	dataA[4].namaMakanan = "SotoMedan"
	dataA[4].hargaMakanan = 18000
	dataA[5].namaMakanan = "IkanAsamPedas"
	dataA[5].hargaMakanan = 18000
	dataA[6].namaMakanan = "SotoPadang"
	dataA[6].hargaMakanan = 16000
	dataA[7].namaMakanan = "Pempek"
	dataA[7].hargaMakanan = 10000
	dataA[8].namaMakanan = "NasiUduk"
	dataA[8].hargaMakanan = 8000
	dataA[9].namaMakanan = "SotoBetawi"
	dataA[9].hargaMakanan = 12000
}

func makananKantinB() {
	dataB[0].namaMakanan = "Gado-Gado"
	dataB[0].hargaMakanan = 12000
	dataB[1].namaMakanan = "GuramePecak"
	dataB[1].hargaMakanan = 16000
	dataB[2].namaMakanan = "EmpalGentong"
	dataB[2].hargaMakanan = 14000
	dataB[3].namaMakanan = "GulaiKepalaIkan"
	dataB[3].hargaMakanan = 20000
	dataB[4].namaMakanan = "SopBuntut"
	dataB[4].hargaMakanan = 18000
	dataB[5].namaMakanan = "NasiGorengBabat"
	dataB[5].hargaMakanan = 18000
	dataB[6].namaMakanan = "LontongOpor"
	dataB[6].hargaMakanan = 16000
	dataB[7].namaMakanan = "SotoAyam"
	dataB[7].hargaMakanan = 13000
	dataB[8].namaMakanan = "RujakCingur"
	dataB[8].hargaMakanan = 25000
	dataB[9].namaMakanan = "CotoMakassar"
	dataB[9].hargaMakanan = 15000
}

func menuAwal() {

	var input int

	makananKantinA() // Inisialisasi array
	makananKantinB()

	fmt.Println("============================================================")
	fmt.Println("---------------------- SELAMAT DATANG ----------------------")
	fmt.Println("============================================================")
	fmt.Println("Silakan Pilih Kantin")
	fmt.Println()
	fmt.Println("1. Kantin A")
	fmt.Println("2. Kantin B")
	fmt.Println("0. Tutup Pogram")
	fmt.Println()
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println()
		menuKantinA()
	case 2:
		fmt.Println()
		menuKantinB()
	case 0:
		fmt.Println()
		fmt.Println("============================================================")
		fmt.Println("----------------------- TERIMA KASIH -----------------------")
		fmt.Println("============================================================")
		fmt.Println()
	default:
		fmt.Println()
		fmt.Println("Pilihan Yang Anda Masukkan Salah.")
		fmt.Println("Silakan Ulangi lagi.")
		fmt.Println()
		menuAwal()
	}
}

func menuKantinA() {

	var input int

	fmt.Println("============================================================")
	fmt.Println("----------------------- Menu Kantin A ----------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Println("1. Lihat Daftar Makanan")
	fmt.Println("2. Cari Makanan")
	fmt.Println("3. Input Makanan")
	fmt.Println("4. Edit Makanan")
	fmt.Println("5. Delete Makanan")
	fmt.Println("0. Kembali ke Menu Awal")
	fmt.Println()
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println()
		daftarlistA()
	case 2:
		fmt.Println()
		searchA()
	case 3:
		fmt.Println()
		inputA()
	case 4:
		fmt.Println()
		editA()
	case 5:
		fmt.Println()
		deleteA()
	case 0:
		fmt.Println()
		menuAwal()
	default:
		fmt.Println()
		fmt.Println("Pilihan Yang Anda Masukkan Salah.")
		fmt.Println("Silakan Ulangi lagi.")
		fmt.Println()
		menuKantinA()
	}
}

func daftarlistA() {

	var input int

	fmt.Println("============================================================")
	fmt.Println("------------------ Daftar Makanan Kantin A -----------------")
	fmt.Println("============================================================")
	fmt.Println()

	for i := 0; i < N; i++ {
		if dataA[i].namaMakanan != "" {
			fmt.Println(dataA[i].namaMakanan, "........................................", dataA[i].hargaMakanan)
		}
	}

	fmt.Println()
	fmt.Println("1. Sorting Nama A - Z")
	fmt.Println("2. Sorting Harga Termurah")
	fmt.Println("0. Kembali ke Menu Kantin")
	fmt.Println()
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println()
		dataA = sortNamaA(dataA)
	case 2:
		fmt.Println()
		dataA = sortHargaA(dataA)
	case 0:
		fmt.Println()
		menuKantinA()
	default:
		fmt.Println()
		fmt.Println("Pilihan Yang Anda Masukkan Salah.")
		fmt.Println("Silakan Ulangi lagi.")
		fmt.Println()
		daftarlistA()
	}

	for i := 0; i < N; i++ {
		if dataA[i].namaMakanan != "" {
			fmt.Println(dataA[i].namaMakanan, "........................................", dataA[i].hargaMakanan)
		}
	}
	menuKantinA()
}

func sortNamaA(dataA kantinA) kantinA {

	fmt.Println("============================================================")
	fmt.Println("-------------------- Berdasarkan Nama A-Z ------------------")
	fmt.Println("============================================================")
	fmt.Println()

	for i := 0; i < N-1; i++ {
		for j := 0; j < N-i-1; j++ {
			if dataA[j].namaMakanan > dataA[j+1].namaMakanan {
				t := dataA[j]
				dataA[j] = dataA[j+1]
				dataA[j+1] = t
			}
		}
	}
	return dataA
}

func sortHargaA(dataA kantinA) kantinA {

	fmt.Println("============================================================")
	fmt.Println("---------------- Berdasarkan Harga Termurah ----------------")
	fmt.Println("============================================================")
	fmt.Println()

	for i := 0; i < N-1; i++ {
		for j := 0; j < N-i-1; j++ {
			if dataA[j].hargaMakanan > dataA[j+1].hargaMakanan {
				t := dataA[j]
				dataA[j] = dataA[j+1]
				dataA[j+1] = t
			}
		}
	}
	return dataA
}

func searchA() {

	var ada bool
	var makanan string

	fmt.Println("============================================================")
	fmt.Println("------------------- Menu Mencari Makanan -------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Makanan Yang di Cari: ")
	fmt.Scan(&makanan)

	ada = false

	for i := 0; i < N; i++ {
		if dataA[i].namaMakanan == makanan {
			fmt.Println()
			fmt.Println(dataA[i].namaMakanan, " ........................................ ", dataA[i].hargaMakanan)
			ada = true
		}
	}

	if ada == false {
		fmt.Println()
		fmt.Println("Maaf, Makanan Tidak Ada di Daftar")
	}
	fmt.Println()
	menuKantinA()
}

func inputA() {

	var input string

	fmt.Println("============================================================")
	fmt.Println("-------------------- Menu Input Makanan --------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Input? (y/n): ")
	fmt.Scan(&input)

	for i := 10; i < N && input == "y"; i++ {
		fmt.Print("Masukkan Nama Makanan: ")
		fmt.Scan(&dataA[i].namaMakanan)
		fmt.Print("Masukkan Harga: ")
		fmt.Scan(&dataA[i].hargaMakanan)
		fmt.Print("Mau Input Lagi? (y/n): ")
		fmt.Scan(&input)
	}
	fmt.Println()
	menuKantinA()
}

func editA() {

	var ada bool
	var makanan string
	var harga int

	fmt.Println("============================================================")
	fmt.Println("--------------------- Menu Edit Makanan --------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Nama Makanan Yang Akan di Edit: ")
	fmt.Scan(&makanan)

	ada = false

	for i := 0; i < N; i++ {
		if dataA[i].namaMakanan == makanan {
			fmt.Print("Masukkan NAMA Makanan Baru: ")
			fmt.Scan(&makanan)
			dataA[i].namaMakanan = makanan
			fmt.Print("Masukkan HARGA Makanan Baru: ")
			fmt.Scan(&harga)
			dataA[i].hargaMakanan = harga
			ada = true
		}
	}

	if ada == false {
		fmt.Println()
		fmt.Println("Maaf, Makanan Tidak Ada di Daftar")
	}
	fmt.Println()
	menuKantinA()
}

func deleteA() {

	var ada bool
	var makanan string

	fmt.Println("============================================================")
	fmt.Println("-------------------- Menu Hapus Makanan --------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Makanan Yang Akan di Hapus: ")
	fmt.Scan(&makanan)

	ada = false

	for i := 0; i < N; i++ {
		if dataA[i].namaMakanan == makanan {
			fmt.Println()
			fmt.Println("Makanan Telah Terhapus.")
			dataA[i].namaMakanan = ""
			dataA[i].hargaMakanan = 0
			ada = true
		}
	}

	if ada == false {
		fmt.Println()
		fmt.Println("Maaf, Makanan Tidak Ada di Daftar")
	}
	fmt.Println()
	menuKantinA()
}

func menuKantinB() {

	var input int

	fmt.Println("============================================================")
	fmt.Println("----------------------- Menu Kantin B ----------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Println("1. Lihat Daftar Makanan")
	fmt.Println("2. Cari Makanan")
	fmt.Println("3. Input Makanan")
	fmt.Println("4. Edit Makanan")
	fmt.Println("5. Delete Makanan")
	fmt.Println("0. Kembali ke Menu Awal")
	fmt.Println()
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println()
		daftarlistB()
	case 2:
		fmt.Println()
		searchB()
	case 3:
		fmt.Println()
		inputB()
	case 4:
		fmt.Println()
		editB()
	case 5:
		fmt.Println()
		deleteB()
	case 0:
		fmt.Println()
		menuAwal()
	default:
		fmt.Println()
		fmt.Println("Pilihan Yang Anda Masukkan Salah.")
		fmt.Println("Silakan Ulangi lagi.")
		fmt.Println()
		menuKantinB()
	}
}

func daftarlistB() {

	var input int

	fmt.Println("============================================================")
	fmt.Println("------------------ Daftar Makanan Kantin B -----------------")
	fmt.Println("============================================================")
	fmt.Println()

	for i := 0; i < N; i++ {
		if dataB[i].namaMakanan != "" {
			fmt.Println(dataB[i].namaMakanan, "........................................", dataB[i].hargaMakanan)
		}
	}

	fmt.Println()
	fmt.Println("1. Sorting Nama A - Z")
	fmt.Println("2. Sorting Harga Termurah")
	fmt.Println("0. Kembali ke Menu Kantin")
	fmt.Println()
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println()
		dataB = sortNamaB(dataB)
	case 2:
		fmt.Println()
		dataB = sortHargaB(dataB)
	case 0:
		fmt.Println()
		menuKantinB()
	default:
		fmt.Println()
		fmt.Println("Pilihan Yang Anda Masukkan Salah.")
		fmt.Println("Silakan Ulangi lagi.")
		fmt.Println()
		daftarlistB()
	}

	for i := 0; i < N; i++ {
		if dataB[i].namaMakanan != "" {
			fmt.Println(dataB[i].namaMakanan, "........................................", dataB[i].hargaMakanan)
		}
	}
	menuKantinB()
}

func sortNamaB(dataB kantinB) kantinB {

	fmt.Println("============================================================")
	fmt.Println("-------------------- Berdasarkan Nama A-Z ------------------")
	fmt.Println("============================================================")
	fmt.Println()

	for i := 0; i < N-1; i++ {
		for j := 0; j < N-i-1; j++ {
			if dataB[j].namaMakanan > dataB[j+1].namaMakanan {
				t := dataB[j]
				dataB[j] = dataB[j+1]
				dataB[j+1] = t
			}
		}
	}
	return dataB
}

func sortHargaB(dataB kantinB) kantinB {

	fmt.Println("============================================================")
	fmt.Println("---------------- Berdasarkan Harga Termurah ----------------")
	fmt.Println("============================================================")
	fmt.Println()

	for i := 0; i < N-1; i++ {
		for j := 0; j < N-i-1; j++ {
			if dataB[j].hargaMakanan > dataB[j+1].hargaMakanan {
				t := dataB[j]
				dataB[j] = dataB[j+1]
				dataB[j+1] = t
			}
		}
	}
	return dataB
}

func searchB() {

	var ada bool
	var makanan string

	fmt.Println("============================================================")
	fmt.Println("------------------- Menu Mencari Makanan -------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Makanan Yang di Cari: ")
	fmt.Scan(&makanan)

	ada = false

	for i := 0; i < N; i++ {
		if dataB[i].namaMakanan == makanan {
			fmt.Println()
			fmt.Println(dataB[i].namaMakanan, " ........................................ ", dataB[i].hargaMakanan)
			ada = true
		}
	}

	if ada == false {
		fmt.Println()
		fmt.Println("Maaf, Makanan Tidak Ada di Daftar")
	}
	fmt.Println()
	menuKantinB()
}

func inputB() {

	var input string

	fmt.Println("============================================================")
	fmt.Println("-------------------- Menu Input Makanan --------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Input? (y/n): ")
	fmt.Scan(&input)

	for i := 10; i < N && input == "y"; i++ {
		fmt.Print("Masukkan Nama Makanan: ")
		fmt.Scan(&dataB[i].namaMakanan)
		fmt.Print("Masukkan Harga: ")
		fmt.Scan(&dataB[i].hargaMakanan)
		fmt.Print("Mau Input Lagi? (y/n): ")
		fmt.Scan(&input)
	}
	fmt.Println()
	menuKantinB()
}

func editB() {

	var ada bool
	var makanan string
	var harga int

	
	fmt.Println("============================================================")
	fmt.Println("--------------------- Menu Edit Makanan --------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Nama Makanan Yang Akan di Edit: ")
	fmt.Scan(&makanan)

	ada = false

	for i := 0; i < N; i++ {
		if dataB[i].namaMakanan == makanan {
			fmt.Print("Masukkan NAMA Makanan Baru: ")
			fmt.Scan(&makanan)
			dataB[i].namaMakanan = makanan
			fmt.Print("Masukkan HARGA Makanan Baru: ")
			fmt.Scan(&harga)
			dataB[i].hargaMakanan = harga
			ada = true
		}
	}

	if ada == false {
		fmt.Println()
		fmt.Println("Maaf, Makanan Tidak Ada di Daftar")
	}
	fmt.Println()
	menuKantinB()
}

func deleteB() {

	var ada bool
	var makanan string

	fmt.Println("============================================================")
	fmt.Println("-------------------- Menu Hapus Makanan --------------------")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Print("Masukkan Makanan Yang Akan di Hapus: ")
	fmt.Scan(&makanan)

	ada = false

	for i := 0; i < N; i++ {
		if dataB[i].namaMakanan == makanan {
			fmt.Println()
			fmt.Println("Makanan Telah Terhapus.")
			dataB[i].namaMakanan = ""
			dataB[i].hargaMakanan = 0
			ada = true
		}
	}

	if ada == false {
		fmt.Println()
		fmt.Println("Maaf, Makanan Tidak Ada di Daftar")
	}
	fmt.Println()
	menuKantinB()
}

func main() {
	menuAwal()
}