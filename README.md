### Portable filebrowser with mobile ui html5

#### Features
- Mobile UI with almost all "usable displays", android and ios ( on ios, can't upload files)
- Fast UI. Json + angular
- Directory fuzzy search / Acces ( style textmate command+T)
- Inline search ( current list )
- Upload mutliple files.
- Big uploads. Tested with 1G files. ( Uploads are streamed to disk )
- File delete / remove / copy / compress
- Dir creation
- File editor with Codemirror ( javascript, html, css, php.. )
- Filesystem json server
    - POST /dir action=createFolder source=name > Will create a folder in dir with name source
    - POST /dir action=delete source=name > will delete
    ... see commands.go
    + GET /dir &format=json
- Download dirs as zip    
- File and video stream.


![screenshot](builds/screenshot.gif)


#### Install

Donwload a binary build: (Stable)

- [Osx 64bits](builds/file_server_darwin_amd64)
- [Linux 64bits](builds/file_server_linux_amd64)

Or compile it:
```go
go get github.com/jordic/fileserver

go build or go install
```

#### Browser compatibility
- Firefox, safari, Chrome.
- Perpahs ie10 but not tested

##### @Todo

+ ~~Add some type of flahs session, to notify user after an upload~~
+ ~~Add the current path, and a direct acces to parent path~~
+ ~~Add file operations, like delete, move or~~ ~~compress zip files..~~
+ ~~Big code refactor to milestone 1.0.~~ 
+ ~~Improbe filesystem json server.~~

- Backend. Add system commands as plugin.. with System services or commands 
    The commands must be, system commands, and should be configured, 
    on json. App, only loads them, and handles execution of them
- Create docs
- UI. Add a Generic command with output ( Perhaps a modal )
- UI. Improve javascript prompt, with some kind of widget
- UI. Add a button on toolbar, with shortcuts to system commands
- UI/Backend. Add a bookmark system. Perhaps a file .bookmark.js on root.

- Back. Add param for CORS handling
- Back. Add some kind of authentification
- Back. Add some kind of permisions...
- Back. Review log system.



##### Changelog

###### v0.9
+ Big refactor using Angular for frontend and api calls for actions

###### v0.5
+ Added version number
+ File deleting operations


