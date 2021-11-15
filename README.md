<p align="right" ><a href="https://github.com/addUsername/UiPath-var-counter/releases/" target="_blank"><img src="https://img.shields.io/badge/version-v0.3.5-blue?style=for-the-badge"/></a></p>
<h2 align="center">
  UiPath-var-counter<br/><br/>
  <img src="https://user-images.githubusercontent.com/60299373/141651241-6533d38e-fa4d-4da9-99e2-25798915ada2.png" />  
  <p align="left"><sub> Simple cli app that identifies variables (and args..) from an <b>UiPath</b> project by reading its `.xaml` files</sub></p>
</h2>

## ⚡️ Quick start

> ### **Download** binary from [releases](https://github.com/addUsername/UiPath-var-counter/releases/)
>
> ### Go to download folder > Right click > **open terminal**
```Powershell
./myvars -path="where_UiPath_project_is"
```

## ⚙️ Building way of quick start
> ### <a href="https://golang.org/doc/install" target="_blank">Install go<img src="https://user-images.githubusercontent.com/60299373/141649494-f7a52c41-8267-471f-aef9-994c9a217cb8.png" width= "40px"/></a> (if needed, there is a[ Dockerfile](https://github.com/addUsername/UiPath-var-counter/blob/main/Dockerfile) also)
``` Powershell
git clone https://github.com/addUsername/UiPath-var-counter

cd UiPath-var-counter

go build

./myvars
 ``` 

## 💡 How to
⚠️If the binary is not installed use `./myvars` or ` ` intead⚠️
> ### 📌 Path
```Powershell
myvars -path="where_project_is_placed"            # Omit or Use "." to indicate current path
```
> ### 💾 Export as .json
```Powershell
myvars -json="where_output.json_should_be_placed" # Use -json="." to indicate current path
```
> ### 📰 Do it again! but.. more
```Powershell
myvars -args -default -path=..                    # Show arguments and default values
```
> ### 💊 Show usage
```Powershell
myvars -help
```

## 🧰 Install
Sadly, install option is not curently working as expected, but just a 2 step procces:
- Copy the binary downloaded/builded before in to the desired folder (Ex: `C:\Program Files\myvars\`)
- Add that route to PATH, by searching "path" on windows bar > environment variables > PATH (Ex: `C:\Program Files\myvars\myvars.exe`)

Now the command "myvars" is defined in powershell