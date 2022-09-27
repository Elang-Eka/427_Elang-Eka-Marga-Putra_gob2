kode peserta 	: 149368582100-427
nama lengkap 	: Elang Eka Marga Putra
Summary sesi 5
      => Channel merupakan sebuah mekanisme untuk Goroutine saling berkomunikasi dengan Goroutine lainnya. 
	   Maksud berkomunikasi disini adalah proses pengiriman dan pertukaran data antara Goroutine satu dengan Goroutine lainnya.
      => Defer merupakan sebuah keyword pada bahasa Go yang digunakan untuk memanggil sebuah function yang dimana
	   proses eksekusi akan di tahan hingga block sebuah function selesai.
      => Funcion exit yang yang berasal dari package os berguna untuk menghentikan suatu program secara paksa.
      => Error merupakan sebuah tipe data pada bahasa Go yang digunakan untuk me-return sebuah error jika ada error yang terjadi. 
	   Umumya, nilai error akan di-return pada pada posisi terakhir dari nilai-nilai return sebuah function.
      => Panic digunakan untuk menampilkan stack traceerror sekaligus menghentikan flow goroutine(karena main() juga merupakan goroutine,
	   makabehaviour yang sama juga berlaku). Setelah adapanic, proses akan terhenti, apapun setelah tidakdi-eksekusi kecuali proses 
	   yang sudah di-defersebelumnya (akan muncul sebelum panic error).
      => Function recover digunakan untuk menangkaperror yang terjadi pada sebuah Goroutine.
