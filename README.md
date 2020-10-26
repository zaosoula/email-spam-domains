[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/zaosoula/email-spam-domains">
    <h3 align="center">Email Spam Domains</h3>
  </a>


  <p align="center">
  A list of 98,522 spam/temporary email domain available in multiples formats
    <br />
    <br />
    <a href="https://raw.githubusercontent.com/zaosoula/email-spam-domains/master/domains.txt">View .txt</a>
    ·
    <a href="https://raw.githubusercontent.com/zaosoula/email-spam-domains/master/domains.json">View .json</a>
    ·
    <a href="https://github.com/zaosoula/email-spam-domains/issues">Report Bug</a>
    ·
    <a href="https://github.com/zaosoula/email-spam-domains/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Sources](#sources)
  * [Built With](#built-with)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)



<!-- ABOUT THE PROJECT -->
## About The Project

### Sources
<small>Updated on 26 Oct 2020</small>
* [MailChecker](https://github.com/FGRibreau/mailchecker/blob/master/list.txt)  
* [StopForumsSpam](https://www.stopforumspam.com/downloads)
* [Disposable Email Domain List](https://github.com/groundcat/disposable-email-domain-list/blob/master/domains.txt)
* [Spam Domains List](https://github.com/tsirolnik/spam-domains-list/blob/master/spamdomains.txt)
* [Our custom list](https://github.com/zaosoula/email-spam-domains/blob/master/src/custom.txt)

### Built With

* [Go](https://golang.org/)


<!-- USAGE EXAMPLES -->
## Usage

### Utils

* [Merge](https://github.com/zaosoula/email-spam-domains/tree/master/utils/merge) - Compile sources in domains.txt
* [ToJson](https://github.com/zaosoula/email-spam-domains/tree/master/utils/toJson) - Convert domains.txt to domains.json
* [IsSpam](https://github.com/zaosoula/email-spam-domains/tree/master/utils/isSpam) - Check if a domain is in the list and return a 0 (false) or 1 (true)
``go run main.go -domain <email domain>``


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/zaosoula/email-spam-domains_name/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Zao Soula — contact@zaosoula.fr

Project Link: [https://github.com/zaosoula/email-spam-domains_name](https://github.com/zaosoula/email-spam-domains_name)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/zaosoula/email-spam-domains.svg?style=flat-square
[contributors-url]: https://github.com/zaosoula/email-spam-domains/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/zaosoula/email-spam-domains.svg?style=flat-square
[forks-url]: https://github.com/zaosoula/email-spam-domains/network/members
[stars-shield]: https://img.shields.io/github/stars/zaosoula/email-spam-domains.svg?style=flat-square
[stars-url]: https://github.com/zaosoula/email-spam-domains/stargazers
[issues-shield]: https://img.shields.io/github/issues/zaosoula/email-spam-domains.svg?style=flat-square
[issues-url]: https://github.com/zaosoula/email-spam-domains/issues
[license-shield]: https://img.shields.io/github/license/zaosoula/email-spam-domains.svg?style=flat-square
[license-url]: https://github.com/zaosoula/email-spam-domains/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/zaosoula
[product-screenshot]: images/screenshot.png
