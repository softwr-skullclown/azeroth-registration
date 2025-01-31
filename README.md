# azeroth-registration

Azeroth Registration is an open source webapp for registering with an [AzerothCore](https://www.azerothcore.org/) private server, though may work with other projects like [MaNGOS](https://www.getmangos.eu/) or [TrinityCore](https://www.trinitycore.org/) as long as the primary auth database models are in alignment (PRs to add support for different models are welcome)

## Why??!

First off let me say I'm a huge fan of the WoW EMU open-source community, they have built a lot of awesome projects with their skill sets. When I decided to run my own private server for fun I looked for an easy registration page I could give friends access to.  Sadly all of the existing ones are written in PHP, which I absolutely despise with a passion (I did work with PHP for many years and it was great for its time) because of it's dynamic scripting nature, security risks, external dependencies to run (e.g. a PHP compatible webserver/fpm/insert solution), templating, you name it.

So here we are...

## What are the goals of this project?

1. Provide a registration and account management experience that makes use of the existing auth database and an SMTP server for simplicity.
1. Simple to run via a single binary or docker container with configured SMTP and Database values.

## Documentation

1. [Developing](./docs/DEVELOPING.md)