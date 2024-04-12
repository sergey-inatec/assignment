#!/bin/bash

# Constants
UBUNTU_ISO_URL="https://releases.ubuntu.com/20.04/ubuntu-20.04.6-desktop-amd64.iso"
DISK_IMAGE="disk.img"
MOUNT_DIR="mnt"
HELLO_WORLD_FILE="hello_world.txt"

# Download Ubuntu ISO
wget -O ubuntu.iso "$UBUNTU_ISO_URL"

# Create disk image
qemu-img create -f qcow2 "$DISK_IMAGE" 4G

# Install Ubuntu
echo "Installing Ubuntu on the disk image..."
sudo qemu-system-x86_64 -hda "$DISK_IMAGE" -cdrom ubuntu.iso -boot d -m 2G -enable-kvm /usr/bin/kvm

# Run QEMU
echo "Booting the disk image with QEMU..."
sudo qemu-system-x86_64 -hda "$DISK_IMAGE" -m 2G -enable-kvm /usr/bin/kvm &

# Wait for QEMU to boot
sleep 30

# Mount the disk image
echo "Mounting the disk image..."
mkdir -p "$MOUNT_DIR"
sudo mount -o romount,rw loop "$DISK_IMAGE" "$MOUNT_DIR"

# Print "hello world" inside the mounted image
echo "hello world" | sudo tee "$MOUNT_DIR/$HELLO_WORLD_FILE"

# Unmount the disk image
echo "Unmounting the disk image..."
sudo umount "$MOUNT_DIR"

# Clean up
rm -rf "$MOUNT_DIR"
rm ubuntu.iso

