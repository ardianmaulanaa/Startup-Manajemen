package main

import (
	"fmt"
)

const NMAXGT = 1024

type mAnggota struct {
	namaAnggota  string
	peranAnggota string
}

type tim [NMAXGT]mAnggota

const NMAX int = 1024

type mStartup struct {
	namaStrtup  string
	bidangUsaha string
	produk      string
	pendanaan   int
	timAnggota  tim
	jmlhAnggota int
	jumStartup  int
	thnBerdiri  int
}

type Startup [NMAX - 1]mStartup

var A Startup
var i int
var StrtUp, bidangUsaha, produk, kode string
var pendanaan, jumStartup, thnBerdiri, opsi, pilih, x int
var found bool

func main() {
	var pilMenu int
	var logIn, pass string
	var akses bool = false
	fmt.Println("--- Selamat datang ---")
	for akses != true {
		fmt.Print("Username: ")
		fmt.Scan(&logIn)
		fmt.Print("Password: ")
		fmt.Scan(&pass)
		if (logIn == "A" && pass == "A") || (logIn == "Alfan" && pass == "AlfanTelkom") {
			akses = true
			for pilMenu != 11 {
				tampilMenu(&pilMenu)
			}
		} else {
			fmt.Println("Username atau password salah. Silahkan coba lagi!")
		}
	}
}

func tampilMenu(pilihan *int) {
	fmt.Println("===========================================")
	fmt.Println("|               MENU UTAMA                |")
	fmt.Println("===========================================")
	fmt.Println("|  1. Menambah Startup                    |")
	fmt.Println("|  2. Mengubah Startup                    |")
	fmt.Println("|  3. Tampilkan Startup                   |")
	fmt.Println("|  4. Menghapus Startup                   |")
	fmt.Println("|  5. Menambah Anggota                    |")
	fmt.Println("|  6. Tampilkan Anggota                   |")
	fmt.Println("|  7. Pencarian Startup                   |")
	fmt.Println("|  8. Pengurutan Startup                  |")
	fmt.Println("|  9. Laporan Startup (Kategori Usaha)    |")
	fmt.Println("| 10. Menghapus Anggota                   |")
	fmt.Println("| 11. Keluar                              |")
	fmt.Println("===========================================")
	fmt.Print("Pilih opsi (1-11): ")

	fmt.Scan(pilihan)

	switch *pilihan {
	case 1:
		tambahStartup()
	case 2:
		mengubahStartup()
	case 3:
		tampilkanStartup()
	case 4:
		menghapusStartup(&opsi, &kode)
	case 5:
		tambahAnggota()
	case 6:
		tampilkanAnggota()
	case 7:
		idx := mencariStartup(&pilih)
		if idx != -1 {
			fmt.Println("\nData startup ditemukan:")
			fmt.Println("===============================================================")
			fmt.Printf("Nama Startup   : %s\n", A[idx].namaStrtup)
			fmt.Printf("Bidang Usaha   : %s\n", A[idx].bidangUsaha)
			fmt.Printf("Produk         : %s\n", A[idx].produk)
			fmt.Printf("Pendanaan      : %d\n", A[idx].pendanaan)
			fmt.Printf("Tahun Berdiri  : %d\n", A[idx].thnBerdiri)
			fmt.Println("===============================================================")
		} else {
			fmt.Println("Data Startup tidak ditemukan")
		}
	case 8:
		pengurutanStartup(&x)
	case 9:
		laporanKategoriBidush()
	case 10:
		menghapusAnggota()
	}
}
func tambahStartup() {
	fmt.Print("Masukkan jumlah Startup yang ingin ditambahkan: ")
	fmt.Scan(&jumStartup)

	for j := 0; j < jumStartup; j++ {
		fmt.Println()
		fmt.Println("===============================================================================")
		fmt.Printf("                          Data Startup ke-%d\n", j+1)
		fmt.Println("===============================================================================")

		fmt.Print("Nama Startup             : ")
		fmt.Scan(&StrtUp)
		A[i].namaStrtup = StrtUp

		fmt.Print("Tahun Berdiri            : ")
		fmt.Scan(&thnBerdiri)
		A[i].thnBerdiri = thnBerdiri

		fmt.Print("Bidang Usaha             : ")
		fmt.Scan(&bidangUsaha)
		A[i].bidangUsaha = bidangUsaha

		fmt.Print("Nama Produk              : ")
		fmt.Scan(&produk)
		A[i].produk = produk

		fmt.Print("Jumlah Pendanaan         : ")
		fmt.Scan(&pendanaan)
		A[i].pendanaan = pendanaan

		i++
	}

}

func mengubahStartup() {
	var newStartup, newProduk, newBidangUsaha string
	var newPendanaan, newThnBerdiri int
	var idx int
	tampilkanStartup()
	if i == 0 {
		fmt.Println("Belum ada Startup yang dapat diubah.")
	} else {
		fmt.Print("Masukkan ID Startup yang ingin diubah: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= i {
			fmt.Println()
			fmt.Println("================================================================")
			fmt.Printf("            Mengubah Data Startup dengan ID: %d\n", idx)
			fmt.Println("================================================================")

			fmt.Print("Nama Baru Startup         : ")
			fmt.Scan(&newStartup)
			A[idx-1].namaStrtup = newStartup

			fmt.Print("Tahun Berdiri Baru        : ")
			fmt.Scan(&newThnBerdiri)
			A[idx-1].thnBerdiri = newThnBerdiri

			fmt.Print("Bidang Usaha Baru         : ")
			fmt.Scan(&newBidangUsaha)
			A[idx-1].bidangUsaha = newBidangUsaha

			fmt.Print("Produk Baru               : ")
			fmt.Scan(&newProduk)
			A[idx-1].produk = newProduk

			fmt.Print("Total Pendanaan Baru      : ")
			fmt.Scan(&newPendanaan)
			A[idx-1].pendanaan = newPendanaan

			fmt.Println()
			fmt.Println("================================================================")
			fmt.Println("               Data Startup berhasil diubah.")
			fmt.Println("================================================================")
		} else {
			fmt.Println("ID Startup tidak ditemukan dalam daftar!")
		}
	}
}

func tampilkanStartup() {
	fmt.Printf("\nDaftar StartUp yang telah ditambahkan:\n")
	fmt.Println("=======================================================================================")
	fmt.Printf("| %-3s | %-20s | %-15s | %-15s | %-10s | %-5s |\n",
		"ID", "Nama StartUp", "Bidang Usaha", "Produk", "Pendanaan", "Tahun")
	fmt.Println("=======================================================================================")

	for j := 0; j < i; j++ {
		fmt.Printf("| %-3d | %-20s | %-15s | %-15s | %-10d | %-5d |\n",
			j+1,
			A[j].namaStrtup,
			A[j].bidangUsaha,
			A[j].produk,
			A[j].pendanaan,
			A[j].thnBerdiri)
	}
	fmt.Println("=======================================================================================")
}

func menghapusStartup(opsi *int, kode *string) {
	var idx int

	tampilkanStartup()
	fmt.Println("==============================================================================")
	fmt.Println(" Pilih opsi penghapusan:")
	fmt.Println(" 1. Hapus seluruh startup berdasarkan ID")
	fmt.Println(" 2. Hapus data tertentu dari startup")
	fmt.Println("==============================================================================")
	fmt.Print("Masukkan pilihan (1/2): ")
	fmt.Scan(opsi)

	if *opsi == 1 {
		if i == 0 {
			fmt.Println("Tidak ada startup yang dapat dihapus.")
		} else {
			fmt.Print("Masukkan ID startup yang ingin dihapus: ")
			fmt.Scan(&idx)

			if idx >= 1 && idx <= i {
				for j := idx - 1; j < i-1; j++ {
					A[j] = A[j+1]
				}
				A[i-1] = mStartup{}
				i--
				fmt.Println("Startup berhasil dihapus.")
			} else {
				fmt.Println("ID tidak valid. Silakan coba lagi.")
			}
		}
	} else if *opsi == 2 {
		if i == 0 {
			fmt.Println("Tidak ada startup untuk diedit.")
		} else {
			fmt.Println("==============================================================================")
			fmt.Println(" Data yang dapat dihapus:")
			fmt.Println(" A. Nama Startup")
			fmt.Println(" B. Tipe/Bidang Usaha")
			fmt.Println(" C. Produk")
			fmt.Println(" D. Dana Anggaran")
			fmt.Println(" E. Tahun Berdiri")
			fmt.Println("==============================================================================")

			tampilkanStartup()

			fmt.Print("Masukkan ID startup yang datanya ingin dihapus: ")
			fmt.Scan(&idx)

			if idx >= 1 && idx <= i {
				fmt.Print("Pilih data yang ingin dihapus (A-E): ")
				fmt.Scan(kode)

				switch *kode {
				case "A", "a":
					A[idx-1].namaStrtup = ""
					fmt.Println("Nama Startup berhasil dihapus.")
				case "B", "b":
					A[idx-1].bidangUsaha = ""
					fmt.Println("Bidang Usaha berhasil dihapus.")
				case "C", "c":
					A[idx-1].produk = ""
					fmt.Println("Produk berhasil dihapus.")
				case "D", "d":
					A[idx-1].pendanaan = 0
					fmt.Println("Dana Anggaran berhasil dihapus.")
				case "E", "e":
					A[idx-1].thnBerdiri = 0
					fmt.Println("Tahun Berdiri berhasil dihapus.")
				default:
					fmt.Println("Pilihan tidak valid. Gunakan huruf A-E.")
				}
			} else {
				fmt.Println("ID tidak ditemukan.")
			}
		}
	} else {
		fmt.Println("Opsi tidak dikenal. Silakan pilih 1 atau 2.")
	}
}

func tambahAnggota() {
	var idx int
	var anggotaBaru mAnggota

	fmt.Println("==============================================================================")
	fmt.Println("                        Tambah Anggota ke StartUp")
	fmt.Println("==============================================================================")

	if i == 0 {
		fmt.Println("Belum ada StartUp yang tersedia.")
		fmt.Println("Silakan tambahkan StartUp terlebih dahulu sebelum menambah anggota.")
	} else {
		tampilkanStartup()
		fmt.Print("Masukkan ID StartUp yang akan ditambahkan anggotanya: ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= i {
			if A[idx-1].jmlhAnggota < NMAXGT {
				fmt.Println("------------------------------------------------------------------------------")
				fmt.Println("                    Masukkan Informasi Anggota Baru")
				fmt.Println("------------------------------------------------------------------------------")
				fmt.Print("Nama Anggota      : ")
				fmt.Scan(&anggotaBaru.namaAnggota)
				fmt.Print("Peran Anggota     : ")
				fmt.Scan(&anggotaBaru.peranAnggota)

				A[idx-1].timAnggota[A[idx-1].jmlhAnggota] = anggotaBaru
				A[idx-1].jmlhAnggota++

				fmt.Println("------------------------------------------------------------------------------")
				fmt.Println("Anggota berhasil ditambahkan ke StartUp:", A[idx-1].namaStrtup)
			} else {
				fmt.Println("Jumlah anggota sudah mencapai batas maksimum!")
			}
		} else {
			fmt.Println("ID StartUp tidak valid. Silakan coba lagi.")
		}
	}

	fmt.Println("==============================================================================")
}

func menghapusAnggota() {
	var idx int
	var namaAnggotaHapus string
	var pos int
	var found bool = false

	fmt.Println("===== Hapus Anggota Startup =====")

	tampilkanStartup()

	fmt.Print("Masukkan ID startup yang anggotanya ingin dihapus: ")
	fmt.Scan(&idx)

	if i != 0 && idx >= 1 && idx <= i {
		if A[idx-1].jmlhAnggota > 0 {
			fmt.Printf("\nDaftar anggota untuk Startup '%s':\n", A[idx-1].namaStrtup)
			fmt.Println("============================================================")
			fmt.Printf("| %-3s | %-20s | %-20s |\n", "No", "Nama Anggota", "Peran")
			fmt.Println("============================================================")
			for j := 0; j < A[idx-1].jmlhAnggota; j++ {
				fmt.Printf("| %-3d | %-20s | %-20s |\n", j+1, A[idx-1].timAnggota[j].namaAnggota, A[idx-1].timAnggota[j].peranAnggota)
			}
			fmt.Println("============================================================")

			fmt.Print("Masukkan nama anggota yang ingin dihapus: ")
			fmt.Scan(&namaAnggotaHapus)

			found = false
			for j := 0; j < A[idx-1].jmlhAnggota; j++ {
				if A[idx-1].timAnggota[j].namaAnggota == namaAnggotaHapus && !found {
					found = true
					pos = j
				}
			}

			if found {
				for k := pos; k < A[idx-1].jmlhAnggota-1; k++ {
					A[idx-1].timAnggota[k] = A[idx-1].timAnggota[k+1]
				}
				A[idx-1].jmlhAnggota--
				fmt.Println("Anggota berhasil dihapus")
			} else {
				fmt.Println("Nama anggota tidak ditemukan")
			}
		} else {
			fmt.Println("Startup ini belum memiliki anggota.")
		}
	} else {
		fmt.Println("Data startup tidak ditemukan.")
	}
}

func tampilkanAnggota() {
	var idx int
	tampilkanStartup()
	if i == 0 {
		fmt.Println("Startup belum tersedia, anggota tidak dapat diketahui.")
	} else {
		fmt.Print("Masukkan ID Startup yang ingin dilihat anggotanya: ")
		fmt.Scan(&idx)
		if idx >= 1 && idx <= i {
			if A[idx-1].jmlhAnggota == 0 {
				fmt.Println("Belum ada anggota yang tersedia")
			} else {
				fmt.Printf("\nDaftar anggota untuk Startup '%s':\n", A[idx-1].namaStrtup)
				fmt.Println("============================================================")
				fmt.Printf("| %-3s | %-20s | %-20s |\n", "No", "Nama Anggota", "Peran")
				fmt.Println("============================================================")
				for j := 0; j < A[idx-1].jmlhAnggota; j++ {
					fmt.Printf("| %-3d | %-20s | %-20s |\n",
						j+1,
						A[idx-1].timAnggota[j].namaAnggota,
						A[idx-1].timAnggota[j].peranAnggota)
				}
				fmt.Println("============================================================")
			}
		} else {
			fmt.Println("ID Startup tidak ditemukan!")
		}
	}
}

func mencariStartup(pilihan *int) int {
	var X string
	var found int = -1
	var low, mid, high, idx int

	fmt.Println("==============================================================================")
	fmt.Println("                         PENCARIAN STARTUP")
	fmt.Println("==============================================================================")
	fmt.Println("Pilih Metode Pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Masukkan pilihan (1/2): ")
	fmt.Scan(pilihan)

	switch *pilihan {
	case 1:
		fmt.Print("Masukkan nama startup atau bidang usaha (Sequential Search): ")
		fmt.Scan(&X)

		for idx < i && found == -1 {
			if A[idx].namaStrtup == X || A[idx].bidangUsaha == X {
				found = idx
			}
			idx++
		}
	case 2:
		fmt.Print("Masukkan nama startup yang ingin dicari (Binary Search): ")
		fmt.Scan(&X)

		low = 0
		high = i - 1

		for low <= high && found == -1 {
			mid = (low + high) / 2
			if A[mid].namaStrtup == X {
				found = mid
			} else if A[mid].namaStrtup < X {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih 1 atau 2.")
	}

	return found
}

func pengurutanStartup(x *int) {
	var idx, idx_min int
	var methods string
	fmt.Println("METHODS")
	fmt.Println("1. Selection sort")
	fmt.Println("2. Insertion sort")
	fmt.Print("Masukkan metode yang diinginkan: ")
	fmt.Scan(x)
	switch *x {
	case 1:
		fmt.Println("1. Descending")
		fmt.Println("2. Ascending")
		fmt.Print("Masukkan urutan yang diinginkan: ")
		fmt.Scan(&methods)
		if methods == "Descending" || methods == "descending" || methods == "1" {
			fmt.Print("Masukkan berapa jumlah Startup yang ingin diurutkan: ")
			fmt.Scan(&idx)
			if idx >= 1 && idx <= jumStartup {
				i1 := 0
				for i1 < idx-1 {
					idx_min = i1
					j := i1 + 1
					for j < idx {
						if A[j].namaStrtup > A[idx_min].namaStrtup {
							idx_min = j
						}
						j++
					}
					A[i1], A[idx_min] = A[idx_min], A[i1]
					i1++
				}
				fmt.Println("Startup berhasil diurutkan dengan Selection Sort berdasarkan nama.")
				tampilkanStartup()
			} else {
				fmt.Println("Data diluar jangkauan!")
			}
		} else if methods == "Ascending" || methods == "ascending" || methods == "2" {
			fmt.Print("Masukkan berapa jumlah Startup yang ingin diurutkan: ")
			fmt.Scan(&idx)
			if idx >= 1 && idx <= jumStartup {
				i1 := 0
				for i1 < idx-1 {
					idx_min = i1
					j := i1 + 1
					for j < idx {
						if A[j].namaStrtup < A[idx_min].namaStrtup {
							idx_min = j
						}
						j++
					}
					A[i1], A[idx_min] = A[idx_min], A[i1]
					i1++
				}
				fmt.Println("Startup berhasil diurutkan dengan Selection Sort berdasarkan nama.")
				tampilkanStartup()
			} else {
				fmt.Println("Data diluar jangkauan!")
			}
		}
	case 2:
		fmt.Println("1. Descending")
		fmt.Println("2. Ascending")
		fmt.Print("Masukkan urutan yang diinginkan: ")
		fmt.Scan(&methods)
		if methods == "Descending" || methods == "descending" || methods == "1" {
			fmt.Print("Masukkan berapa jumlah Startup yang ingin diurutkan: ")
			fmt.Scan(&idx)
			if idx >= 1 && idx <= jumStartup {
				i2 := 1
				for i2 <= idx-1 {
					j2 := i2
					temp := A[j2]
					for j2 > 0 && temp.namaStrtup > A[j2-1].namaStrtup {
						A[j2] = A[j2-1]
						j2--
					}
					A[j2] = temp
					i2++
				}
				fmt.Println("Startup berhasil diurutkan dengan Insertion Sort berdasarkan nama.")
				tampilkanStartup()
			} else {
				fmt.Println("Data diluar jangkauan!")
			}
		} else if methods == "Ascending" || methods == "ascending" || methods == "2" {
			fmt.Print("Masukkan berapa jumlah Startup yang ingin diurutkan: ")
			fmt.Scan(&idx)
			if idx >= 1 && idx <= jumStartup {
				i2 := 1
				for i2 <= idx-1 {
					j2 := i2
					temp := A[j2]
					for j2 > 0 && temp.namaStrtup < A[j2-1].namaStrtup {
						A[j2] = A[j2-1]
						j2--
					}
					A[j2] = temp
					i2++
				}
				fmt.Println("Startup berhasil diurutkan dengan Insertion Sort berdasarkan nama.")
				tampilkanStartup()
			} else {
				fmt.Println("Data diluar jangkauan!")
			}
		}
	}
}

func lapJumlahStartup(bidUsh int) {
	var idx int
	fmt.Print("Masukkan kategori Bidang Usaha: ")
	fmt.Scan(&bidUsh)
	if idx != -1 {
		fmt.Println("\nData startup ditemukan:")
		fmt.Println("===============================================================")
		fmt.Printf("Nama Startup   : %s\n", A[idx].namaStrtup)
		fmt.Printf("Bidang Usaha   : %s\n", A[idx].bidangUsaha)
		fmt.Printf("Produk         : %s\n", A[idx].produk)
		fmt.Printf("Pendanaan      : %d\n", A[idx].pendanaan)
		fmt.Printf("Tahun Berdiri  : %d\n", A[idx].thnBerdiri)
		fmt.Println("===============================================================")
	} else {
		fmt.Println("Data Startup tidak ditemukan")
	}
}

func laporanKategoriBidush() {
	var kategori string
	var count int
	fmt.Print("Masukkan kategori bidang usaha: ")
	fmt.Scan(&kategori)

	fmt.Println("\nKategori Data startup ditemukan berdasarkan kategori bidang usaha: ", kategori)
	for i := 0; i < len(A); i++ {
		if A[i].bidangUsaha == kategori {
			fmt.Println("===============================================================")
			fmt.Printf("Nama Startup   : %s\n", A[i].namaStrtup)
			fmt.Printf("Bidang Usaha   : %s\n", A[i].bidangUsaha)
			fmt.Printf("Produk         : %s\n", A[i].produk)
			fmt.Printf("Pendanaan      : %d\n", A[i].pendanaan)
			fmt.Printf("Tahun Berdiri  : %d\n", A[i].thnBerdiri)
			fmt.Println("===============================================================")
			count++
		}
	}
	fmt.Println("Jumlah Startup di kategori", kategori, ":", count)
}
