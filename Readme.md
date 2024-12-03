# create project 
1. mkdir go-send-mail
2. init project go `go mod init gosendmail`
3. install dependency 
    `$ go get github.com/jackc/pgx/v5`
    `$ go get -u gorm.io/gorm`
    `$ go get -u github.com/gofiber/fiber/v2` 
    `$ go get -u github.com/gofiber/fiber/v2/middleware/logger`
    `$ go get github.com/go-gomail/gomail`
    `$ go get github.com/joho/godotenv`
`