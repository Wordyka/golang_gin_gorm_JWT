# golang_gin_gorm_JWT

JSON Web Token (JWT) adalah sebuah random token berbentuk string panjang yang digunakan untuk melakukan sistem otentikasi dan pertukaran informasi. Token yang dihasilkan akan disimpan oleh user pada cookies browser atau local storage, ketika user ingin mengakses halaman tertentu maka harus menyertakan token tersebut. Untuk struktur JWT terdiri dari tiga bagian yaitu header, payload dan signature.

1. Header, berisi algoritma dan jenis token yang digunakan
2. Playload, berisi data atau informasi yang ingin dikirimkan ke server JWT misalnya berupa id, nama user dan lain lain,
3. Signature. yaitu verify signature berupa hasil gabungan dari header dan payload yang telah di encode serta penambahan kode secretnya

Adapun tahapan yang ada penggunaan JWT adalah sebagai berikut (Rani and Sangeetha 2021).
• Client pertama kali akan melakukan login dengan mengajukannya kepada server
• Kemudian diberikan suatu key yaang bersifat private key yang digunakan sebagai signature dari token dan JWT dikirimkan kepada client
• Token kemudian disimpan pada cookies browser / local storage agar setiap kali client melakukan login, dia dapat menambahkan permintaan dengan token tersebut.
• Ketika client melakukan suatu request HTTP apapun, maka di perlu melampirkan token dan mengirimkannya ke server untuk diotorisasi dalam membuat respond atas request client tersebut.

Berikut adalah diagram yang menjelaskan proses client mengakses service menggunakan Autentikasi JWT.

![Sequence Diagram1](https://user-images.githubusercontent.com/77338737/163722232-f91ef9da-1769-4061-bf0c-4f3ed3c998e9.png)
