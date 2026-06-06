#!/bin/bash

# and-agy proot setup script

set -e

echo "=== and-agy proot Setup ==="
echo ""

# Check if running in Termux
echo "Checking Termux environment..."
if [ ! -d "$PREFIX" ]; then
    echo "Error: Not running in Termux"
    exit 1
fi

# Install proot
echo "Installing proot..."
pkg install -y proot

# Install proot-distro
echo "Installing proot-distro..."
pkg install -y proot-distro

# Install Ubuntu distribution
echo "Installing Ubuntu distribution..."
proot-distro install ubuntu

# Login to Ubuntu and install dependencies
echo "Setting up Ubuntu environment..."
proot-distro login ubuntu -- bash -c "
apt update && apt upgrade -y
apt install -y git golang make
"

echo ""
echo "=== proot Setup Complete ==="
echo "To use proot environment:"
echo "  proot-distro login ubuntu"
echo ""
echo "Then clone and build and-agy:"
echo "  git clone https://github.com/Chieji/and-agy.git"
echo "  cd and-agy"
echo "  go build -o agy ."
echo "  ./agy"