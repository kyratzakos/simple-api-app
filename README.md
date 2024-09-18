# GoLang MongoDB CRUD API
A RESTful API built with Go that performs standard CRUD operations on a products resource, backed by a MongoDB database. This API includes the ability to create, read, update, and delete products, with built-in validation for required fields and timestamp management.

### Features
- Create, read, update, and delete operations on `products`.
- Uses MongoDB for persistent storage.
- Dynamic field handling: Only non-empty fields are inserted or updated in MongoDB.
- Meta timestamps: Automatically handles `createdAt` and `updatedAt` fields.
- Field validation: Only allowed fields are accepted, with checks for required fields like `name` and `price`.

##
### Table of Contents
- [Installaton](#installation)
- [API Endpoints](#api-endpoints)
- [Running the Project](#running-the-project)
- [Future Improvements](#future-improvements)

### Installation
To run this project locally, ensure that you have the following prerequisites installed:

### Prerequisites
- [Go](https://go.dev/doc/install) (v1.18 or later)
- [MongoDB](https://www.mongodb.com/docs/manual/installation/)
- [Git](https://git-scm.com/)

### Steps
1. Clone the Repository:

```bash []
git clone https://github.com/kyratzakos/simple-api-app.git
cd simple-api-app
```
2. Install Dependencies:

   Run the following command in the project root directory to install the required Go dependencies:

```bash
go mod tidy
```
3. Set Up MongoDB:

   Ensure MongoDB is running on your machine. You can start MongoDB using:

- Linux/Mac:
```bash
sudo service mongod start
```
- Windows:
```bash
net start MongoDB
```
4. Set Environment Variables:

Create a copy of [.env.example](./.env.example) file in the root directory with name `.env`:

##
### API Endpoints
1. Create Product
   - Endpoint: POST /products
   - Description: Creates a new product.
   - Body Parameters:
```json
{
  "title": "Product A",
  "description": "A high-quality product",
  "price": 199.99,
  "category": "Optional",
  "meta": {
    "createdAt": "optional_timestamp",
    "updatedAt": "optional_timestamp"
  }
}
```
   - Required Fields: `title`, `price`
   - Response: 201 Created, with the inserted product `id`.
2. Get Products (Paginated)
   - Endpoint: `GET /products?page=1&limit=10`
   - Description: Retrieves all products with pagination.
   - Query Parameters: page, limit
   - Response: 200 OK with paginated product list.
3. Get Single Product
   - Endpoint: `GET /products/{id}`
   - Description: Retrieves a single product by its MongoDB ObjectId.
   - Response: 200 OK with the product data.
4. Update Product
   - Endpoint: PUT /products/{id}
   - Description: Updates an existing product's fields dynamically (only non-empty fields are updated).
   - Body Parameters:
```json
{
  "name": "Updated Product Name",
  "price": 299.99,
  "meta": {
    "updatedAt": "2024-09-17T12:34:56Z"
  }
}
```
   - Response: 200 OK with a success message.
5. Delete Product
   - Endpoint: `DELETE /products/{id}`
   - Description: Deletes a product by its MongoDB ObjectId.
   - Response: 200 OK with a success message.

##
### Running the Project
Once you have set up the environment variables and MongoDB is running, you can start the server with:

```bash
go run main.go
```
The API will be available at http://localhost:3000.

##
### Future Improvements
- Add unit tests for the CRUD functionality.
- Implement authentication and authorization.
- Add more comprehensive validation for request payloads.
