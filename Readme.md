# arc-dps-launcher

The base idea of this project is to provide a simple method of updating the [GuildWars2](https://www.guildwars2.com) tool [arcdps, a gw2 dps meter (and general combat metrics tool)](https://www.deltaconnected.com/arcdps/).

## Reason for development

I'm playing GuildWars2 for a lot of years now and the only tool i've ever used, bringing A LOT of benefits to the game, is arcdps. The tool is just awesome by itsself but has one little problem I always got annoyed by: the installation and update mechanism. Its not complicated at all, placing the right .dll file in the right folder but EVERY TIME Arenanet patches the game arcdps needs to be updated, therefore requiring to be downloaded and placed in the right folder again and again (to avoid game crashes).

This is where this project shines (I hope). Everytime you start your game you just have to run the arc-dps-launcher and your current version of arc-dps is automatically updated to the latest version. After the update is done, it automatically launches GuildWars2 for you. If arc-dps isnt installed it even installs it for you.

I developed a lot of versions of this idea in the past few years starting with a really hacky java and python application but always kept them for me. The reason behind is the pain of installation in case of java and a lot of problems and erros due to a lack of development experience and no motivation to actually step up with the project. With me learning GO I figured its finally time to publish this tool and improving it. Hope you enjoy!

## Features

* Easy installation
* Installs Arc-Dps if necessary
* Keeps up to date version of Arc-Dps
* Launches GuildWars2
* Supports Dx9 and Dx11

## Installation

Simply download the latest pre-compiled binaries from the release-section (either for Dx9 or Dx11) and place it in the same folder your Gw2.exe is located. If you want the fancy way create a link on your desktop so you can launch GuildWars2 with the arc-dps-updater from there.
