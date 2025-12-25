
# Gotion
**Gotion** is a lightweight CLI tool that lets you create, list, and manage notes directly from your terminal. It stores everything locally and is perfect for quick note-taking without leaving your command line.



###  4. Install to the local machine
| Action        | Command           | Result                                                  |
|---------------|-------------------|---------------------------------------------------------|
| Development   | `make build`      | Creates gotion.exe in your current folder for quick testing. |
| New Release   | `git tag v1.0.1`  | Marks your code with a version number.                  |
| Official Install | `make install` | Installs the new version to your system globally.       |
| Run Anywhere  | `gotion`          | Run your app from any folder or drive.                  |


### Automated Release Process

step 1: 
  git add .
  git commit -m "Add new awesome feature"
  git push origin main

step 2:
  git tag v1.1.0
  git push origin v1.1.0