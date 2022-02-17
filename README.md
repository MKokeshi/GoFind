# GoFind

## How to Run?
1. Clona la repository.
2. Setta il tuo `GOPATH` dalla root della repository.
3. Runna `make install` nella root della repository.
4. Runna gofind così:
```shell script
gofind  -start <path-to-starting-directory> -pattern <pattern-to-match-against>
```

## Note
1. I caratteri regex speciali devono essere esclusi quando si fornisce l'opzione '-pattern'.
2. Per impostazione predefinita, proprio come 'find' di GNU, 'gofind' mostrerà tutti gli errori che incontra in 'stderr'. Se desideri silenziare tali errori, assicurati di reindirizzare "stderr" al dispositivo "null", ovvero virare su un "2> / dev / null" alla fine del comando.

## Esempio
Supponiamo che il percorso assoluto della directory di lavoro corrente sia `/Users/Joker/DevilProjects` e che abbia una gerarchia come segue:
```shell script

.
├── Makefile
├── README.md
├── bin
├── pkg
└── src
    └── user
        └── gofind
            ├── find
            │   └── versions
            │       ├── v1
            │       │   └── finder_v1.go
            │       └── v2
            │           └── finder_v2.go
            ├── finder.go
            ├── finder_test.go
            ├── main.go
            ├── parse_flags_test.go
            ├── string_queue.go
            └── string_queue_test.go
```
Se volessimo runnare:
```shell script
gofind -start . -pattern "finder.*\.go" 2>/dev/null
```
L'output sarebbe:
```shell script
start: /Users/Joker/DevilProjects, pattern: finder.*\.go
Utilizzando 12 workers !
/Users/Joker/DevilProjects/src/user/gofind/finder.go
/Users/Joker/DevilProjects/src/user/gofind/finder_test.go
/Users/Joker/DevilProjects/src/user/gofind/find/versions/v1/finder_v1.go
/Users/Joker/DevilProjects/src/user/gofind/find/versions/v2/finder_v2.go
```