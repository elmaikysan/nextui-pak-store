<div align="center">
<img src=".github/resources/banner.png" width="auto" alt="Mortar wordmark">

![GitHub License](https://img.shields.io/github/license/UncleJunVip/nextui-pak-store?style=for-the-badge)
![GitHub Release](https://img.shields.io/github/v/release/UncleJunVIP/nextui-pak-store?sort=semver&style=for-the-badge)
![GitHub Repo stars](https://img.shields.io/github/stars/UncleJunVip/nextui-pak-store?style=for-the-badge)
![GitHub Downloads (specific asset, all releases)](https://img.shields.io/github/downloads/UncleJunVIP/nextui-pak-store/Pak.Store.pak.zip?style=for-the-badge&label=Downloads)

</div>

---

## How do I setup Pak Store?

1. Own a TrimUI Brick and have a SD Card with NextUI configured.
2. Connect your device to a Wi-Fi network.
3. Download the latest Pak Store release from this repo.
4. Unzip the release download.
5. Copy the entire Pak Store.pak folder to `SD_ROOT/Tools/tg5040`.
6. Reinsert your SD Card into your device.
7. Launch `Pak Store` from the `Tools` menu and enjoy all the amazing Paks made by the community!

---

## I want my Pak in Pak Store!

Awesome! To get added to Pak Store you have to complete the following steps:

1. Create a `pak.json` file at the root of your repo. An example can be seen below.
2. Make sure your release is tagged properly and matches the version number in `pak.json`.
3. Make sure the file name of the release artifact matches what is in `pak.json`.
4. Once all of these steps are complete, please file an issue with a link to your repo.

---

## Sample pak.json
```json
{
  "name": "Pak Store",
  "version": "v0.1.0",
  "type": "TOOL",
  "description": "A Pak Store in this economy?!",
  "author": "K-Wall",
  "repo_url": "https://github.com/scalysoot/nextui-pak-store",
  "release_filename": "Pak.Store.pak.zip",
  "banners": {
    "BRICK": ".github/resources/banner.png"
  },
  "platforms": [
    "tg5040"
  ],
  "update_ignore": [
    "Folder/To/Ignore",
    "bin/FileToIgnore.jpg"
  ],
  "launch": "launch.sh"
}
```

---

Enjoy! ✌️