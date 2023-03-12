module github.com/zhiyanfoo/module-a/seperate-lib2

go 1.20

require (
	github.com/zhiyanfoo/module-a v0.2.1-0.20230311205151-a1c18eabc27e // indirect
	github.com/zhiyanfoo/module-a/seperate-lib1 v0.0.0-20230311210701-9f89960a3c93 // indirect
)

replace (
    github.com/zhiyanfoo/module-a => ../localpath-a
)
