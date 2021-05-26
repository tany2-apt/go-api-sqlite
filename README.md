# go-api-sqlite
This is an application written in go which uses GORM to interact with sqliteDB.
The server runs on port:8080.
The Table has the following attributes:
Id      
	FirstName
	LastName  
	City    
	Phone   
	Height   
	Gender   
	Password  
	Married  
	Created  
	Updated  

# 🐱‍🏍The api performs following operations at the given endpoints
✨
Fetches by list of IDs
pi/v1/user/gets

✨
Fetches record by a given ID
/api/v1/user/{id}

✨
Fetches records of all users present in DB
/api/v1/user/fetch

✨
Insert a new record into the BB
/api/v1/user/create

# For encrypting the password of an user: 
SHA-1 has been used

# Date and time 
Epoch convention is used for storing the data and time of 'Creation and Updation' of a record. 
