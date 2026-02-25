# Go2048
Greetings folks, go2048 is a 2048 game using Raylib written in Golang.

# Flatpak image
There is not any official flatpak image to download from flathub, but you can built-it by yourself using 'org.flatpak.2048.yml' file.

# Build
It is as easy as using this command:

```
make build
```

> Please use 'make' utile for building the project because there is more process than just compiling the binary inside the 'Makefile'.

# Why
Because it's fun. That's the main point of recreational programming.

# Wayland
You can easly compile the project for wayland using 'BUILD_MODE' VAR in 'Makefile'.
But for some reason, Wayland support doesn't work properly in Hyprland, though it works in Sway.
For fixing the problem in hypr, easly just build the project with BUILD_MODE=x11;

# More
If you see any problems with the project, I would be happy to read your PL. This project is not my main focus then I will barely look check out the source code, your help can keep this game alive :3.


