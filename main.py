#!/usr/bin/env python

import sqlite3

from sqlite3 import Error
from typing import Union
from venv import create
from fastapi import FastAPI

app = FastAPI()

"""
    Root route !

    Return only a dict with ping/pong
"""
@app.get('/')
def index() -> dict:
    return {'ping': 'pong'}

"""
    Search for a station
"""
@app.get('/station')
def search_station(q: Union[str, None] = None, limit: int = 10) -> dict:
    result = _search_station(q, limit)
    return {'query': q, 'limit': limit, 'result_count': len(r), 'result': result}

"""
    Search for all stations on the same line as the one provided
"""
@app.get('/station/{station_id}')
def info_stations(station_id: str) -> dict:
    line, result = _search_line(station_id)
    return {'station_id': station_id, 'line': line, 'result': result}

"""
    Search for the shortest way to go from Station 1 to Station 2
"""
@app.get('/shortest/{from_station}/{to_station}')
def shortest_way(from_station: str, to_station: str) -> dict:
    return {'from': from_station, 'to': to_station}

"""
    Search for the fastest way to go from Station 1 to Station 2
"""
@app.get('/fastest/{from_station}/{to_station}')
def fastest_way(from_station: str, to_station: str) -> dict:
    return {'from': from_station, 'to': to_station}

##########################################################

database_file = 'database.db'

def create_connection():
    conn = None
    try:
        conn = sqlite3.connect(database_file)
    except Error as e:
        print(e)

    return conn

def close_connection(conn):
    if conn:
        conn.close()

def _search_station(q, limit = 10):
    result = []

    conn = create_connection()
    cur = conn.cursor()
    cur.execute(f'SELECT Name, Reference FROM stations WHERE Name LIKE \'%{q}%\' LIMIT ?', (limit,))

    rows = cur.fetchall()

    for row in rows:
        result.append({
            'Name': row[0],
            'Ref': row[1]
        })
    return result

'''
    NOT WORKING AS INTENDED .... TOO MUCH STATIONS
    AND THE ORDER IS WRONG
'''
def _search_line(station_id):
    result = []

    conn = create_connection()
    cur = conn.cursor()
    
    cur.execute("SELECT LineName FROM stations WHERE Reference = ?", (station_id,))
    rows = cur.fetchall()
    for row in rows:
        line = row

    line_name = line[0]

    cur.execute("SELECT Name, Reference FROM stations WHERE LineName = ? GROUP BY Name ORDER BY id", (line_name,))
    rows = cur.fetchall()
    i = 0
    for row in rows:
        result.append({
            'Order': i,
            'Name': row[0],
            'Ref': row[1]
        })
        i += 1

    return line_name, result