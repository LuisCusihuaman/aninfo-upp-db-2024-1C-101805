# PSA Database Seeder and Setup

The PSA database, seeding it with fake data using Gorm and Faker in Go, and running it in a Docker container.

## 🌟 Features

- **Database Setup**: Instructions to set up a MySQL database using Docker.
- **Database Seeding**: Seed the database with initial and related data using Go, Gorm, and Faker.
  ![docs.png](docs.png)

## 🛠️ Installation

To set up the PSA database on your local environment, follow these steps:

1. **Run MySQL Database in Docker**:

   ```bash
   docker run --name=db-psa \
   -p 3306:3306 \
   -v mysql-psa:/var/lib/mysql \
   -e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
   -e MYSQL_DATABASE=psa \
   -d mysql:8.0.36 \
   --entrypoint "/usr/local/bin/docker-entrypoint.sh" \
   mysql:8.0.36 psa.db
   ```

2. **Clone the Repository**:

   ```bash
   git clone https://github.com/tribu-a-2024-1c/psa-database-seeder
   cd psa-database-seeder
   ```

3. **Install Dependencies**:

   Ensure you have Go installed. Then, install the necessary Go packages:

   ```bash
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   go get -u github.com/bxcodec/faker/v3
   ```

4. **Database Seeding**:

   The Go file `main.go` is included in the repository with the following content to seed the database:

   ```go
   package main

   func main() {
       dsn := "root:@tcp(127.0.0.1:3306)/psa?charset=utf8mb4&parseTime=True&loc=Local"
       db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
           NamingStrategy: schema.NamingStrategy{
               SingularTable: true,
           },
       })
       fake := faker.New()
       rand.New(rand.NewSource(time.Now().UnixNano()))
	   ...
       seedInitialTables(db, fake, primaryKeys, usedIntIDs)
       seedRelatedTables(db, fake, primaryKeys, usedIntIDs)

   ```

5. **Run the Seeder**:

   Execute the Go file to seed the database:

   ```bash
   go run main.go
   ```

## 🤝 Contributors

Contributions are welcome! Feel free to submit a pull request or open an issue for any improvements or bug fixes.

