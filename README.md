# Quran Obsidian Template

I created this template to take notes while studying Quran, using the popular note-taking app [Obsidian](https://obsidian.md/)

### Folder Structure
Some folders are created by default to group similar topics together, and you can link the particular Ayats during your study.
The whole folder structure is as follows:
```
Quranic Guidance
| Concepts
| - Artifact
|   - Place of Worship
|   - ...
| - ...
| - Weather Phenomenon
|
| Surahs
| - 1 Al-Fatihah (The Opener)
| - ...
| - 114 An-Nas (Mankind)
|
| Names Of Allah.md
```

The Ayats are grouped in folders by Surah, and each Ayat is in a separate file inside the folder.
For example, the second Ayat of the last Surah appears as `114-2.md`
```
Surahs
| - ...
| - 113 Al-Falaq (The Daybreak)
|   - 113-1.md
|   - ...
| - 114 An-Nas (Mankind)
|   - 114-1.md
|   - 114-2.md
|   - ...
```


### Names of Allah
The 99 names of Allah are listed in a table along with their meanings.
Their location in the Quran is automatically linked to the particular Ayat(s).
You can hover over the link to see a preview, or Ctrl + Click to open that Ayat in a separate tab.

![image](https://user-images.githubusercontent.com/54888682/221333424-44ef609a-13e6-4366-8f83-f7da4f823434.png)



### Ayats
Each of the 6236 Ayats of the Quran is in a separate page, grouped by Surahs.

Tags are present in the metadeta at the top of every such page.
By default, a `meccan` or `medinan` tag is added to every Ayat according to its place of revelation.
You can add your own tags to group similar-themed Ayats together.

![image](https://user-images.githubusercontent.com/54888682/221333511-ecd04854-97a1-4019-b62d-5d5ba9fdba59.png)


### Graph View
![image](https://user-images.githubusercontent.com/54888682/221333572-00aecdce-716f-4197-beaf-956aa92504ae.png)


---

### How to Use
The easiest method is to download the suitable executable file according to your OS (Windows, MacOS, or Linux) from the [releases](https://github.com/omeiirr/quran-obsidian-template/releases/tag/v1.0.0) page and execute it on your machine. A new folder "Quranic Guidance" will be setup in the working directory as described in [Folder Structure](https://github.com/omeiirr/quran-obsidian-template/edit/main/README.md#folder-structure).

If you have Go installed on your machine, you can also use it to setup the project.
1. Clone this repository with `git clone https://github.com/omeiirr/quran-obsidian-template.git`
2. `cd quran-obsidian-template`
3. `go run .`

If you are unable to get the executable running on your system, you can directly download the compressed `QuranicGuidance` zip folder from [here](https://github.com/omeiirr/quran-obsidian-template/releases/download/v1.1.0/QuranicGuidance.zip) which contains all the features, and decompress it on your device.

Happy learning! :)

### Suggested Add-ons

You are suggested to install the following Obsidian plug-ins for the best experience:
- [Omnisearch](https://github.com/scambier/obsidian-omnisearch) for fuzzy search across the vault for specific words/phrases
- [Dynamic RTL](https://github.com/mwxgaf/obsidian-dynamic-rtl) for automatic right-to-left alignment of Arabic text.
- [Automatically Reveal Active File](https://github.com/shichongrui/obsidian-reveal-active-file) to reveal the active file automatically when you open it

The Arabic font used in the demonstrations is KFGQPC Uthmanic Script Hafs.

The theme being used is [Things](https://github.com/colineckert/obsidian-things).

### Contributing
All contributions are welcome. Please feel free to submit issues and PRs for feature additions, bug-fixes, or enhancements.

### References
[Quran corpus](https://corpus.quran.com) has been used to source the `Concepts`.
This project also uses the [quran-json API](https://github.com/risan/quran-json) and [allah-names](https://github.com/AzeemGhumman/allah-names) libraries.
I extend my gratitude to their authors for their work.
