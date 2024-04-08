# PSN Downloader

`psnupdatedownloader` is a command-line tool written in Go that allows users to download update packages for PlayStation games by providing the game ID.

## Building (not required if using release binaries)

Before you can run `psnupdatedownloader`, make sure you have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).

## Installation

To install `psnupdatedownloader`, follow these steps:

1.  Clone the repository or download the source code to your local machine.
2.  Navigate to the directory containing the source code.
3.  Compile the program by running `go build -o psnupdatedownloader`.

## Usage

To use `psnupdatedownloader`, run the executable from the command line, followed by the game ID you wish to download updates for:

`./psnupdatedownloader <GameID>`

For example, to download updates for the game with ID BCUS12345, you would run:

`./psnupdatedownloader BCUS12345`

The tool will create a directory with the game ID as its name and download all available update packages into this directory.

## License

`psnupdatedownloader` is licensed under the GNU General Public License v3.0 (GPLv3), a popular open-source license that ensures users are free to use, modify, and distribute the software.

For more details, see the LICENSE file or visit GNU General Public License v3.0.
