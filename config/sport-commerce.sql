CREATE TYPE valid_roles AS ENUM ('Customer', 'Admin');
-- Users Table
CREATE TABLE Users (
    UserID SERIAL PRIMARY KEY,
    Username VARCHAR(100) NOT NULL, 
    Password VARCHAR(50) NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    Address VARCHAR(255),
    Balance DECIMAL(10, 2),
    Role valid_roles NOT NULL DEFAULT 'Customer' -- new
);

-- Categories Table
CREATE TABLE Categories (
    CategoryID SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL
);

-- Products Table
CREATE TABLE Products (
    ProductID SERIAL PRIMARY KEY,
    CategoryID INT NOT NULL REFERENCES Categories(CategoryID),
    Name VARCHAR(100) NOT NULL,
    Description TEXT,
    Price DECIMAL(10, 2) NOT NULL,
    QuantityInStock INT NOT NULL
);

-- Carts Table
CREATE TABLE Carts (
    CartID SERIAL PRIMARY KEY,
    UserID INT NOT NULL REFERENCES Users(UserID),
    ProductID INT NOT NULL REFERENCES Products(ProductID),
    Quantity INT NOT NULL,
    SubTotal DECIMAL(10, 2) NOT NULL
);

CREATE TYPE valid_order_status AS ENUM ('Pending', 'Completed', 'Cancelled');
CREATE TYPE valid_payment_method AS ENUM ('Cash', 'Member Balance', 'Debit Card', 'Credit Card'); -- new
-- Orders Table
CREATE TABLE Orders (
    OrderID SERIAL PRIMARY KEY,
    UserID INT NOT NULL REFERENCES Users(UserID),
    OrderDate DATE NOT NULL,
    ProductID INT NOT NULL REFERENCES Products(ProductID), -- new
    Quantity INT NOT NULL, -- new
    SubTotal DECIMAL(10, 2) NOT NULL, -- new
    Status valid_order_status NOT NULL DEFAULT 'Pending',
    PaymentMethod valid_payment_method NOT NULL
);