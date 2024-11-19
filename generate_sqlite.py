import sqlite3
import csv
import os
import sys

# Configuration
DB_NAME = "my_database.db"
TSV_FILE = "data/modelSeed.tsv"
TABLE_NAME = "data"

# Check if the TSV file exists
if not os.path.exists(TSV_FILE):
    print(f"Error: File '{TSV_FILE}' not found!")
    sys.exit(1)

# Infer column names and data types
with open(TSV_FILE, "r", encoding="utf-8") as file:
    reader = csv.reader(file, delimiter="\t")
    
    # Extract header and first row
    header = next(reader)
    first_row = next(reader, None)

    # Infer data types
    def infer_type(value):
        if value.isdigit():
            return "INTEGER"
        try:
            float(value)
            return "REAL"
        except ValueError:
            return "TEXT"

    data_types = [infer_type(value) for value in first_row] if first_row else ["TEXT"] * len(header)

# Construct CREATE TABLE statement
create_table_sql = f"CREATE TABLE {TABLE_NAME} ("
create_table_sql += ", ".join(f"{col} {dtype}" for col, dtype in zip(header, data_types))
create_table_sql += ");"

# Load data into SQLite
try:
    conn = sqlite3.connect(DB_NAME)
    cursor = conn.cursor()
    
    # Drop table if it exists
    cursor.execute(f"DROP TABLE IF EXISTS {TABLE_NAME}")
    cursor.execute(create_table_sql)

    # Import TSV data
    with open(TSV_FILE, "r", encoding="utf-8") as file:
        reader = csv.DictReader(file, delimiter="\t")
        columns = reader.fieldnames
        insert_sql = f"INSERT INTO {TABLE_NAME} ({', '.join(columns)}) VALUES ({', '.join(['?'] * len(columns))})"
        cursor.executemany(insert_sql, (list(row.values()) for row in reader))

    conn.commit()
    print(f"Data imported successfully into '{TABLE_NAME}' in database '{DB_NAME}'.")
except sqlite3.Error as e:
    print(f"SQLite error: {e}")
finally:
    if conn:
        conn.close()

