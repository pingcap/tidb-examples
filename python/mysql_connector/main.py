import mysql.connector

# Connect to the database
mydb = mysql.connector.connect(
  host="localhost",
  port="4000",
  user="root",
  passwd="",
  database="tidb_example"
)

# create a database cursor
mycursor = mydb.cursor()

# create a tables
mycursor.execute("CREATE TABLE IF NOT EXISTS orders (oid INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT, cid INT UNSIGNED, price FLOAT);")
mycursor.execute("CREATE TABLE IF NOT EXISTS customer (cid INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), gender ENUM ('Male', 'Female') NOT NULL)")


		#"INSERT INTO customer (name, gender) value (%s, %s);"
		#"INSERT INTO orders (cid, price) value (1,10.23);",

add_customer = ("INSERT INTO customer (name, gender) VALUES (%(name)s, %(gender)s);")
add_order = "INSERT INTO orders (cid, price) VALUES ({}, {});"



data_customers = [
    {'name': 'Ben', 'gender': 'Male'},
    {'name': 'Alice', 'gender': 'Female'},
    {'name': 'Peter', 'gender': 'Male'},
]
data_orders = [
    [1.3, 4.0, 52.0, 123.0, 45.0],
    [2.4, 23.4],
    [100.0],
]

# Insert new employee
for data_customer in data_customers:
    mycursor.execute(add_customer, data_customer)
    mydb.commit()

cid = 1
for price in data_orders[cid-1]:
    #mycursor.execute(add_order, (cid, price))
    mycursor.execute(add_order.format(cid, price))
    cid = cid + 1
    mydb.commit()

# query customer
mycursor.execute("SELECT * FROM customer")
myresult = mycursor.fetchall()
for x in myresult:
  print(x)

# update orders
mycursor.execute("UPDATE orders SET price = %s WHERE oid = %s", (100.0, 1))
mydb.commit()

# join_result
mycursor.execute("SELECT customer.name, orders.price FROM customer INNER JOIN orders ON customer.cid = orders.cid")
myresult = mycursor.fetchall()
for x in myresult:
  print(x)

# close the database connection
mycursor.close()
mydb.close()
