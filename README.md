<p align="center">
<strong><a href="https://github.com/Pix4Devs/CursedSpirits">Github</a></strong>
|
<strong><a href="https://github.com/Pix4Devs">Other Tools</a></strong>
|
<strong><a href="https://discord.gg/bNhXhypyTS">Discord</a></strong>
|
<strong><a href="https://pix4.dev">Pix4.dev</a></strong>
</p>

# L7 Stress Tester - CursedSpirits
CursedSpirits is a powerful stress testing tool designed to assess the robustness and performance of your web applications through Layer 7 stress testing. This tool allows you to simulate heavy traffic.

<p align="center">
<img src="https://www.hindustantimes.com/ht-img/img/2023/07/28/1600x900/Screenshot_2023-07-27_234919_1690524990508_1690525009794.png" width="450" class="frame">
</p>

> **Exterminator:**<br>
> CLI tool that wraps and empowers CursedSpirits with some hot-reloads and amplifications.
> <br><a href="https://github.com/Z3NTL3/Exterminator/">View</a>

## Stats
We have recorded a whopping 400 000 requests per second on our 8 core dedicated server with 1 Gbps network bandwidth.
<p align="center">
<img src="https://images-ext-1.discordapp.net/external/H9bTk-XvqRyQ5JjHgx19_mU1P6G_KsDS2_4USksEYLU/https/camo.githubusercontent.com/56f79ca67dbc72081b9619508e3e6b256e4621ba1953db2ce6710cceddfc0a72/68747470733a2f2f6d656469612e646973636f72646170702e6e65742f6174746163686d656e74732f3935363331303834303436343737333230302f313134333435303535323730363031313235362f696d6167652e706e673f77696474683d31343430266865696768743d363038?width=1440&height=607" >
</p>


## Getting Started
CursedSpirits requires golang to build. Follow the instructions at https://go.dev/doc/install to get your toolchain.
To use CursedSpirits, you need to follow these steps:

 1. Clone the Repository
	 ```sh
	 git clone https://github.com/Pix4Devs/CursedSpirits.git
	 cd CursedSpirits
	```
 2. Compile the Project by running `go build`
 3. **Populate the Context Directory:** CursedSpirits requires specific files in the context directory to run effectively.

	-   `proxies.txt`: A list of proxy addresses that will be used for sending requests.
	-   `accepts.txt`: A list of accept headers to be used in requests.
	-   `refs.txt`: A list of referer headers to be used in requests.

You can use the `scrape` command to automatically populate `proxies.txt` with scraped proxy addresses.
## Usage
CursedSpirits provides various commands to help you stress test your applications effectively. For information about these commands, run `./CursedSpirits help`

```plain
Usage:
  ./CursedSpirits [command]

Available Commands:
  check       Check utilities
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  scrape      Scrape proxies
  version     Prints out the current VERSION of CursedSpirits

Flags:
  -h, --help   help for ./CursedSpirits

Additional help topics:
  ./CursedSpirits start      Start flood

Use "./CursedSpirits [command] --help" for more information about a command.
```


We recommend running the following commands in succession, this assures the tool is ran with fresh and tested proxies.
```sh
./CursedSpirits scrape
./CursedSpirits check proxy
./CursedSpirits start --url <target url>
```
## Contributing

We welcome contributions from the community to enhance CursedSpirits. If you encounter any issues or have suggestions for improvements, please feel free to open an issue or create a pull request.

## Maintainers

 - [Z3ntl3](https://github.com/Z3ntl3)
 - [Midas](https://github.com/MidasVanVeen)

## License

This project is licensed under the [MIT License](https://mit-license.org).

----------

Pix4Devs - Empowering Developers with Powerful Tools, Visit us at [https://pix4.dev](https://pix4.dev/)
