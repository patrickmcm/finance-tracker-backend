import csv
import psycopg

with psycopg.connect("dbname=finance_tracker user=test password=test") as conn:
    with open('instruments_lse.csv', newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        print(reader.fieldnames)
        for row in reader:
            if row['Trading Currency'] == 'GBX' or row['Trading Currency'] == 'GBP':
                statement = f"INSERT INTO instruments VALUES (\'{row['\ufeffTIDM']}\', \'{row['Instrument Name'].replace('\'', '')}\', \'{row['Trading Currency']}\', \'{row['ISIN']}\', 'ETF')"
                conn.execute(statement)