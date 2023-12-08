<h1 align="center">Smart Dormitory</h1>
<h2 align="center">Prototipe Tugas Besar II3120 Layanan STI</h2>

## Anggota
**Kelompok 5 K1**

| NIM      | Nama                    |
|----------|-------------------------|
| 18221055 | Mochamad Syahrial A.	   |
| 18221065 | Josua Adriel Sinabuntar |
| 18221077 | Riandra Diva Auzan	     |
| 18221101 | Ilmagita Nariswari	     |
| 18221113 |Abraham Megantoro S.	   |

## Deskripsi

Repositori ini bertujuan untuk menyimpan *source code* untuk *prototype* dari Smart Dormitory untuk kebutuhan Tugas Besar II3120 Layanan STI. Smart Dormitory adalah fasilitas hunian dormitory yang telah ditingkatkan melalui solusi layanan IT (Information Technology). Prototipe ini menerapkan proses bisnis **penerimaan penghuni baru** menurut *role* admin sebagaimana yang terdapat di dalam proses bisnis Smart Dormitory.

## Panduan

### Teknologi yang diperlukan
1. Browser (Chrome, Firefox, dsb.)
2. Internet
3. Perangkat yang terhubung ke internet

### Login
1. Buka [http](https://smart-dormitory.vercel.app)https://smart-dormitory.vercel.app
2. Klik 'Login' untuk masuk ke laman login. Masukkan *username* dan *password* yaitu:
   
   ```
   username: admin
   password: admin
   ```

### Seleksi Penghuni
1. Pada sidebar, klik 'Seleksi Penghuni' untuk melihat calon-calon penghuni yang telah mendaftarkan diri ke Smart Dormitory.
2. Terdapat beberapa status yang dimiliki calon penghuni. Secara berurutan, statusnya yaitu:
   * Belum Direview: Data calon penghuni baru masuk dan belum direview admin.
   * Menunggu Pembuatan Kontrak: Data calon penghuni telah direview dan menunggu dibuatkan kontrak.
   * Menunggu Pembayaran: Calon penghuni perlu membayar.
3. Klik *icon* pensil pada salah satu calon penghuni dengan status **Belum Direview**.

   ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/0911db06-d735-47ef-9d72-a3127088c7b5)

3. Pastikan data penghuni telah betul semua. Klik **Tolak** untuk menolak calon penghuni atau **Lanjut ke Pembuatan Kontrak** untuk pembuatan kontrak. 

   ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/f7ac0a05-74a4-4a2f-ab62-4945dc703b00)

5. Status penghuni akan berubah menjadi **Menunggu Pembuatan Kontrak**. Klik *icon* pensilnya lagi untuk melanjutkan proses seleksi.

   ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/db889803-b6b0-4619-8d82-51dcf1b8ff26)

6. Alokasikan kamar pada calon penghuni dengan *dropdown* dan masukkan pin akses pada kolom-kolom tersebut. Setelah itu, klik **Kirim Kontak**.

   ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/52e1b18a-6693-4035-8f17-74f4da8fed10)

   Kontrak yang berhasil dibuat akan memicu munculnya *toast* seperti berikut.

   ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/b7f7733d-7ee9-4dc7-9329-28a70e43a5f2)

7. Status calon penghuni akan berubah menjadi **Menunggu Pembayaran**. Klik *icon* pensil kembali untuk melanjutkan proses seleksi.

   ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/fb20c421-fe77-4ce9-a81e-8485bc628e96)

8. Pembayaran dikonfirmasi melalui rekening Smart Dormitory. Jika sudah terkonfirmasi, maka Pembayaran dapat di-*checklist*.

   ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/29480bf1-489a-4a05-84b1-8d73e925d7ab)

9. Setelah di-*checklist*, data calon penghuni akan masuk menjadi penghuni dan dapat dilihat pada laman **Penghuni**.

    ![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/d1284be9-2789-4c75-a735-6c0fc7c6c302)


### Penghuni

Laman penghuni berisikan data-data penghuni yang telah diterima dan dialokasikan kamar.

![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/6e03d2e0-d46e-46cd-9e5b-8b3750a2263e)

Penghuni juga dapat dihapus apabila sudah bukan menjadi anggota dormitory lagi.

![image](https://github.com/abrahammegantoro/go-smart-dormitory/assets/52821168/e19c4fe1-ef95-4f00-b5d7-1c50e8f82515)





