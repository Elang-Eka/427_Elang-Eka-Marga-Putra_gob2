kode peserta 	: 149368582100-427
nama lengkap 	: Elang Eka Marga Putra
Summary sesi 4
      => Materi interface, interface adalah sebuah tipe data pada bahasa Go yang merupakan kumpulan dari definisi satu atau lebih method. 
	   Kita dapat menggunakan tipe data interface jika kita telah mengimplementasikan method-method yang telah didefinisikan oleh 
	   interface tersebut melalui tipe data lainnya. 
	   contoh : 
		type shape interface{
			area() float64
		}
      => Materi empty interface, empty interface merupakan suatu tipe data yang dapat menerima segala tipe data pada bahasa Go. 
	   Maka dari itu, sebuah variable dengan tipe data empty interface dapat diberikan nilai dengan tipe data yang berbeda-beda.
      => Reflect digunakan untuk melakukan  inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.
	   Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer,
	   dan banyak lagi.Go menyediakan package reflect, berisikan banyak sekali fungsi untuk keperluan reflection.
	   Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang paling penting untuk diketahui, yaitu
	   reflect.ValueOf() dan reflect.TypeOf().
	   1. Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi yang
	      berhubungan dengan nilai pada variabel yang dicari..
	   2. Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi yang
		berhubungan dengan tipe data variabel yang dicari
      => Concurrency adalah mengeksekusi sebuah proses secara independen atau berhadapan dengan lebih dari satu tugas dalam waktu yang sama.
      => Goroutine merupakan sebuah thread ringan pada bahasa Go untuk melakukan concurrency. 
	   Jika dibandingkan dengan thread biasa, Goroutine memiliki ukuran thread yang jauh lebih ringan. 
	   Pada saat kita mengeksekusi sebuah Goroutine, maka satu Goroutine hanya membutuhkan 2kb memori saja, 
	   sedangkan satu thread biasa dapat menghabiskan 1-2mb memori.Goroutine bersifat asynchronous sehingga proses nya tidak 
	   saling tunggu dengan Goroutine lainnya.
      => Waitgroup merupakan sebuah struct dari package sync, yang digunakan untuk melakukan sinkronisasi terhadap goroutine.
