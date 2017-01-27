gambas
======

A simple Mock Server mimicking the GBIF API, for testing purposes.

Usage
-----

``$ gambas``

=> available at http://127.0.0.1:8080

**Listen to different IP/port:**
    
``$ gambas -listen=0.0.0.0:4567``

=> available on all network interfaces, port 4567

**More info:**

``$ gambas --help``

Development
-----------

- gambas is written in golang
- run tests with ``$ go test``