# Run your android QEMU with on :5555

```bash
qemu-system-x86_64 -hda android.qcow2 -boot d -m 2048 -enable-kvm -net user,hostfwd=tcp::5555-:5555 -net nic

```

# android-tools

```bash
sudo pacman -S android-tools

# connect to the android emulator
adb connect 127.0.0.1:5555

# see all connected devices
adb devices

# see logs used for debugging
adb logcat

# disconnect to a device
adb disconnect localhost:5900

# use for restarting
adb kill-server
adb start-server


adb shell
su
cd /data/data/com.example.myapp/databases/

```
