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

![image](https://user-images.githubusercontent.com/54888682/211848142-5abbfbdd-4fac-4bf6-a79d-b988ab6a4405.png)


### Ayats
Each of the 6236 Ayats of the Quran is in a separate page, grouped by Surahs.

Tags are present in the metadeta at the top of every such page.
By default, a `meccan` or `medinan` tag is added to every Ayat according to its place of revelation.
You can add your own tags to group similar-themed Ayats together.

![image](https://user-images.githubusercontent.com/54888682/211851606-b35c3b16-a493-452a-a049-f45e76acc98b.png)


### Graph View
![image](https://user-images.githubusercontent.com/54888682/211851887-903016c7-9969-4e77-b4fc-54b3c1ecf42d.png)

---

### How to Use
The easiest method is to download the suitable executable file according to your OS (Windows, MacOS, or Linux) from the [releases](https://github.com/omeiirr/quran-obsidian-template/releases/tag/v1.0.0) page and execute it on your machine. A new folder "Quranic Guidance" will be setup in the working directory as described in [Folder Structure](https://github.com/omeiirr/quran-obsidian-template/edit/main/README.md#folder-structure).

If you have Go installed on your machine, you can also use it to setup the project.
1. Clone this repository with `git clone https://github.com/omeiirr/quran-obsidian-template.git`
2. `cd quran-obsidian-template`
3. `go run .`

### Contributing
All contributions are welcome. Please feel free to submit issues and PRs for feature additions, bug-fixes, or enhancements.

### References
[Quran corpus](https://corpus.quran.com) has been used to source the `Concepts`.
This project also uses the [quran-json API](https://github.com/risan/quran-json) and [allah-names](https://github.com/AzeemGhumman/allah-names) libraries.
I extend my gratitude to their authors for their work.
