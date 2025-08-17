# made User struct in -------------->users/user.go
# database connection ----------------> cmd/db/connect.go
# called ConnectDB in -------------------> cmd/main.go
# made initTable to make table if not exist in -------->cmd/db/initDB.go
# called initTables() to ------------------>cmd/main.go
# 
# made JSONContentTypeMiddleware middleware in ------------> pkg/middleware/json.go
# creating controller for users in ------------------> internal/users/controller/go
# created createUser and also added bcrypt in --------------->controller.go
# learnt how bcrypt works
# added UserResponse struct below User struct in -----------------------> user.go
#
# separted insert Query of database from controller to --------------->repository.go
# similarly Handles password hashing, validation, and calling repository ---------------->service.go
# 

1. Model / Entity (internal/users/user.go)
2. Repository (DB queries) (internal/users/repository.go)
3. Service (business logic) (internal/users/service.go)
4. Controller (HTTP handlers) (internal/users/controller.go)
5. Routes (internal/users/routes.go)
# 
# made user in psql with code sudo -u postgres psql and then same id pass to tableplus
#
#
#
# router(routes.go)----->GetUsersController(controller.go)----->GetAllUsersService(service.go)----->FetchUsers(repository.go)
#
# CreateUsersController DONE 
# GetUserController DONE
#
#
# 
#
#
#
#
#
#
#
#