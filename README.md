<p align="right" ><a href="https://github.com/addUsername/UiPath-var-counter/releases/" target="_blank"><img src="https://img.shields.io/badge/version-v0.3.1-blue?style=for-the-badge"/></a></p>
<h2 align="center">
  UiPath-var-counter<br/><br/>
  <img src="https://user-images.githubusercontent.com/60299373/141651241-6533d38e-fa4d-4da9-99e2-25798915ada2.png" />  
  <p align="left"><sub> Simple cli app that identifies variables (and args..) from an **UiPath** project by reading its `.xaml` files</sub></p>
</h2>

## ⚡️ Quick start

> ### **Download** binary from [releases](https://github.com/addUsername/UiPath-var-counter/releases/)
>
> ### Copy it to UiPath project's **root folder**
> 
> ### Right click > **open terminal**

```Powershell
myvars -help
```

## ⚙️ Building way of quick start
> ### <a href="https://golang.org/doc/install" target="_blank">Install go<img src="https://user-images.githubusercontent.com/60299373/141649494-f7a52c41-8267-471f-aef9-994c9a217cb8.png" width= "40px"/></a> (if needed)
``` Powershell
git clone https://github.com/addUsername/UiPath-var-counter

cd UiPath-var-counter

go build!


./myvars . 
 ``` 

## 💡 How to
> ### 📌 Path
```Powershell
myvars -path="where_project_is_placed"           # Omit or Use "." to indicate current path
```
> ### 💾 Export table as .csv
```Powershell
myvars -csv="where_export.csv_should_be_placed"  # Use "." to indicate current path
```
> ### 📰 Do it! but.. arguments
```Powershell
myvars -args  
```
> ### 💊 Show usage
```Powershell
myvars -help
```

## 🧰 Install

Run console as Admin and type `./myvars -install`, it places the binary in the specified folder (default `C:\Program Files`) and add it to PATH, now you could use `myvars` wherever you please, or do it manually.
