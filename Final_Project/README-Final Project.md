# API MyGram
=========================================================================

##Biodata 
Kode Peserta : 149368582100-427
Nama		 : Elang Eka Marga Putra
Link Github  : https://github.com/Elang-Eka/427_Elang_Eka_Marga_Putra_gob2/tree/master/Final_Project
Alamat	 : Jl Singomenggolo, Dsn. malang ganting, Ds. Ganting, Gedangan, Sidoarjo

##Paduan Running Aplikasi	: 
1. Jalankan Xampp Apcahe dan MySQL
2. Buka VSCode di folder MyGram
3. Buka terminal dan jalankan go run main.go
4. Klik allow_access firewall ketika muncul popup
5. Buka Postman untuk menjalankan API aplikasi sesuai router yang tersedia

##Router :
=> Auth user Register and Login
	1. http://localhost:8080/users/register 	(POST) 	(formKey: username,password(value 6 character),age,email)
	2. http://localhost:8080/users/login 	(POST)	(formKey: email,password)

=> User Router
	1. http://localhost:8080/users 	(GET)
	2. http://localhost:8080/users 	(PUT) 	(formKey: username,email,age)
	3. http://localhost:8080/users 	(DELETE) 
	4. http://localhost:8080/users/all 	(GET)

=> Social Media Router
	1. http://localhost:8080/socialmedias 			(GET)
	2. http://localhost:8080/socialmedias 			(POST) 	(formKey: name,social_media_url)
	3. http://localhost:8080/socialmedias/:socialMediaId 	(PUT) 	(formKey: name,social_media_url)
	4. http://localhost:8080/socialmedias/:socialMediaId 	(DELETE)
	5. http://localhost:8080/socialmedias/:userId		(GET)

=> Photo Router
	1. http://localhost:8080/photos 		(POST) 	(formKey: photo_url,title,caption)
	2. http://localhost:8080/photos 		(GET)
	3. http://localhost:8080/photos/:photoId 	(PUT) 	(formKey: photo_url,title,caption)
	4. http://localhost:8080/photos/:photoId 	(DELETE)

=> Comment Router
	1. http://localhost:8080/comments 			(POST) 	(formKey:message,photo_id)
	2. http://localhost:8080/comments/:photoId 	(GET)
	3. http://localhost:8080/comments/:commentId 	(PUT) 	(formKey:message)
	4. http://localhost:8080/comments/:commentId 	(DELETE)
