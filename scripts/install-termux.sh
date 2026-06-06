#!/bin/bash

# and-agy Termux installation script

set -e

echo "=== and-agy Termux Installation ==="
echo ""

# Check if running in Termux
echo "Checking Termux environment..."
if [ ! -d "$PREFIX" ]; then
    echo "Error: Not running in Termux"
    exit 1
fi

# Update packages
echo "Updating packages..."
pkg update && pkg upgrade -y

# Install dependencies
echo "Installing dependencies..."
pkg install -y git golang make

# Clone repository
echo "Cloning repository..."
cd $HOME
if [ -d "and-agy" ]; then
    cd and-agy
    git pull
else
    git clone https://github.com/Chieji/and-agy.git
    cd and-agy
fi

# Build
echo "Building and-agy..."
go build -o agy -ldflags "-s -w" .

# Install
echo "Installing..."
cp agy $PREFIX/bin/agy
chmod +x $PREFIX/bin/agy

# Cleanup
echo "Cleaning up..."
cd $HOME
rm -rf and-agy

echo ""
echo "=== Installation Complete ==="
echo "Run 'agy' to start the application"
echo ""
echo "To set up authentication:"
echo "  export GEMINI_API_KEY=your-api-key"
echo "  agy auth login gemini"