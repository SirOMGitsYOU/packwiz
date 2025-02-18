# packwiz
packwiz is a command line tool for creating Minecraft modpacks. Instead of managing JAR files directly, packwiz creates TOML metadata files which can be easily version-controlled and shared with git (see an example pack [here](https://github.com/packwiz/packwiz-example-pack)). You can then [export it to a CurseForge or Modrinth modpack](https://packwiz.infra.link/tutorials/hosting/curseforge/), or [use packwiz-installer](https://packwiz.infra.link/tutorials/installing/packwiz-installer/) for an auto-updating MultiMC instance.

packwiz is great for...

- Distributing private modpacks for servers
- Creating modpacks for CurseForge and Modrinth

packwiz is not so great for...

- Managing downloaded mod files (use [Curse/GDLauncher or another CLI](https://gist.github.com/comp500/13ae6f058221196077fb19953ac608c7))
- People who want a GUI (though there are some [third-party efforts](https://github.com/ExoPlant/packwiz-gui))

Join my Discord server if you need help [here](https://discord.gg/Csh8zbbhCt)!

## Features
- Git-friendly TOML-based metadata format
- MultiMC pack installer/updater, with support for optional mods and fast automatic updates - perfect for servers!
- Pack distribution with HTTP servers, with a built in local server for testing
- Easy installation and updating of multiple mods at once from CurseForge and Modrinth
- Exporting to CurseForge and Modrinth packs
- Importing from CurseForge packs
- Server-only and Client-only mod handling
- Creation of remote file metadata from JAR files for CurseForge mods

## Installation
Prebuilt binaries are available from [GitHub Actions](https://github.com/wafflecoffee/packwiz/actions) - the UI is a bit terrible, but essentially select the top build, then download the artifact ZIP for your system at the bottom of the page. To run the executable, add the folder where you downloaded it to your PATH environment variable ([see tutorial for Windows here](https://www.howtogeek.com/118594/how-to-edit-your-system-path-for-easy-command-line-access/)) or move it to where you want to use it.

In future I will have a lot more installation options, but you can also compile from source:

1. Install Go (1.18 or newer) from https://golang.org/dl/
2. Run `go install github.com/wafflecoffee/packwiz@latest`. Be patient, it has to download and compile dependencies as well!

## Documentation
See https://packwiz.infra.link/ for the full packwiz documentation!