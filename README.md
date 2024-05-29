# ASCII Art Output

## Project Overview

ASCII Art Output is a command-line tool developed in Go that converts text input into ASCII art using predefined character sets. This tool allows users to specify an output file and choose from multiple ASCII art styles to enhance the visual representation of their text.

### Objectives

The primary objective of this project is to provide a flexible and user-friendly tool for generating ASCII art from text input. It offers the following features:

- Conversion of text input into ASCII art using customizable character sets.
- Support for specifying an output file for the ASCII art.
- Multiple ASCII art styles to choose from, including standard, shadow, and thinkertoy.


<!-- TABLE OF CONTENTS -->
<details>
  <summary style="font-weight: bold; font-size: 1.4em;" >Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#optional-ascii-art-styles">Optional ASCII Art Styles</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


## About The Project

ASCII Art Output is designed to provide users with a simple yet powerful tool for creating ASCII art from text input. It utilizes a set of ASCII characters to represent text in a visually appealing manner. The project consists of three main components:

- **main.go**: The entry point of the program, responsible for parsing command-line arguments, loading ASCII characters, and generating ASCII art.
- **loadascii.go**: A package that handles the loading of ASCII characters from file.
- **printascii.go**: A package that prints the ASCII art to the specified output file.

### Built With

- Go Programming Language

## Getting Started

To get started with ASCII Art Output, follow the instructions below.

### Prerequisites

Before running the program, ensure that you have the following prerequisites:

- Go installed on your machine.
- Basic understanding of Go programming language.

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Vincent-Omondi/ascii-art-output.git
    ```

2. Navigate to the project directory:

    ```sh
    cd ascii-art-output
    ```

## Usage

To use ASCII Art Output, follow these steps:

1. Prepare a text file containing the text you want to convert into ASCII art.

2. Run the program with the following command:

    ```sh
    go run . --output=<outputFileName.txt> <textToConvert>
    ```

    Replace `<outputFileName.txt>` with the desired name for the output file, and `<textToConvert>` with the text you want to convert into ASCII art.

## Optional ASCII Art Styles

ASCII Art Output supports multiple ASCII art styles, including:

- **Standard**: A basic ASCII art style.
- **Shadow**: ASCII art with shadow effects.
- **Thinkertoy**: ASCII art with a playful design.

To specify a different ASCII art style, use the following command format:

```sh
go run . --output=<outputFileName.txt> <textToConvert> <style>
```

Replace `<style>` with one of the following options: `standard`, `shadow`, or `thinkertoy`.

## Roadmap

The following features are planned for future releases:

- Support for additional ASCII art styles.
- Add support for color.
- Add feature to specify text alignment
- Integration with third-party ASCII art libraries.
- Development of a graphical user interface (GUI) for easier usage.

## Contributing

Contributions to ASCII Art Output are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your fork.
5. Open a pull request to merge your changes into the main branch.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

If you have any questions or suggestions regarding ASCII Art Output, feel free to reach out to the project maintainers:


- **[X - @vinomondi_1](https://tweeter.com/vinomondi_1)**
- **[X - @seth](https://tweeter.com/seth)**
- **[Github - Vincent](https://github.com/Vincent-Omondi/)**
- **[Github - Seth](https://github.com/aTHOOH)**

<p align="right">(<a href="#project-overview">back to top</a>)</p>


## Acknowledgments

Special thanks to the creators of the ASCII art character sets used in this project.

