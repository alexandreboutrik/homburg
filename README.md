# Homburg

<img align="right" width="128px" src="./.media/homburg.png">

![Version Badge](https://img.shields.io/badge/version-v0.0.0--prealpha-darkred)

Homburg is a backend solution for creating and managing forum platforms. It employs a proof-of-concept authentication system with asymmetric post-quantum cryptography, eliminating the need for traditional IDs and passwords. This approach not only enforces privacy (against leaks) but also ensures protection against spoofing attacks, including those that could be perpetrated by the platform admins.  
Designed with simplicity and ease-of-use in mind, Homburg also offers robust configuration and management, providing endpoints for adminitrators to effectively control the forum environment.

## Development status

Statistics:

| Language  | Files  | Lines   | Blanks  | Comments |
|:---------:|:------:|:-------:|:-------:|:--------:|
| Go        | 13     | 440     | 105     | 58       |
| Markdown  | 3      | 132     | 39      | 0        |
| Bash      | 1      | 40      | 8       | 1      |
| **Total** | **17** | **612** | **152** | **59**  |

Features:

| Feature                  | Status                             | Description |
|:-------------------------|:-----------------------------------|:------------|
| PGA                      | :black_square_button: In progress | Authentication based on asymmetric cryptography. |
| Forum & Thread           | :white_square_button: Planned     | CRUD endpoints for forums and threads. |
| Post                     | :white_square_button: Planned     | CRUD endpoints for posts. |
| Markdown                 | :white_square_button: Planned     | Markdown syntax support for posts. |
| Reporting                | :white_square_button: Planned     | Reporting system for posts and profile pictures. |
| Admin panel              | :white_square_button: Planned     | Endpoints for easy management by the admin. |
| Filter keywords          | :white_square_button: Planned     | Block and/or filter keyword rules for posts. |
| Likes                    | :white_square_button: Planned     | Like system for posts |
| User titles              | :white_square_button: Planned     | User title system according to a set rule of achievements. |
| Captcha                  | :white_square_button: Planned     | Captcha support. |
| Email notification       | :white_square_button: Planned     | Configuration options to send a notification email to the user when something happens. |

## Instructions

Install the dependencies:

```bash
$ go get .
```

Configure the SQL server:

```bash
$ sudo vim /etc/postgresql-16/postgresql.conf
# uncomment or add the following lines:
listen_addresses = 'localhost'
port = 9900

$ sudo vim /etc/postgresql-16/pg_hba.conf # if exists
# change to
local all postgres trust
local all all md5

$ sudo rc-service postgresql-16 start
```

Run the server:

```bash
$ ./reset.sh
$ go run .
```

## Frontend

The available frontend are listed below:

- None.

<!---
| Frontend | Plataform | Stable | Version |
| [hwc](https://github.com/alexandreboutrik/homburg-web-client) | browser | x | 0.0.1 |
--->

You can also implement your own following the REST API documentation.

## Documentation

| REST api                                   | DEV guide | USER guide | 
|:------------------------------------------:|:---------:|:----------:|
| [English](./documentation/RESTAPI-en.md)   | -         | -          |
| [Português](./documentation/RESTAPI-br.md) | -         | -          |
| [Français](./documentation/RESTAPI-fr.md)  | -         | -          |

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute the code as needed. See the [LICENSE](LICENSE) file for more information.
