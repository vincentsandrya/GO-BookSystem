# GO-BookSystem


<!-- DOCUMENTATION API (POSTMAN .JSON)-->
List end point : 
1. POST {{base_url}}/api/users/login
 - Kegunaan : API ini berfungsi untuk login dengan akun yang sudah terdaftar, bisa mengunakan (admin, password). API ini akan mengembalikan data user beserta token yang akan di simpan di cache
2. GET {{base_url}}/api/categories/
 - Kegunaan : API ini akan mengembalikan semua data kategori buku yang tersedia
3. GET {{base_url}}/api/categories/:id
 - Kegunaan : API ini akan mengembalikan detail data kategori buku dengan spesifik id kategori yang diberikan
4. POST {{base_url}}/api/categories/
 - Kegunaan :  API ini berfungsi untuk menambahkan kategori buku, dengan validasi nama harus unik
5. DELETE {{base_url}}/api/categories/:id
 - Kegunaan : API ini untuk menghapus data kategori berdasarkan id kategori yang dikirimkan, dengan validasi jika id tersedia
6. GET {{base_url}}/api/categories/:id/books
 Kegunaan : API ini berfungsi untuk mendapatkan semua data buku yang memiliki kategori sesuai dengan id kategori yang diberikan
7. GET {{base_url}}/api/books
 Kegunaan : API ini untuk mendapatkan semua data buku yang tersedia
8. GET {{base_url}}/api/books/:id
 Kegunaan : API ini untuk mendapatkan data buku spesifik dari id buku yang diberikan
9. POST {{base_url}}/api/books/
 Kegunaan : API ini untuk menambahkan data buku
10. DELETE {{base_url}}/api/books/:id
 - Kegunaan : API ini untuk menghapus data buku berdasarkan id buku yang dikirimkan, dengan validasi jika id tersedia

hasil export .json dari postman

{
	"info": {
		"_postman_id": "07e945b1-2c63-48fb-9fd0-1be046ccfe9b",
		"name": "Sanber",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "15995448"
	},
	"item": [
		{
			"name": "Quiz 3",
			"item": [
				{
					"name": "Categories",
					"item": [
						{
							"name": "GetCategories",
							"request": {
								"method": "GET",
								"header": [],
								"url": {
									"raw": "{{base_url}}/api/categories",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"categories"
									]
								}
							},
							"response": []
						},
						{
							"name": "GetCategoriesById",
							"request": {
								"method": "GET",
								"header": [],
								"url": {
									"raw": "{{base_url}}/api/categories/1",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"categories",
										"1"
									]
								}
							},
							"response": []
						},
						{
							"name": "InsertCategories",
							"request": {
								"method": "POST",
								"header": [],
								"body": {
									"mode": "raw",
									"raw": "{\r\n    \"name\" : \"Koran\",\r\n    \"created_by\" : \"admin\",\r\n    \"modified_by\" : \"admin\"\r\n}",
									"options": {
										"raw": {
											"language": "json"
										}
									}
								},
								"url": {
									"raw": "{{base_url}}/api/categories",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"categories"
									]
								}
							},
							"response": []
						},
						{
							"name": "DeleteCategory",
							"request": {
								"method": "DELETE",
								"header": [],
								"url": {
									"raw": "{{base_url}}/api/categories/2",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"categories",
										"2"
									]
								}
							},
							"response": []
						},
						{
							"name": "GetBooksByCategoryId",
							"request": {
								"method": "GET",
								"header": [],
								"url": {
									"raw": "{{base_url}}/api/categories/2/books",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"categories",
										"2",
										"books"
									]
								}
							},
							"response": []
						}
					]
				},
				{
					"name": "Book",
					"item": [
						{
							"name": "GetBooks",
							"request": {
								"method": "GET",
								"header": [],
								"url": {
									"raw": "{{base_url}}/api/books",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"books"
									]
								}
							},
							"response": []
						},
						{
							"name": "GetBookById",
							"request": {
								"method": "GET",
								"header": [],
								"url": {
									"raw": "{{base_url}}/api/books/1",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"books",
										"1"
									]
								}
							},
							"response": []
						},
						{
							"name": "InsertBook",
							"request": {
								"method": "POST",
								"header": [],
								"body": {
									"mode": "raw",
									"raw": "{\r\n    \"title\": \"Si Kancil (Bagian 1)\",\r\n    \"description\": \"Buku ini menceritakan perjalan petualangan hewan kancil yang cerdas\",\r\n    \"image_url\": \"https://image.popmama.com/content-images/post/20231017/cover%20kancil%20dan%20buaya-FXhSeoHMneVdOUBDhZyUbqexq6JrHMQm.jpg\",\r\n    \"release_year\": 2020,\r\n    \"price\": 25000,\r\n    \"total_page\": 30,\r\n    \"category_id\": 5,\r\n    \"created_by\": \"admin\",\r\n    \"modified_by\": \"admin\"\r\n}",
									"options": {
										"raw": {
											"language": "json"
										}
									}
								},
								"url": {
									"raw": "{{base_url}}/api/books",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"books"
									]
								}
							},
							"response": []
						},
						{
							"name": "DeleteBook",
							"request": {
								"method": "DELETE",
								"header": [],
								"url": {
									"raw": "{{base_url}}/api/books/1",
									"host": [
										"{{base_url}}"
									],
									"path": [
										"api",
										"books",
										"1"
									]
								}
							},
							"response": []
						}
					]
				},
				{
					"name": "Login",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Username\" : \"admin\",\r\n    \"Password\" : \"password\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{base_url}}/api/users/login",
							"host": [
								"{{base_url}}"
							],
							"path": [
								"api",
								"users",
								"login"
							]
						}
					},
					"response": []
				}
			]
		}
	]
}