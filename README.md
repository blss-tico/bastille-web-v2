# Bastille-Web (This project is under development...)

> A web interface for the FreeBSD Bastille Jails management. 

## 💻 Requisites

Before begin, you must to know that this project was tested with:

- FreeBSD 14+;
- Golang 1.24+;
- Bastille 1.0.1.250714+;
- ttyd 1.7.7+;

## 🚀 Install bastille-web

To install the bastille-web, follow these steps.

Install dependencies:

```
pkg install -y go125 bastille ttyd
```

Clone the project and build it:

```
git clone https://github.com/blss-tico/bastille-web.git
cd bastille-web
go build
```

## ☕ Usando bastille-web

To use bastille-web, with default ip/port option:
```
sudo ./bastille-web
```

Ip address and port are defined in .env file.

## 📝 Licença

Esse projeto está sob licença. Veja o arquivo [LICENÇA](LICENSE.md) para mais detalhes.
