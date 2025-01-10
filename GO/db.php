<?php
// Database configuration
$host = "localhost"; // Replace with your host
$username = "root";  // Replace with your MySQL username
$password = "";      // Replace with your MySQL password
$dbname = "your_database_name"; // Replace with your database name

// Create a connection
$conn = new mysqli($host, $username, $password, $dbname);

// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}

// SQL query to create the Users table
$sql = "CREATE TABLE IF NOT EXISTS Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE
)";

// Execute the query
if ($conn->query($sql) === TRUE) {
    echo "Table Users created successfully.";
} else {
    echo "Error creating table: " . $conn->error;
}

// Insert sample data into the Users table
$sqlInsert = "INSERT INTO Users (name, email) VALUES ('John Doe', 'johndoe@example.com')";
if ($conn->query($sqlInsert) === TRUE) {
    echo "Sample data inserted successfully.";
} else {
    echo "Error inserting data: " . $conn->error;
}

// Fetch and display data from the Users table
$sqlFetch = "SELECT * FROM Users";
$result = $conn->query($sqlFetch);

if ($result->num_rows > 0) {
    while ($row = $result->fetch_assoc()) {
        echo "ID: " . $row["id"] . " - Name: " . $row["name"] . " - Email: " . $row["email"] . "<br>";
    }
} else {
    echo "No records found.";
}

// Close the connection
$conn->close();
?>
