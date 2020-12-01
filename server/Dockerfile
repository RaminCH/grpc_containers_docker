# 1-ya stupen sborki 
#sborka na osnove goroutine
FROM golang:alpine as builder 

#Obnovlayu konfigi -> slepku OS v konteynere
#i zagrujayu 'git' dla podderjki 'go.mod'
RUN apk update
RUN apk upgrade
RUN add --no-cache git 

#Sozdayu rabochuyu direktoriyu proyekta v konteynere
RUN mkdir /app 
#Perexodim v rabochuyu direktoriyu v konteynere
WORKDIR /app

#Kopiruyem vse iz direktorii 'server' v 'app'
COPY . .

#Yesli zabili sozdat go.mod i go.sum (go mod init) -> delayem sebe straxovku, yesli ix net opovestit -> yesli yest to vse ok
COPY go.mod .
COPY go.sum . 

#Pozvolayet poluchit vse zavisimosti iz go.mod
RUN go mod download


#CGO ENABLED - pozvolayet ispolzovat/ne ispolzovat komplilator CGO
#dla sborki nozkourovnevix bibliotek yazika Go napisannix na C

#CGO_ENABLED=0 oznachayet chto mi perekompiliruyem vse paketi pod okrujeniye yazika Go
#GOOS - pod kakoye okrujeniye ispolzuyem kompilator yazika Go 
#go build -a -installsuffix cgo -> sobrat proyek s uchetom obxoda CGO
# -o -> kak nazvat vixodnoy biinarnik (v nashem sluchaye 'server')
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server 

#na etom etape mi imeyem GORU NENUJNOGO MUSORA v goroutine + server.bin


# 2-ya stupen sborki 
# Podgotavlivayem rabochiy konteyner
# v kotorom budet minimum funktsionala, neobxodimogo dla vipolneniya .bin (ex: server.bin)


