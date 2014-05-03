### HTTP FILE SERVER in GOlang


I'm learning golang, and this is the first app I build, just for 
get inside the code. Missed and ugly tool for sysadmins, once, up and runing, 
shares current folder and subfolders via http, allowing uploading files, and 
on future file operations ( delete, rename, edit )

Start from commandline:

> file_server 
>  -dir="directory" ( default to current . )
>  -port=":8080"
>  -log=true/false show console logs...



### @Todo

+ Add some type of flahs session, to notify user after an upload
+ Add the current path, and a direct acces to parent path
- Add file operations, like delete, move or uncompress zip files..
- Refactor code, i'm building as learning golang.





