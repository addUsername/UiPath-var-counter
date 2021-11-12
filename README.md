# UiPath-var-counter


Simple cli program that identify variables from an UiPath project

## What
Simple cli program that identify variables from an UiPath project

``` 
PS \> myvars -path="C:\my_UiPath_project" -verbose

[1 - 2]    Main.xaml 
┌───────────────────────────┬─────────────────────────────┬──────┐
  var                       │   String                    │   0
  variable2                 │   String                    │   2
  variableInner             │   Int32                     │   1
  variable3                 │   String                    │   0
└───────────────────────────┴─────────────────────────────┴──────┘
```

## How
#### Easy
- Download myvars executable from realeases
#### test this xd
- [Install go](https://golang.org/doc/install) (if needed)
- `git clone https://github.com/addUsername/UiPath-var-counter`
- `cd myvars`
- `go build`