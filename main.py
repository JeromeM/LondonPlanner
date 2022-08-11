#!/usr/bin/env python

import sqlite3

from sqlite3 import Error
from typing import Union
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
    r = _search_station(q, limit)
    return {'query': q, 'limit': limit, 'result_count': len(r), 'result': r}

"""
    Search for all stations on the same line as the one provided
"""
@app.get('/station/{station_id}')
def info_stations(station_id: str) -> dict:
    return {'station_id': station_id}

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