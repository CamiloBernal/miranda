# Miranda Project
A golang + qt Evernote Client

**Stack**:

For UI: Qt binding for Go: https://github.com/therecipe/qt

Ideal project structure:

your-application-name
├── cmd
|   ├── your-application-name-cli
|   |    └── main.go
|   ├── your-application-name-api
|   |    └── main.go
├── internal/
|   ├── package01
|   |    ├── file01.go
|   |    └── file02.go
|   ├── package02
|   |    ├── file01.go
|   |    └── file02.go
├── pkg/
|   ├── library01
|   |    ├── file01.go
|   |    └── file02.go
|   ├── library02
|   |    ├── file01.go
|   |    └── file02.go
├── vendor/
├── LICENSE
└── README.md

*ref: https://blog.friendsofgo.tech/posts/como_estructurar_tus_aplicaciones_go/


# How to generate code step by step

1. Install the latest Thrift. *(`sudo pacman -S thrift` if you are using Arch)*
1. Clone evernote-thrift repo `https://github.com/evernote/evernote-thrift`
1. Generate golang version specs through this command:

     ```bash
     thrift -strict -nowarn \
       --allow-64bit-consts \
       --allow-neg-keys \
       --gen go:package_prefix=github.com/camilobernal/miranda/,thrift_import=github.com/apache/thrift/lib/thrift \
       -strict -nowarn --allow-64bit-consts \
       --allow-neg-keys \
       --gen go:package_prefix=github.com/camilobernal/miranda/,thrift_import=github.com/apache/thrift/lib/go/thrift \
       -I src/ -r \
       --out github.com/camilobernal/miranda src/UserStore.thrift
     ```
   (thanks to:  github.com/dreampuf/evernote-sdk-golang )

