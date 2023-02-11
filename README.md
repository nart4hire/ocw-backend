# ITBOpenCourseWare Backend

## Development Guide

1. Checkout ke branch staging + pull

   ```sh
   git checkout staging
   git pull
   ```

2. Buat branch baru dari staging dengan format `feat/s<nomor sprint>-sb<nomor sb>-<nama fitur dipisah dengan strip (-)>`. Jika membuat perubahan dari fitur yang sudah ada di staging, formatnya menjadi `fix/s<nomor sprint>-sb<nomor sb>-<apa yang diubah dari fitur tersebut>`

   ```sh
   git checkout -b feat/s1-sb1-login
   ```

3. Lakukan perubahan pada branch tersebut, commit dengan format `feat(<scope fitur>): <isi perubahan>`. Jika melakukan perubahan pada fitur yang sudah ada, formatnya menjadi `fix(<scope fitur>): <isi perubahan>`, lain-lainnya bisa dilihat di [semantic commit](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716)

   ```sh
   git commit -m "feat(login): add login page"
   ```

4. Push branch ke remote

   ```sh
   git push origin feat/s1-sb1-login
   ```

5. Buat merge request ke branch staging, tambah assignee diri sendiri dan reviewer

<!-- TODO: how to handle conflict -->