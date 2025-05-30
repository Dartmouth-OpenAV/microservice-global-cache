#!/bin/bash

# Global Cache Relay Microservice API Test Script
# This script tests all endpoints from the Global Cache Relay Postman collection

# Configuration variables - Update these values as needed
MICROSERVICE_URL="localhost:8080"
DEVICE_FQDN="globalcache-device.local"
PORT_NUMBER_1="1"
PORT_NUMBER_2="2"
PORT_NUMBER_3="3"

echo "Starting Global Cache Relay Microservice API Tests..."
echo "Microservice URL: $MICROSERVICE_URL"
echo "Device FQDN: $DEVICE_FQDN"
echo "Test Ports: $PORT_NUMBER_1, $PORT_NUMBER_2, $PORT_NUMBER_3"
echo "=============================================="

# GET State for all ports
echo "Testing GET State for Port 1..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_1"
sleep 1

echo "Testing GET State for Port 2..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_2"
sleep 1

echo "Testing GET State for Port 3..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_3"
sleep 1

echo "=============================================="
echo "Starting SET State operations..."
echo "=============================================="

# SET State to "on" for all ports
echo "Testing SET State to 'on' for Port 1..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_1" \
     -H "Content-Type: application/json" \
     -d "\"on\""
sleep 1

echo "Testing SET State to 'on' for Port 2..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_2" \
     -H "Content-Type: application/json" \
     -d "\"on\""
sleep 1

echo "Testing SET State to 'on' for Port 3..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_3" \
     -H "Content-Type: application/json" \
     -d "\"on\""
sleep 1

# SET State to "off" for all ports
echo "Testing SET State to 'off' for Port 1..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_1" \
     -H "Content-Type: application/json" \
     -d "\"off\""
sleep 1

echo "Testing SET State to 'off' for Port 2..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_2" \
     -H "Content-Type: application/json" \
     -d "\"off\""
sleep 1

echo "Testing SET State to 'off' for Port 3..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_3" \
     -H "Content-Type: application/json" \
     -d "\"off\""
sleep 1

echo "=============================================="
echo "Starting SET Trigger operations..."
echo "=============================================="

# SET Trigger for all ports (1 second duration)
echo "Testing SET Trigger for Port 1 (1 second)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/trigger/$PORT_NUMBER_1" \
     -H "Content-Type: application/json" \
     -d "\"1\""
sleep 1

echo "Testing SET Trigger for Port 2 (2 seconds)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/trigger/$PORT_NUMBER_2" \
     -H "Content-Type: application/json" \
     -d "\"2\""
sleep 1

echo "Testing SET Trigger for Port 3 (3 seconds)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/trigger/$PORT_NUMBER_3" \
     -H "Content-Type: application/json" \
     -d "\"3\""
sleep 1

echo "=============================================="
echo "Starting SET Timed Trigger operations..."
echo "=============================================="

# SET Timed Trigger combinations
echo "Testing SET Timed Trigger (Port 1 -> Port 2, 5 seconds)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/timedtrigger/$PORT_NUMBER_1/$PORT_NUMBER_2" \
     -H "Content-Type: application/json" \
     -d "\"5\""
sleep 1

echo "Testing SET Timed Trigger (Port 2 -> Port 3, 10 seconds)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/timedtrigger/$PORT_NUMBER_2/$PORT_NUMBER_3" \
     -H "Content-Type: application/json" \
     -d "\"10\""
sleep 1

echo "Testing SET Timed Trigger (Port 1 -> Port 3, 15 seconds)..."
curl -X PUT "http://$MICROSERVICE_URL/$DEVICE_FQDN/timedtrigger/$PORT_NUMBER_1/$PORT_NUMBER_3" \
     -H "Content-Type: application/json" \
     -d "\"15\""
sleep 1

echo "=============================================="
echo "Final state check - Getting current states..."
echo "=============================================="

# Final state check
echo "Final GET State for Port 1..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_1"
sleep 1

echo "Final GET State for Port 2..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_2"
sleep 1

echo "Final GET State for Port 3..."
curl -X GET "http://$MICROSERVICE_URL/$DEVICE_FQDN/state/$PORT_NUMBER_3"
sleep 1

echo "=============================================="
echo "All Global Cache Relay API tests completed!"
echo "=============================================="
