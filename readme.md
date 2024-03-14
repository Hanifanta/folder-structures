## Generate Folder Structure for Go Project
- ---
### Description
 This is a simple script to generate a folder structure for a Go project. It creates the following folders:
- app
- db 
  - migrations
- delivery 
  - http { controller, middleware, dto{request, response}, router}
- domain
  - entity
  - repository
  - usecase
- shared
  - config
  - logger
  - token
  - utils

- --
### How to use
- Clone the repository
```bash
git clone https://github.com/Hanifanta/folder-structures.git
```
- Copy `main` file to your project folder and run it.
- ---
### Author
- [Hanifanta](https://github.com/Hanifanta)