package main

import (
	"errors"
	"fmt"
	"hello/database"
	"hello/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	Pilihan()
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data")
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v\n\n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}

	fmt.Printf("Update user's email: %+v \n\n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err)
		return
	}

	fmt.Println("New Product Data:", Product)
}

func getUsersWithProduct() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with product", err)
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v\n\n", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error Deleting Product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted\n\n", id)
}

func Pilihan() {
	var pilih int
	var loop = 1
	for i := 0; i <= loop; i++ {
		if loop == 1 {
			fmt.Println("==============================")
			fmt.Println("1. Create User")
			fmt.Println("2. Get User By Id")
			fmt.Println("3. Update User By Id")
			fmt.Println("4. Create Product")
			fmt.Println("5. Get Users With Product")
			fmt.Println("6. Delete Product By Id")
			fmt.Println("7. Exit")
			fmt.Println("==============================")
			fmt.Print("Masukkan Pilihan : ")
			fmt.Scanln(&pilih)
			if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 || pilih == 5 || pilih == 6 {
				switch pilih {
				case 1:
					var email string
					fmt.Print("Masukkan Email user yang akan dibuat: ")
					fmt.Scanln(&email)
					createUser(email)
					i = 0
					loop = 1
				case 2:
					var id uint
					fmt.Print("Masukkan Id user yang akan di lihat: ")
					fmt.Scanln(&id)
					getUserById(id)
					i = 0
					loop = 1
				case 3:
					var email string
					var id uint
					fmt.Print("Masukkan Id user yang akan di ubah: ")
					fmt.Scanln(&id)
					fmt.Print("Masukkan email user yang baru: ")
					fmt.Scanln(&email)
					updateUserById(id, email)
					i = 0
					loop = 1
				case 4:
					var id uint
					var brand string
					var nama string
					fmt.Print("Masukkan Id user: ")
					fmt.Scanln(&id)
					fmt.Print("Masukkan brand produk")
					fmt.Scanln(&brand)
					fmt.Print("Masukkan nama produk")
					fmt.Scanln(&nama)
					createProduct(id, brand, nama)
					i = 0
					loop = 1
				case 5:
					getUsersWithProduct()
					i = 0
					loop = 1
				case 6:
					var id uint
					fmt.Print("Masukkan Id user: ")
					fmt.Scanln(&id)
					deleteProductById(id)
					i = 0
					loop = 1
				case 7:
					i = 1
					loop = 2
				}
			} else {
				i = 0
				loop = 1
				fmt.Println("Pilihan tidak ada menu hanya 1 hingga 7")
			}
		} else {
			fmt.Println("Terima Kasih")
		}
	}
}
