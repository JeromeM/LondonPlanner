# LONDON PLANNER

London Planner is an API to search London Underground stations.  
You can search for station name, information, and plan your shortest / faster journey from station A to station B.

## Installation

Run the following command for development run :  
`make run_dev`

Run the following command for production run :  
`make run`
  
## Routes  

_Search a station_ :
> /station?q={SEARCH}&limit={LIMIT}

_Search for all stations on the same line as the one provided_ :
> /station/{STATION_ID}

_Search for the shortest way to go from Station 1 to Station 2_ :
> /shortest/{FROM_STATION}/{TO_STATION}

_Search for the fastest way to go from Station 1 to Station 2_ :
> /fastest/{FROM_STATION}/{TO_STATION}