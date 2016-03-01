## The problem
So one day I was away from home and wanted to torrent a TV show to watch when I got home. Most torrent clients offer remote access, however my desktop is not accessible from outside of my home network.

## The solution

So I made this server and client to initiate torrent downloads via Transmission from remote locations. Currently, you start and leave running one service on your torrenting computer. Then you invoke the client from your remote computer, once per torrent file.

I used this project as an excuse to try out three new technologies, so it's not the most elegant.
* pusher.com - Pub/sub as a service. Their APIs are really nice, however their free tier is pretty restrictive on how much traffic I can handle.
* Kotlin - I wrote the server side of this system in Kotlin. The language was fast to pick up because of the "synergy" (I can't believe I'm using that word) of the JVM ecosystem.
* Go - I've dabbled with Go before, but this was the first Go project I built from scratch. Essentially it is just a CLI that allows you to upload files.

When I have time, I will write an iOS client for when I don't have a laptop with me.

## Usage

### Server

You will need Java, Gradle and Transmission installed.

    git clone https://github.com/aj-michael/remotetorrent-backend
    cd remotetorrent-backend
    gradle build
    java -jar build/libs/remotetorrent-1.0.0.jar /home/adam/TorrentDownloads


### Client

You need to have Go installed and paths set up correctly.

    go get github.com/aj-michael/remotetorrent
    go install github.com/aj-michael/remotetorrent
    remotetorrent ubuntu-14.04-iso.torrent
