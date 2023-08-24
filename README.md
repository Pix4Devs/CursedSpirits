<p align="center">
<strong><a href="https://github.com/Pix4Devs/CursedSpirits">Github</a></strong>
|
<strong><a href="https://github.com/Pix4Devs">Other Tools</a></strong>
|
<strong><a href="https://discord.gg/bNhXhypyTS">Discord</a></strong>
|
<strong><a href="https://pix4.dev">Pix4.dev</a></strong>
</p>

# L7 Stress Tester - CrusedSpirits
CursedSpirits is a powerful stress testing tool designed to assess the robustness and performance of your web applications through Layer 7 stress testing. This tool allows you to simulate heavy traffic.

<p align="center">
<img src="https://www.hindustantimes.com/ht-img/img/2023/07/28/1600x900/Screenshot_2023-07-27_234919_1690524990508_1690525009794.png" width="450" class="frame">
</p>

## Stats
We have recorded a whopping 400 000 requests per second on our 8 core dedicated server with 1 Gbps network bandwidth.

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

This project is licensed under the [MIT License](https://github.com/Pix4Devs/CursedSpirits/LICENSE).

----------

Pix4Devs - Empowering Developers with Powerful Tools, Visit us at [https://pix4.dev](https://pix4.dev/)
