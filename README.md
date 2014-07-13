### Html5 FileBrowser, Angular + Go

![screenshot](builds/screenshot.gif)

#### Features

- Fast UI. Json + angular
- Directory fuzzy search / Acces
- Inline search ( current list )
- Upload mutliple files.
- Big uploads. Tested with 1G files. ( Uploads are streamed to disk )
- File delete / remove / copy / compress
- Dir creation
- File editor with Codemirror
- Also can act as an api file json server... ( Improbed in future relases )
    + PUT /dir file
    + GET /dir &format=json
- Download dirs as zip    

#### Install

Donwload a binary build: (Stable)

- [Osx 64bits](builds/file_server_osx)
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


