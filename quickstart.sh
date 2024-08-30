#!/bin/bash

# Initialize backend
echo "Initializing backend..."
cd backend
make up

# Check if the backend started successfully
if [ $? -ne 0 ]; then
    echo "Backend initialization failed. Exiting..."
    exit 1
fi

# Initialize frontend
echo "Initializing frontend..."
cd ../frontend
npm install

if [ $? -ne 0 ]; then
    echo "npm install failed. Exiting..."
    exit 1
fi

npm run build

if [ $? -ne 0 ]; then
    echo "npm run build failed. Exiting..."
    exit 1
fi

npm run preview

if [ $? -ne 0 ]; then
    echo "npm run preview failed. Exiting..."
    exit 1
fi

echo "Quickstart complete!"