import csv
import math

import pandas
import psycopg
import yfinance as yf
from google.protobuf.internal.well_known_types import Timestamp

with psycopg.connect("dbname=finance_tracker user=test password=test") as conn:
    with open('instruments_lse.csv', newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        tickers = []
        for row in reader:
            if row['Trading Currency'] == 'GBX' or row['Trading Currency'] == 'GBP':
                tickers.append(row['\ufeffTIDM']+".L")

        testTickers = ['VUAG.L', 'CSH2.L']

        data = yf.download(tickers, period='5d', group_by='ticker', keepna=False)

        formattedData = dict()
        for tickerName in tickers:
            formattedData[tickerName] = dict()
            for priceType, priceData in data[tickerName].items():
                for date, price in priceData.items():
                    if not (date in formattedData[tickerName]):
                        formattedData[tickerName][date] = dict()
                    formattedData[tickerName][date][priceType.lower()] = price


        for symbol, priceData in formattedData.items():
            for date, price in priceData.items():
                if math.isnan(price['open']):
                    continue
                conn.execute("INSERT INTO instrument_prices VALUES (%s, %s, %s, %s, %s, %s, %s)", (
                    symbol.removesuffix('.L'),
                    date,
                    price['open'],
                    price['close'],
                    price['high'],
                    price['low'],
                    price['volume'],
                ))
