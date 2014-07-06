### HTTP FILE SERVER in GOlang

v0.9. Big code refactor (UI) using angular and port of almost all methods to 
api-json. 
At the momment only compatible with Browsers with xmlhttprequest 2.0 (firefox,safari,chrome) not tested in ie10, v0.9 is in branch devel.


--
I'm learning golang, and this is the first app I build, just for 
get inside the code. Missed and ugly tool for sysadmins, once, up and runing, 
shares current folder and subfolders via http, allowing uploading files, and 
on future file operations ( delete, rename, edit )

Start from commandline:

> file_server 
>  -dir="directory" ( default to current . )
>  -port=":8080"
>  -log=true/false show console logs...

And browser your system or server at:

http://localhost:8080



##### @Todo

+ ~~Add some type of flahs session, to notify user after an upload~~
+ ~~Add the current path, and a direct acces to parent path~~
- ~~Add file operations, like delete, move or~~ uncompress zip files..
- Big code refactor to milestone 1.0.
    - PACKAGE filsistem json .. with common operations... 
    - Frontend
- Add some kind of authentification
- Add some kind of permisions...
- Review log system.


##### Changelog

###### v0.9
+ Big refactor using Angular for frontend and api calls for actions

###### v0.5
+ Added version number
+ File deleting operations


